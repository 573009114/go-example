package demo

import (
	"encoding/json"
	"fmt"
)

//序列化数据
type Person struct {
	Name string
	Age  int
}

func JsonDecode() {
	p := &Person{Name: "张三", Age: 16}
	structJson, err := json.Marshal(p)
	if err != nil {
		fmt.Println("结构体序列化为json失败:", err)
	}
	fmt.Println(string(structJson))
}
