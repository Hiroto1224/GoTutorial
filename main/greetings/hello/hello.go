package main

import (
	"fmt"
	"log"

	// 他のフォルダのファイルのimport方法
	"greetings/greetings"
)

func main() {

	// Loggerのプロパティを設定
	log.SetPrefix("greetings: ")
	// ログエントリのプリフィックスと印刷を無効にするフラグ
	log.SetFlags(0)

	names := []string{"Gladys", "Samantha", "Darrin"}

	// greetingsファイルにて作成したHello関数を呼び出す
	messages, err := greetings.Hellos(names)

	// エラーが返却された場合はコンソールにエラーを出力し、プログラムを終了する。
	if err != nil {
		log.Fatal(err)
	}

	for name, message := range messages {
		fmt.Printf("to: %v, message: %v\n", name, message)
	}
}
