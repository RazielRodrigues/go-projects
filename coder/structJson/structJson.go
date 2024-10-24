package main

import (
	"encoding/json"
	"fmt"
)

func main() {

	type Computer struct {
		ID   int    `json:"id"`
		NAME string `json:"name"`
	}

	computer := Computer{1, "Raziel"}

	json1, _ := json.Marshal(computer)

	fmt.Println(string(json1))

	var co2 Computer
	json.Unmarshal([]byte(json1), &co2)
	fmt.Println(co2)

}
