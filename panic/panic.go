package main

import (
	"fmt"
	"os"
	"time"

	//"github.com/pkg/errors"
)

type ErrorWithTime struct {
	message string
	timestamp int64
}

func (e *ErrorWithTime) Error() string {
	return fmt.Sprintf("error: %s\nunix-time: %d\n", e.message, e.timestamp)
}

func New(message string) error {
	return &ErrorWithTime{
		message: message,
		timestamp: time.Now().UnixMicro(),
	}
}

func main()  {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(result)
}

func divide(x, y int) (result int, err error)  {
	defer func() {
		if v, ok := recover().(error); ok {
			fmt.Println("Don't panic:", v)
			err = New(v.Error())
		}
	} ()

	return x / y, nil	
}