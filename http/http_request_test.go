package http

import (
	"testing"
)

// go test -run ^Test.* . -v
func TestHttpGet(t *testing.T) {
	result, err := HttpGet("https://myip.ipip.net/", nil)
	if err != nil {
		t.Errorf("execute error : %v", err)
	}
	t.Logf("return data : %s", result)
}
