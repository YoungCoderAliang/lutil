package lutil

import (
	"encoding/json"
	"fmt"
)

func DecodeJson(s string) {
	var unknow interface{}
	json.Unmarshal([]byte(s), &unknow)
	printValue(unknow)
}

func printValue(unknow interface{}) {
	switch t := unknow.(type) {
	case map[string]interface{}:
		fmt.Print("{")
		pt := false
		for k, v := range t {
			if !pt {
				pt = true
			} else {
				fmt.Print(", ")
			}
			fmt.Print("\"" + k + "\":")
			printValue(v)
		}
		fmt.Print("}")
	case []interface{}:
		fmt.Print("[")
		for i, v := range t {
			if i > 0 {
				fmt.Print(", ")
			}
			printValue(v)
		}
		fmt.Print("]")
	case string:
		fmt.Print("\"")
		fmt.Print(unknow)
		fmt.Print("\"")
	default:
		// 数值打印会丢失精度
		fmt.Print(unknow)
	}
}
