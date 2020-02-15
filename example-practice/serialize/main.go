package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	// 序列化
	JsonStruct()
	JsonMap()
	JsonSlice()
	// 反序列化
	UnMarshalStruct()
	UnMarshalStruct2()
	UnMarshalMap()
	UnMarshalSlice()

}

type Monster struct {
	Name     string
	Age      int
	Birthday string
	Sal      float64
	Skill    string
	Pic      json.RawMessage
}

func JsonStruct() {
	str := `{"name":"ll","age":"11"}`
	bt := []byte(str)
	monster := Monster{Name: "iphone xr", Age: 10, Birthday: "2019-02-01", Pic: bt}
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Printf("序列号错误 err=%v\n", err)
	}
	fmt.Printf("struct序列化后=%v\n", string(data))
}

func JsonMap() {
	var map_a map[string]interface{}
	map_a = make(map[string]interface{})
	map_a["name"] = "leighj"
	map_a["age"] = 11
	data, err := json.Marshal(map_a)
	if err != nil {
		fmt.Printf("序列号错误 err=%v\n", err)
	}
	fmt.Printf("map序列化后=%v\n", string(data))
}

func JsonSlice() {
	var slice_a []map[string]interface{}
	var m1 map[string]interface{}
	m1 = make(map[string]interface{})
	m1["name"] = "leighj"
	m1["age"] = 12

	m2 := make(map[string]interface{})
	m2["name"] = "tom"
	m2["age"] = 20

	slice_a = append(slice_a, m1)
	slice_a = append(slice_a, m2)
	data, err := json.Marshal(slice_a)
	if err != nil {
		fmt.Printf("序列化错误 err=%v\n", err)
	}
	fmt.Printf("slice序列化后=%v\n", string(data))
}

func UnMarshalStruct() {

	var monster Monster // monster := Monster{}

	str := "{\"Name\":\"iphone xs\",\"Age\":10,\"Birthday\":\"2019-02-01\",\"Sal\":0,\"Skill\":\"\"}"
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Printf("反序列化错误 err=%v\n", err)
	}
	fmt.Printf("struct反序列化后=%v,monster.Name=%v\n", monster, monster.Name)
}

func UnMarshalStruct2() {

	var monster Monster // monster := Monster{}

	str := "{\"Pic\":{\"name\":\"ll\",\"age\":\"11\"}}"
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Printf("反序列化错误 err=%v\n", err)
	}
	fmt.Printf("struct反序列化后=%v,monster.Pic=%v\n", monster, string(monster.Pic))
}

func UnMarshalMap() {

	var map_a map[string]interface{}
	//注意:反序列化 map,不需要 make,因为 make 操作被封装到 Unmarshal 函数
	str := "{\"age\":11,\"name\":\"leighj\"}"
	err := json.Unmarshal([]byte(str), &map_a)
	if err != nil {
		fmt.Printf("反序列化错误 err=%v\n", err)
	}
	fmt.Printf("map反序列化后=%v\n", map_a)
}

func UnMarshalSlice() {

	var slice_a []map[string]interface{}
	//注意:反序列化 map,不需要 make,因为 make 操作被封装到 Unmarshal 函数
	str := "[{\"age\":12,\"name\":\"leighj\"},{\"age\":20,\"name\":\"tom\"}]"
	err := json.Unmarshal([]byte(str), &slice_a)
	if err != nil {
		fmt.Printf("反序列化错误 err=%v\n", err)
	}
	fmt.Printf("slice反序列化后=%v\n", slice_a)
}
