package common_test

import (
	"bytes"
	"encoding/json"
	"github.com/fenek-dev/go-twitter/src/common"
	"testing"
)

func NewStringResponse(code int, message string, data interface{}) string {
	r := common.NewResponse(code, message, data)
	out, err := json.Marshal(r)
	if err != nil {
		panic(err)
	}

	return string(out)
}

func NewBodyBuffer(t *testing.T, body interface{}) *bytes.Buffer {
	t.Helper()
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(body)
	if err != nil {
		t.Fatalf("Error reading response body: %v", err)
	}
	return &buf
}
