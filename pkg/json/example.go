package main

import (
	"encoding/json"
	"fmt"
)

/*
func Marshal(v interface{}) ([]byte, error)
将struct对象序列化成json
将map序列化成json
将slice序列化成json

func Unmarshal(data []byte, v interface{}) error
将json反序列化成struct对象
将json反序列化到map中
将json反序列化到slice中
*/

type User struct {
	Username string
	Password string
	Firends  []string
}

func main() {
	user := User{}
	user.Username = "Tom"
	user.Password = "123456"
	user.Firends = []string{"Li", "Fei"}

	// 将struct 转成 json 字符串，结构体中的字段首字母必须大写，否则无法解析
	if userJSON, err := json.Marshal(user); err == nil {
		fmt.Println(string(userJSON))
	}

	// 将 slice 转成 json 字符串
	arr := []string{"Apple", "Orange", "Banana"}
	if arrJSON, err := json.Marshal(arr); err == nil {
		fmt.Println(arrJSON)
	}

	// 将 map 转成 json 字符串
	m := map[string]string{"id1": "name1", "id2": "name2"}
	if mJSON, err := json.Marshal(m); err == nil {
		fmt.Println(string(mJSON))
	}

	// json 转成 struct
	jsonStr := `{"Username":"Tom", "Password":"123456", "Firends":["fri1","fri2"]}`
	var userJSON User
	if err := json.Unmarshal([]byte(jsonStr), &userJSON); err == nil {
		fmt.Println(userJSON)
	}

	// json 转成 slice
	jsonFruit := `["Apple", "Orange", "Banana"]`
	var arrFruit []string
	if err := json.Unmarshal([]byte(jsonFruit), &arrFruit); err == nil {
		fmt.Println(arrFruit)
	}

	// json 转成 map
	jsonMap := `{"key1":"name1", "key2":"name2"}`
	var mapKey map[string]string
	if err := json.Unmarshal([]byte(jsonMap), &mapKey); err == nil {
		fmt.Println(mapKey)
	}

	str, _ := json.Marshal("test")
	fmt.Println(string(str))
}
