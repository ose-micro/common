package common

import (
	"encoding/json"
	"fmt"
)

func JsonToAny(raw any, cast any) error {
	data, err := json.Marshal(raw)
	if err != nil {
		return fmt.Errorf("failed to marshal raw data: %w", err)
	}

	err = json.Unmarshal(data, &cast)
	if err != nil {
		return fmt.Errorf("failed to unmarshal to any: %w", err)
	}

	return nil
}

func AnyToJson(cast any) (string, error) {
	data, err := json.Marshal(cast)
	if err != nil {
		return "", fmt.Errorf("failed to marshal any: %w", err)
	}

	return string(data), nil
}
