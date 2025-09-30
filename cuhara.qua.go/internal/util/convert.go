package util

import (
	openapi_types "github.com/oapi-codegen/runtime/types"
)

func PtrToString(str *string) string {
	if str == nil {
		return ""
	}

	return *str
}

func StringToPtr(str string) *string {
	if str == "" {
		return nil
	}

	return &str
}

func PtrToInt64(num *int64) int64 {
	if num == nil {
		return 0
	}

	return *num
}

func Int64ToPtr(num int64) *int64 {
	if num == 0 {
		return nil
	}

	return &num
}

func Int64PtrToInt(num *int64) int {
	if num == nil {
		return 0
	}

	return int(*num)
}

func PtrToInt(num *int) int {
	if num == nil {
		return 0
	}

	return *num
}

func IntToPtr(num int) *int {
	if num == 0 {
		return nil
	}

	return &num
}

func EmailPtrToStringPtr(email *openapi_types.Email) *string {
	if email == nil {
		return nil
	}
	s := string(*email)
	return &s
}

