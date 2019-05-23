package monster

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

//Car struct
type Monster struct {
	Name  string
	Age   int
	Skill string
}

func (this *Monster) Store() bool {

	data, err := json.Marshal(this)
	if err != nil {
		return false
	}

	//save file
	filePath := "/Users/jessica/Dev/go/src/github.com/leighjpeter/go-learning/example-testing/monster.ser"
	err = ioutil.WriteFile(filePath, data, 0666)
	if err != nil {
		return false
	}
	return true
}

func (this *Monster) Restore() bool {
	filePath := "/Users/jessica/Dev/go/src/github.com/leighjpeter/go-learning/example-testing/monster.ser"
	data, err := ioutil.ReadFile(filePath)

	if err != nil {
		fmt.Println("ReadFile err =", err)
		return false
	}
	//2.使用读取到 data []byte ,对反序列化
	err = json.Unmarshal(data, this)
	if err != nil {
		fmt.Println("UnMarshal err =", err)
		return false
	}
	return true
}
