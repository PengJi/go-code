package main

import (
	"encoding/json"
	"fmt"
)

type addr struct {
	Province string `json:"province"`
	City string `json:"city"`
}

type stu struct {
	Name string `json:"name"`
	Age int `json:"age"`
	Addr addr `json:"addr"`
}

func main() {
	var xm = stu{Name:"zhangsan", Age:18, Addr:addr{Province:"prov1", City:"city1"}}

	js, err := json.Marshal(xm)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(js))

	var xxm stu
	err = json.Unmarshal(js, &xxm)
	fmt.Println(xxm)
}