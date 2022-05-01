package helper

import (
	"fmt"

	"github.com/mitchellh/mapstructure"
)

func FillStruct_Interface(v interface{}, result interface{}) {
	data := v.(map[string]interface{})
	FillStruct(data, result)
}

func FillStruct(data map[string]interface{}, result interface{}) {
	if err := mapstructure.Decode(data, &result); err != nil {
		fmt.Println(err)
	}
}
