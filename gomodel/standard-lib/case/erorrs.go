package _case

import (
	"errors"
	"fmt"
	"log"
	"time"
)

type cuError struct {
	Code string
	Msg  string
	Time time.Time
}

func (err cuError) Error() string {
	return fmt.Sprintf("code:%s,msg:%s,time:%s", err.Code, err.Msg, err.Time.Format("2006-01-02 15:04:05"))
}
func getErrorCase(code, msg string) error {
	return cuError{
		Code: code,
		Msg:  msg,
		Time: time.Now(),
	}
}

func ErrorsCase() {
	var err error
	err = errors.New("this is a error")
	fmt.Println(err)

	var a, b = -1, -2
	res, err := sum(a, b)
	if err != nil {
		log.Println(res, err)
		cusErr, ok := err.(cuError)
		if ok {
			fmt.Println("cusErr:", cusErr)
		}

	}
}

func sum(a, b int) (int, error) {
	if a <= 0 && b <= 0 {
		return 0, getErrorCase("1001", "参数错误两数求和不能同时小于0")
	}
	return a + b, nil
}
