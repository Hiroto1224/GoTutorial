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

// Hellos 名前と挨拶を関連づけるマップを返却する
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		messages[name] = message
	}
	return messages, nil
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
