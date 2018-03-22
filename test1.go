package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Dictionary struct {
	Dict map[string]string `json:"datas"`
}

type Tdictmap struct {
	Records []Record `json:"RECORDS"`
}

type Record struct {
	Eng string `json:"eng"`
	Zh  string `json:"zh"`
}

func main() {
	b, err := ioutil.ReadFile("D:/hyf/liteide/tdict.json")
	if err != nil {
		fmt.Print(err)
	}
	var sp Tdictmap
	err = json.Unmarshal(b, &sp)
	if err != nil {
		panic(err)
	}

	Dictionart := make(map[string]string)

	for i, _ := range sp.Records {

		Dictionart[sp.Records[i].Eng] = sp.Records[i].Zh
	}

	var Recipient string = ""

	for {
		if Recipient == "E;" {
			break
		} else {
			fmt.Println("请输入需要查询的英文单词")
			fmt.Scanln(&Recipient)
			var Output = Dictionart[Recipient]
			Output, ok := Dictionart[Recipient]
			if ok {
				fmt.Println(Output)
			} else {
				fmt.Println("未找到")
			}
		}

	}
	//	if err != nil {
	//		fmt.Println(err)
	//	} else {
	//		fmt.Println(sp)
	//	}
	os.Exit(0)
}
