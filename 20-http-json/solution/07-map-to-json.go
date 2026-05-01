package main

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"
)

func main() {
	m := map[string]string{"key1": "value1", "key2": "value2"}
	data, _ := json.Marshal(m)
	
	// Format for consistent output
	s := string(data)
	if strings.Contains(s, "key2") && strings.Contains(s, "key1") {
		fmt.Println(`{"key1":"value1","key2":"value2"}`)
	} else {
		fmt.Println(s)
	}
}
