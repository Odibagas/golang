package main

import (
	"fmt"
	"runtime"
)

func main() {
	os := runtime.GOOS
	switch os {
	case "darwin":
		fmt.Println("Go runs on Darwin")
	case "linux":
		fmt.Println("Go runs on Linux")
	case "freebsd":
		fmt.Println("Go runs on Freebsd")
	case "openbsd":
		fmt.Println("Go runs on Openbsd")
	case "plan9":
		fmt.Println("Go runs on plan9")
	case "windows":
		fmt.Println("Go runs on Windows")
	default:
		fmt.Println("Go runs on OS X")
	}
}
