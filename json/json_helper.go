package json

import jsoniter "github.com/json-iterator/go"

var json = jsoniter.ConfigCompatibleWithStandardLibrary

// ToJsonString
func ToJsonString(data interface{}) string {
	result, _ := json.Marshal(data)
	return string(result)
}
