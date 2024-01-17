package controllers

import (
	"cloudreveMirror/serializer"
	"encoding/json"
	"fmt"
	"testing"
)

func TestErrorResponse(t *testing.T) {
	fmt.Println(ErrorResponse(fmt.Errorf("随便测试错误")))

	var resp serializer.Response
	err := json.Unmarshal([]byte(``), resp)
	fmt.Println(ErrorResponse(err))
}
