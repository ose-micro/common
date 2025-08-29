package common

import "fmt"

// ToStringMap converts map[string]interface{} to map[string]string
func ToStringMap(in map[string]interface{}) (map[string]string, error) {
	out := make(map[string]string)
	for k, v := range in {
		switch val := v.(type) {
		case string:
			out[k] = val
		default:
			return nil, fmt.Errorf("invalid type for key %s: expected string, got %T", k, v)
		}
	}
	return out, nil
}

// ToInterfaceMap converts map[string]string to map[string]interface{}
func ToInterfaceMap(in map[string]string) map[string]interface{} {
	out := make(map[string]interface{})
	for k, v := range in {
		out[k] = v
	}
	return out
}
