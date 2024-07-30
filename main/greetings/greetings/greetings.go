package greetings

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	// 名前が定義されていない場合は、メッセージと共にエラーを返却する
	if name == "" {
		return "", errors.New("empty name")
	}

	// := 宣言と初期化を行うショートカット
	message := fmt.Sprintf("Hi, %v. welcome!", name)
	return message, nil
}
