package common

import (
	"log"
	"strconv"
	"time"

	"google.golang.org/protobuf/types/known/anypb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
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

func StrPtr(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func IntPtr(i int) *int {
	if i == 0 {
		return nil
	}
	return &i
}

func BoolPtr(b bool) *bool {
	if b {
		return nil
	}
	return &b
}

func FloatPtr(f float64) *float64 {
	if f == 0 {
		return nil
	}
	return &f
}

func TimePtr(t time.Time) *time.Time {
	if t.IsZero() {
		return nil
	}
	return &t
}

func ToAny(val interface{}) *anypb.Any {
	if val == nil {
		return nil
	}

	var anyVal *anypb.Any
	var err error

	switch v := val.(type) {
	case string:
		anyVal, err = anypb.New(wrapperspb.String(v))
	case int:
		anyVal, err = anypb.New(wrapperspb.Int32(int32(v)))
	case int32:
		anyVal, err = anypb.New(wrapperspb.Int32(v))
	case int64:
		anyVal, err = anypb.New(wrapperspb.Int64(v))
	case float32:
		anyVal, err = anypb.New(wrapperspb.Float(v))
	case float64:
		anyVal, err = anypb.New(wrapperspb.Double(v))
	case bool:
		anyVal, err = anypb.New(wrapperspb.Bool(v))
	default:
		// If it's already a proto.Message, pack directly
		if msg, ok := v.(anypb.Any); ok {
			anyVal, err = anypb.New(&msg)
		} else {
			log.Printf("Unsupported type for Any: %T\n", val)
			return nil
		}
	}

	if err != nil {
		log.Println("Failed to pack Any:", err)
		return nil
	}

	return anyVal
}

func FromAnyToPrimitive(val *anypb.Any) interface{} {
	if val == nil {
		return nil
	}

	switch val.TypeUrl {
	case "type.googleapis.com/google.protobuf.StringValue":
		wrapped := &wrapperspb.StringValue{}
		if err := val.UnmarshalTo(wrapped); err != nil {
			log.Println("Failed to unmarshal string:", err)
			return ""
		}
		return wrapped.Value

	case "type.googleapis.com/google.protobuf.BoolValue":
		wrapped := &wrapperspb.BoolValue{}
		if err := val.UnmarshalTo(wrapped); err != nil {
			log.Println("Failed to unmarshal bool:", err)
			return false
		}
		return wrapped.Value

	case "type.googleapis.com/google.protobuf.Int32Value":
		wrapped := &wrapperspb.Int32Value{}
		if err := val.UnmarshalTo(wrapped); err != nil {
			log.Println("Failed to unmarshal int32:", err)
			return int32(0)
		}
		return wrapped.Value

	case "type.googleapis.com/google.protobuf.Int64Value":
		wrapped := &wrapperspb.Int64Value{}
		if err := val.UnmarshalTo(wrapped); err != nil {
			log.Println("Failed to unmarshal int64:", err)
			return int64(0)
		}
		return wrapped.Value

	case "type.googleapis.com/google.protobuf.UInt32Value":
		wrapped := &wrapperspb.UInt32Value{}
		if err := val.UnmarshalTo(wrapped); err != nil {
			log.Println("Failed to unmarshal uint32:", err)
			return uint32(0)
		}
		return wrapped.Value

	case "type.googleapis.com/google.protobuf.UInt64Value":
		wrapped := &wrapperspb.UInt64Value{}
		if err := val.UnmarshalTo(wrapped); err != nil {
			log.Println("Failed to unmarshal uint64:", err)
			return uint64(0)
		}
		return wrapped.Value

	case "type.googleapis.com/google.protobuf.FloatValue":
		wrapped := &wrapperspb.FloatValue{}
		if err := val.UnmarshalTo(wrapped); err != nil {
			log.Println("Failed to unmarshal float:", err)
			return float32(0)
		}
		return wrapped.Value

	case "type.googleapis.com/google.protobuf.DoubleValue":
		wrapped := &wrapperspb.DoubleValue{}
		if err := val.UnmarshalTo(wrapped); err != nil {
			log.Println("Failed to unmarshal double:", err)
			return float64(0)
		}
		return wrapped.Value

	default:
		log.Printf("Unsupported type: %s\n", val.TypeUrl)
		return nil
	}
}
