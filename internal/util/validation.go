package util

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func Validate[T any](data T) map[string]string {
	err := validator.New().Struct(data)
	result := map[string]string{}
	if err != nil {
		for _, value := range err.(validator.ValidationErrors) {
			result[value.StructField()] = TranslateTag(value)
		}
	}
	return result
}

func TranslateTag(fd validator.FieldError) string {
	if fd.ActualTag() == "required" {
		return fmt.Sprintf("field %s must be filled in", fd.StructField())
	}
	return "validation failed"
}
