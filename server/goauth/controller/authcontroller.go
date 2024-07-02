package controller

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/sessions"
	"gogosing/server/goauth/domain"
	"gogosing/server/goauth/oauth"
	"gogosing/server/goauth/ui"
	"golang.org/x/oauth2"
	"io/ioutil"
	"net/http"
	"time"
)

type AuthController struct {
	SessionStore *sessions.CookieStore
}

const (
	Session = "Session"
	state   = "state"
	Code    = "Code"
)

func (controller *AuthController) Serve(mux *http.ServeMux) {
	mux.HandleFunc(controller.renderAuthView())
	mux.HandleFunc(controller.authenticationCallback())
}

// 유저 식별을 위한 랜덤 state 값을 가진 로그인 링크를 랜더링 해주는 핸들러
const AuthEndpoint = "/auth"

func (controller *AuthController) renderAuthView() (string, func(http.ResponseWriter, *http.Request)) {
	return AuthEndpoint,
		func(writer http.ResponseWriter, request *http.Request) {
			session, _ := controller.SessionStore.Get(request, Session)
			session.Options = &sessions.Options{
				Path:   AuthEndpoint,
				MaxAge: int(30 * time.Minute),
			}

			token := oauth.GetRandToken()
			fmt.Println("token -> ", token)
			session.Values[state] = token
			session.Save(request, writer)

			url := oauth.GetRandomLoginURL(token)
			println("[URL] -> " + url)
			ui.Render(writer, "auth", url)
		}
}

// OAuth 콜백 핸들러
const AuthCallbackEndpoint = "/auth/callback"

func (controller *AuthController) authenticationCallback() (string, func(http.ResponseWriter, *http.Request)) {
	return AuthCallbackEndpoint,
		func(writer http.ResponseWriter, request *http.Request) {
			session, _ := controller.SessionStore.Get(request, Session)

			savedToken, err := getSavedToken(session, writer, request)

			if err != nil {
				http.Error(writer, err.Error(), http.StatusInternalServerError)
				return
			}

			// Session에 저장된 토큰과 Request의 토큰 값을 비교한다. (검증)
			if savedToken != request.FormValue(state) {
				http.Error(writer, "Invalid token", http.StatusUnauthorized)
				return
			}

			token, err := oauth.Config.Exchange(oauth2.NoContext, request.FormValue(Code))
			if err != nil {
				http.Error(writer, err.Error(), http.StatusBadRequest)
				return
			}

			//client := oauth.Config.Client(oauth2.NoContext, token)
			client := oauth.Config.Client(context.Background(), token)
			userInfoResp, err := client.Get(AuthCallbackEndpoint)
			if err != nil {
				http.Error(writer, err.Error(), http.StatusBadRequest)
				return
			}
			defer userInfoResp.Body.Close()
			userInfo, err := ioutil.ReadAll(userInfoResp.Body)
			if err != nil {
				http.Error(writer, err.Error(), http.StatusBadRequest)
				return
			}
			var authUser domain.User
			json.Unmarshal(userInfo, &authUser)

			session.Options = &sessions.Options{
				Path:   AuthCallbackEndpoint,
				MaxAge: 86400,
			}

			session.Values["user"] = authUser.Email
			session.Values["username"] = authUser.Name
			session.Save(request, writer)

			http.Redirect(writer, request, HomeEndpoint, http.StatusFound)
		}
}

func getSavedToken(session *sessions.Session, writer http.ResponseWriter, request *http.Request) (string, error) {
	savedToken := session.Values[state]
	delete(session.Values, state)
	session.Save(request, writer)
	fmt.Println("savedToken -> ", savedToken)

	token, ok := savedToken.(string)
	if !ok {
		fmt.Println(token)
		return "", errors.New("변환 실패")
	}
	return token, nil
}
