package datautil

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/michaelmass/utils/fileutil"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type dataEncoderFunc func(interface{}) ([]byte, error)

var encodeDataTypeMap = map[string]dataEncoderFunc{
	"yaml": marshalYamlData,
	"yml":  marshalYamlData,
	"json": marshalJSONData,
}

// MarshalToFile encodes the data into a file
func MarshalToFile(data interface{}, filename string) error {
	content, err := Marshal(data, filepath.Ext(filename))

	if err != nil {
		return errors.Wrap(err, "Error encoding data")
	}

	err = ioutil.WriteFile(filename, content, fileutil.DefaultFilePerm)

	if err != nil {
		return errors.Wrapf(err, "Error creating file %s", filename)
	}

	return nil
}

// Marshal encodes data into byte array
func Marshal(data interface{}, datatype string) ([]byte, error) {
	dataMarshalFunc, ok := encodeDataTypeMap[strings.TrimPrefix(datatype, ".")]

	if !ok {
		return nil, errors.Errorf("Error encoding type %s isn't supported", datatype)
	}

	return dataMarshalFunc(data)
}

func marshalYamlData(data interface{}) ([]byte, error) {
	return yaml.Marshal(data)
}

func marshalJSONData(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}
