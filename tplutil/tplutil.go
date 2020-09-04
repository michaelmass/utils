package tplutil

import (
	"io"
	"io/ioutil"
	templ "text/template"

	"github.com/Masterminds/sprig"
)

var funcMap = sprig.TxtFuncMap()

// Render renders a template into io.Writer
func Render(content []byte, data interface{}, wr io.Writer) error {
	t, err := templ.New("default").Funcs(funcMap).Parse(string(content))

	if err != nil {
		return err
	}

	t.Option("missingkey=error")

	err = t.Execute(wr, data)

	if err != nil {
		return err
	}

	return nil
}

// RenderFile renders a template file into io.Writer
func RenderFile(reader io.Reader, data interface{}, wr io.Writer) error {
	content, err := ioutil.ReadAll(reader)

	if err != nil {
		return err
	}

	t, err := templ.New("default").Funcs(funcMap).Parse(string(content))

	if err != nil {
		return err
	}

	t.Option("missingkey=error")

	err = t.Execute(wr, data)

	if err != nil {
		return err
	}

	return nil
}
