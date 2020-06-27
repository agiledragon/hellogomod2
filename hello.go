package hellogomod2

import (
	"errors"
	"fmt"
)

func Hello(who, t string) (string, error) {
	switch t {
	case "en":
		return fmt.Sprintf("Hello, %s!", who), nil
	case "cn":
		return fmt.Sprintf("你好, %s!", who), nil
	default:
		return "", errors.New("unknown language")
	}
}
