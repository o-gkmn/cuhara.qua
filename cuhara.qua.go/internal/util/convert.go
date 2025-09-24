package util

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
