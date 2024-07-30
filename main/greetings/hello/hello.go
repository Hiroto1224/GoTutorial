package main

import (
	"fmt"
	// 他のフォルダのファイルのimport方法
	"greetings/greetings"
)

func main() {
	// greetingsファイルにて作成したHello関数を呼び出す
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
