package main

import (
	"errors"
	"fmt"
)

type errUserNameExist struct {
	UserName string
}

func (e errUserNameExist) Error() string {
	return fmt.Sprintf("username %s already exist", e.UserName)
}

func IsErrUserNameExist(err error) bool {
	_, ok := err.(errUserNameExist)
	return ok
}

// 自定义错误 errors.New
func checkUserNameExist(username string) (bool, error) {
	if username == "foo" {
		return true, errUserNameExist{UserName: username}
	}

	if username == "bar" {
		return true, errors.New("username bar already exist")
	}
	return false, nil
}

func main() {
	if _, err := checkUserNameExist("foo"); err != nil {
		if IsErrUserNameExist(err) {
			fmt.Println(err)
		}
	}

	if _, err := checkUserNameExist("bar"); err != nil {
		if IsErrUserNameExist(err) {
			fmt.Println(err)
		} else {
			fmt.Println("IsErrUserNameExist is false")
		}
	}
}
