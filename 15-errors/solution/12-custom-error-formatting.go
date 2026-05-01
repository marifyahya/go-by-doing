package main

import "fmt"

type ApiError struct {
	Code int
	Msg  string
}

func (e ApiError) Error() string {
	return fmt.Sprintf("[ERROR] code=%d, msg=%s", e.Code, e.Msg)
}

func main() {
	err := ApiError{Code: 500, Msg: "internal error"}
	fmt.Println(err.Error())
}
