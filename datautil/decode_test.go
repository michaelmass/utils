package datautil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	mapStringExpectedJSONDecode = map[string]interface{}(map[string]interface{}{"": "", "empty": "", "key": "value", "test": "test"})
	mapStringExpectedYamlDecode = map[interface{}]interface{}(map[interface{}]interface{}{"": "", "empty": "", "key": "value", "test": "test"})
)

func unmarshalTestingUtil(t *testing.T, value string, expected interface{}, datatype string) {
	var data interface{}
	err := Unmarshal([]byte(value), datatype, &data)
	assert.Equal(t, expected, data)
	assert.Nil(t, err)
}

func unmarshalTestingUtilErr(t *testing.T, value string, expected interface{}, datatype string) {
	var data interface{}
	err := Unmarshal([]byte(value), datatype, &data)
	assert.Equal(t, expected, data)
	assert.NotNil(t, err)
}

func TestDecode(t *testing.T) {
	unmarshalTestingUtil(t, mapStringExpectedJSON, mapStringExpectedJSONDecode, "json")
	unmarshalTestingUtil(t, mapStringExpectedYaml, mapStringExpectedYamlDecode, "yaml")
	unmarshalTestingUtil(t, mapStringExpectedJSON, mapStringExpectedJSONDecode, ".json")
	unmarshalTestingUtil(t, mapStringExpectedYaml, mapStringExpectedYamlDecode, ".yaml")

	unmarshalTestingUtil(t, boolValueExpectedJSON, boolValue, "json")
	unmarshalTestingUtil(t, boolValueExpectedYaml, boolValue, "yaml")
	unmarshalTestingUtil(t, boolValueExpectedJSON, boolValue, ".json")
	unmarshalTestingUtil(t, boolValueExpectedYaml, boolValue, ".yaml")
}
