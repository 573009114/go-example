package demo

import (
	"encoding/json"
	"fmt"
)

//反序列化
type People struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func JsonEndcode() {
	jsonStructStr := `{"name":"张三", "age": 18}`
	p := &People{}
	err := json.Unmarshal([]byte(jsonStructStr), &p)
	if err != nil {
		fmt.Println("Unmarshal  jsonStr err!", err)
	}

	fmt.Println(p.Name)
}
