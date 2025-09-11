package common

import "strconv"

func Str2Int(str string) int {
	i, err := strconv.Atoi(str)
	if err != nil {
		return 0
	}

	return i
}

func Int2Str(i int) string {
	return strconv.Itoa(i)
}

func StrToStrPointer(str string) *string {
	if str == "" {
		return nil
	}

	return &str
}

func StrPointerToStr(str *string) string {
	if str == nil {
		return ""
	}

	return *str
}
