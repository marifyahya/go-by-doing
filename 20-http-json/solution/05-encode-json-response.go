package main

import (
	"encoding/json"
	"fmt"
	"net/http/httptest"
)

func main() {
	w := httptest.NewRecorder()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"data": "value"})
	fmt.Printf("Content-Type: %s\n", w.Header().Get("Content-Type"))
	fmt.Print(w.Body.String())
}
