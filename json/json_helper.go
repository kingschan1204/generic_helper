package json

import (
	"fmt"
	jsoniter "github.com/json-iterator/go"
	"strconv"
	"strings"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// ToJsonString
func ToJsonString(data interface{}) string {
	result, _ := json.Marshal(data)
	return string(result)
}

func GetByExp(jsonData string, expression string) (bool, interface{}) {
	var v interface{}
	json.Unmarshal([]byte(jsonData), &v)
	exps := strings.Split(expression, ".")
	var result interface{} = nil
	var exits = false
	for i := 0; i < len(exps); i++ {
		if i == 0 {
			exits, result = Get(v, exps[i])
		} else {
			if exits {
				exits, result = Get(result, exps[i])
			}
		}
		if result == nil || !exits {
			break
		}
	}
	return exits, result
}

func Get(data interface{}, exp string) (bool, interface{}) {
	switch v := data.(type) {
	case []interface{}:
		index, _ := strconv.Atoi(exp)
		if index >= len(v) {
			fmt.Println("index out of capacity ！the array maximum  is ", len(v), "and index start with 0")
			return false, nil
		}
		return true, v[index]
	case map[string]interface{}:
		val, ok := v[exp]
		if ok {
			return true, val
		}
		return false, val
	default:
		fmt.Println("unknown -> ", v)
		return false, nil
	}
	return false, nil
}
