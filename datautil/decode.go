package datautil

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type dataDecoderFunc func([]byte, interface{}) error

var decodeDataTypeMap = map[string]dataDecoderFunc{
	"yaml": unmarshalYamlData,
	"yml":  unmarshalYamlData,
	"json": unmarshalJSONData,
}

// UnmarshalFromFile decodes the data from a file
func UnmarshalFromFile(filename string, data interface{}) error {
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		return errors.Wrapf(err, "Error reading file %s", filename)
	}

	return Unmarshal(content, filepath.Ext(filename), data)
}

// Unmarshal decodes data from byte array
func Unmarshal(content []byte, datatype string, data interface{}) error {
	dataUnmarshalFunc, ok := decodeDataTypeMap[strings.TrimPrefix(datatype, ".")]

	if !ok {
		return errors.Errorf("Error decoding type %s isn't supported", datatype)
	}

	return dataUnmarshalFunc(content, data)
}

func unmarshalYamlData(content []byte, data interface{}) error {
	err := yaml.Unmarshal(content, data)

	if err != nil {
		return errors.Wrap(err, "Error decoding yaml content")
	}

	return nil
}

func unmarshalJSONData(content []byte, data interface{}) error {
	err := json.Unmarshal(content, data)

	if err != nil {
		return errors.Wrap(err, "Error decoding json content")
	}

	return nil
}
