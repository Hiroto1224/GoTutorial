package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hello(name string) (string, error) {
	// 名前が定義されていない場合は、メッセージと共にエラーを返却する
	if name == "" {
		return "", errors.New("empty name")
	}

	// := 宣言と初期化を行うショートカット
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// 　用意した複数の挨拶をランダムで返却する
func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}
