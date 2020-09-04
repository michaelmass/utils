package typeutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var mapInterface = map[string]interface{}{
	"mapI": map[string]interface{}{
		"key": "value",
	},
	"nil":         nil,
	"emptystring": "",
	"string":      "test",
	"booltrue":    true,
	"boolfalse":   false,
	"int":         1,
	"float":       0.3,
}

func TestStringFromMapI(t *testing.T) {
	str, err := StringFromMapI("string", mapInterface)
	assert.Equal(t, "test", str)
	assert.Nil(t, err)

	str, err = StringFromMapI("emptystring", mapInterface)
	assert.Equal(t, "", str)
	assert.Nil(t, err)

	str, err = StringFromMapI("nil", mapInterface)
	assert.Equal(t, "", str)
	assert.NotNil(t, err)

	str, err = StringFromMapI("mapI", mapInterface)
	assert.Equal(t, "", str)
	assert.NotNil(t, err)
}

func TestIntFromMapI(t *testing.T) {
	i, err := IntFromMapI("int", mapInterface)
	assert.Equal(t, 1, i)
	assert.Nil(t, err)

	i, err = IntFromMapI("float", mapInterface)
	assert.Equal(t, 0, i)
	assert.NotNil(t, err)

	i, err = IntFromMapI("nil", mapInterface)
	assert.Equal(t, 0, i)
	assert.NotNil(t, err)

	i, err = IntFromMapI("bool", mapInterface)
	assert.Equal(t, 0, i)
	assert.NotNil(t, err)
}

func TestBoolFromMapI(t *testing.T) {
	b, err := BoolFromMapI("boolfalse", mapInterface)
	assert.Equal(t, false, b)
	assert.Nil(t, err)

	b, err = BoolFromMapI("booltrue", mapInterface)
	assert.Equal(t, true, b)
	assert.Nil(t, err)

	b, err = BoolFromMapI("float", mapInterface)
	assert.Equal(t, false, b)
	assert.NotNil(t, err)

	b, err = BoolFromMapI("nil", mapInterface)
	assert.Equal(t, false, b)
	assert.NotNil(t, err)
}

func TestFloat64FromMapI(t *testing.T) {
	f, err := Float64FromMapI("float", mapInterface)
	assert.Equal(t, 0.3, f)
	assert.Nil(t, err)

	f, err = Float64FromMapI("int", mapInterface)
	assert.Equal(t, 0.0, f)
	assert.NotNil(t, err)

	f, err = Float64FromMapI("booltrue", mapInterface)
	assert.Equal(t, 0.0, f)
	assert.NotNil(t, err)

	f, err = Float64FromMapI("nil", mapInterface)
	assert.Equal(t, 0.0, f)
	assert.NotNil(t, err)
}
