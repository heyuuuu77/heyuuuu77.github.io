package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	s := "zebra"

	fmt.Println(json.Marshal(s))
}
