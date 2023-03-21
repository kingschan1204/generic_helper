package json

import (
	"fmt"
	"testing"
)

func TestToJsonString(t *testing.T) {
	data := map[string]interface{}{"name": "test", "pass": true, "lv": 1}
	fmt.Println(ToJsonString(data))
}
