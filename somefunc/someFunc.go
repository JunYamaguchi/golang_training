package somefunc

import (
	"errors"
	"fmt"
)

func someFunc(someString string) (int, error) {
	if someString == "hoge" {
		fmt.Println("It's OK")
		return 0, nil
	}
	return 1, errors.New("someFunc() argument must be hoge")
}
