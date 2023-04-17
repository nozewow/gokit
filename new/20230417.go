package main

import (
	"fmt"

	jsoniter "github.com/json-iterator/go"
)

type Demo struct {
	Num []int64 `json:"num"`
}

func main() {
	d := Demo{}

	data := `{"num":[1,2,3]}`
	jsoniter.UnmarshalFromString(data, &d)

	for _, v := range d.Num {
		fmt.Println(v)
	}
	str := "lv1"
	fmt.Println(string(str[2]))
	fmt.Println(fmt.Sprintf("vip%v", string(str[2])))
	fmt.Println(len(str))
}
