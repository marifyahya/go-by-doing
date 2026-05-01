package main

import (
	"fmt"
	"net/http/httptest"
)

func main() {
	rr := httptest.NewRecorder()
	rr.Write([]byte("Response sent"))
	fmt.Println(rr.Body.String())
}
