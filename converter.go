package common

import (
	"strconv"
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

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

func ParseStrPtnToStr(str *string) string {
	if str == nil {
		return ""
	}

	return *str
}

func ParseTimeToTimestamp(val *time.Time) *timestamppb.Timestamp {
	if val != nil {
		return timestamppb.New(*val)
	}
	return nil
}

func ParseStrToPtn(val string) *string {
	if val == "" {
		return nil
	}

	return &val
}

func ParseTimestampToTimePtn(val *timestamppb.Timestamp) *time.Time {
	if val == nil {
		return nil
	}

	date := val.AsTime()
	return &date
}
