package fileutil

import (
	"io/ioutil"

	"github.com/pkg/errors"
)

// Utility constants for managing files
const (
	DefaultFilePerm = 0644
	DefaultDirPerm  = 0755
)

// Copy copies a file from a source to a destination
func Copy(source string, destination string) error {
	content, err := ioutil.ReadFile(source)

	if err != nil {
		return errors.Wrapf(err, "Error reading file %s", source)
	}

	err = ioutil.WriteFile(destination, content, DefaultFilePerm)

	if err != nil {
		return errors.Wrapf(err, "Error writing file %s", destination)
	}

	return nil
}
