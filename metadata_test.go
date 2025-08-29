package common_test

import (
	"testing"

	"github.com/ose-micro/common"
	"github.com/stretchr/testify/assert"
)

func TestToStringMap(t *testing.T) {
	raw := map[string]interface{}{
		"org:view":      "allow",
		"campaign:view": "deny",
	}

	result, err := common.ToStringMap(raw)
	assert.NoError(t, err)
	assert.Equal(t, map[string]string{
		"org:view":      "allow",
		"campaign:view": "deny",
	}, result)
}

func TestToStringMap_InvalidType(t *testing.T) {
	raw := map[string]interface{}{
		"org:view": true, // not a string
	}

	_, err := common.ToStringMap(raw)
	assert.Error(t, err)
}

func TestToInterfaceMap(t *testing.T) {
	input := map[string]string{
		"org:view":      "allow",
		"campaign:view": "deny",
	}

	result := common.ToInterfaceMap(input)

	assert.Equal(t, map[string]interface{}{
		"org:view":      "allow",
		"campaign:view": "deny",
	}, result)
}
