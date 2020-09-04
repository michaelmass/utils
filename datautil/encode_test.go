package datautil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	nilValueExpectedJSON = "null"
	nilValueExpectedYaml = "null\n"

	intValueExpectedJSON = "1"
	intValueExpectedYaml = "1\n"

	boolValueExpectedJSON = "false"
	boolValueExpectedYaml = "false\n"

	mapStringExpectedJSON = "{\"\":\"\",\"empty\":\"\",\"key\":\"value\",\"test\":\"test\"}"
	mapStringExpectedYaml = "\"\": \"\"\nempty: \"\"\nkey: value\ntest: test\n"

	structValueExpectedJSON = "{\"Test\":\"testing\",\"Value\":123}"
	structValueExpectedYaml = "test: testing\nvalue: 123\n"
)

var (
	nilValue interface{} = nil

	intValue = 1

	boolValue = false

	mapString = map[string]string{
		"test":  "test",
		"key":   "value",
		"empty": "",
		"":      "",
	}

	structValue = struct {
		Test       string
		Value      int
		unexported bool
	}{
		Test:       "testing",
		Value:      123,
		unexported: false,
	}
)

func marshalTestingUtil(t *testing.T, value interface{}, expected string, datatype string) {
	content, err := Marshal(value, datatype)
	assert.Equal(t, expected, string(content))
	assert.Nil(t, err)
}

func TestMarshal(t *testing.T) {
	marshalTestingUtil(t, nilValue, nilValueExpectedJSON, "json")
	marshalTestingUtil(t, nilValue, nilValueExpectedYaml, "yaml")
	marshalTestingUtil(t, nilValue, nilValueExpectedYaml, "yml")
	marshalTestingUtil(t, nilValue, nilValueExpectedJSON, ".json")
	marshalTestingUtil(t, nilValue, nilValueExpectedYaml, ".yaml")
	marshalTestingUtil(t, nilValue, nilValueExpectedYaml, ".yml")

	marshalTestingUtil(t, intValue, intValueExpectedJSON, "json")
	marshalTestingUtil(t, intValue, intValueExpectedYaml, "yaml")
	marshalTestingUtil(t, intValue, intValueExpectedYaml, "yml")
	marshalTestingUtil(t, intValue, intValueExpectedJSON, ".json")
	marshalTestingUtil(t, intValue, intValueExpectedYaml, ".yaml")
	marshalTestingUtil(t, intValue, intValueExpectedYaml, ".yml")

	marshalTestingUtil(t, boolValue, boolValueExpectedJSON, "json")
	marshalTestingUtil(t, boolValue, boolValueExpectedYaml, "yaml")
	marshalTestingUtil(t, boolValue, boolValueExpectedYaml, "yml")
	marshalTestingUtil(t, boolValue, boolValueExpectedJSON, ".json")
	marshalTestingUtil(t, boolValue, boolValueExpectedYaml, ".yaml")
	marshalTestingUtil(t, boolValue, boolValueExpectedYaml, ".yml")

	marshalTestingUtil(t, mapString, mapStringExpectedJSON, "json")
	marshalTestingUtil(t, mapString, mapStringExpectedYaml, "yaml")
	marshalTestingUtil(t, mapString, mapStringExpectedYaml, "yml")
	marshalTestingUtil(t, mapString, mapStringExpectedJSON, ".json")
	marshalTestingUtil(t, mapString, mapStringExpectedYaml, ".yaml")
	marshalTestingUtil(t, mapString, mapStringExpectedYaml, ".yml")

	marshalTestingUtil(t, structValue, structValueExpectedJSON, "json")
	marshalTestingUtil(t, structValue, structValueExpectedYaml, "yaml")
	marshalTestingUtil(t, structValue, structValueExpectedYaml, "yml")
	marshalTestingUtil(t, structValue, structValueExpectedJSON, ".json")
	marshalTestingUtil(t, structValue, structValueExpectedYaml, ".yaml")
	marshalTestingUtil(t, structValue, structValueExpectedYaml, ".yml")
}
