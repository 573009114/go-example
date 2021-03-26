package demo

import "fmt"

func Mkmap() map[string]string {
	// 直接创建
	m2 := make(map[string]string)
	// 然后赋值
	m2["a"] = "aa"
	m2["b"] = "bb"
	return m2
}
func LoadMap() {
	m := Mkmap()
	//循环获取
	for k, v := range m {
		if k == "a" {
			fmt.Println(k)
			fmt.Println(v)
		}
	}
	//单独获取某个值
	fmt.Println(m["a"])
}
