package main

import (
	"fmt"
	"runtime"
	"time"
)

func printOs() {

	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X Yeah")
	case "linux":
		fmt.Println("Linuxxxx")
	default:
		fmt.Printf("다윈 리눅스 아님 -> %s\n", os)
	}
}

func isSaturday() {
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today is Saturday")
	case today + 1:
		fmt.Println("Tomorrow is Saturday")
	case today + 2:
		fmt.Println("2일 후가 토욜임")
	default:
		fmt.Println("아직 멀었음 ㅅㄱ")
	}
}

func sayHelloToMyFriend() {
	now := time.Now()
	switch {
	case now.Hour() < 12:
		fmt.Println("Good Morning! Yeah!!")
	case now.Hour() < 17:
		fmt.Println("Good Afternoon! Yeah~")
	default:
		fmt.Println("자라")
	}
}
