package typeutil

import "github.com/pkg/errors"

// StringFromMapI returns a string from a map interface
func StringFromMapI(key string, values map[string]interface{}) (string, error) {
	valueInt, ok := values[key]

	if !ok {
		return "", errors.Errorf("Error value not found for key %s", key)
	}

	value, ok := valueInt.(string)

	if !ok {
		return "", errors.Errorf("Error invalid value isn't a string for key %s", key)
	}

	return value, nil
}

// IntFromMapI returns an int from a map interface
func IntFromMapI(key string, values map[string]interface{}) (int, error) {
	valueInt, ok := values[key]

	if !ok {
		return 0, errors.Errorf("Error value not found for key %s", key)
	}

	value, ok := valueInt.(int)

	if !ok {
		return 0, errors.Errorf("Error invalid value isn't an int for key %s", key)
	}

	return value, nil
}

// BoolFromMapI returns a bool from a map interface
func BoolFromMapI(key string, values map[string]interface{}) (bool, error) {
	valueInt, ok := values[key]

	if !ok {
		return false, errors.Errorf("Error value not found for key %s", key)
	}

	value, ok := valueInt.(bool)

	if !ok {
		return false, errors.Errorf("Error invalid value isn't a bool for key %s", key)
	}

	return value, nil
}

// Float64FromMapI returns a float64 from a map interface
func Float64FromMapI(key string, values map[string]interface{}) (float64, error) {
	valueInt, ok := values[key]

	if !ok {
		return 0, errors.Errorf("Error value not found for key %s", key)
	}

	value, ok := valueInt.(float64)

	if !ok {
		return 0, errors.Errorf("Error invalid value isn't an float64 for key %s", key)
	}

	return value, nil
}
