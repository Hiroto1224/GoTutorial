package greetings

import "fmt"

func Hello(name string) string {
	// := 宣言と初期化を行うショートカット
	message := fmt.Sprintf("Hi, %v. welcome!", name)
	return message
}
