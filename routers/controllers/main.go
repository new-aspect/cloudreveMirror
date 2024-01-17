package controllers

import (
	"cloudreveMirror/serializer"
	"fmt"
	"gopkg.in/go-playground/validator.v8"
)

// ErrorResponse 返回错误消息
func ErrorResponse(err error) serializer.Response {
	if ve, ok := err.(validator.ValidationErrors); ok {
		for _, e := range ve {
			return serializer.ParamErr(
				fmt.Sprintf("%s %s", e.Field, e.Tag),
				err,
			)
		}
	}
	return serializer.ParamErr("参数错误", err)
}
