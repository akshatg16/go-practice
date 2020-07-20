package main

import ("fmt"
       "encoding/json"
)
type detail struct {
    Title  string `json:"title"`
}

func main() {

	m := make (map [string]string)

	m["name"] = "akshat"
	m["address"] = "jaipur"

	fmt.Println("map:", m)

	v1 := m["name"]
	fmt.Println("v1: ", v1)

	v2 := m["address"]
	fmt.Println("v2: ", v2)
	
	fmt.Println("len:", len(m))
	
	 jsonString, _ := json.Marshal(m)

    fmt.Println(m)
    fmt.Println(jsonString)
}

	