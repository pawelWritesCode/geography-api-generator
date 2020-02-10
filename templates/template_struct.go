package templates

import (
	"bytes"
	"generator/backend-go/generators"
	"io/ioutil"
	"os"
	"path"
	"text/template"
)

//Template struct holds information about template
type Template struct {
	Resource
	Payload   *template.Template
	Variables generators.RandomVariables
}

//Render returns rendered template as string
func (t Template) Render() (string, error) {
	var tpl bytes.Buffer
	err := t.Payload.Execute(&tpl, t.Variables)

	return tpl.String(), err
}

//RenderAndEmplace renders given template and put it in given place in filesystem.
func (t Template) RenderAndEmplace() error {
	renderedTemplate, err := t.Render()

	if err != nil {
		return err
	}

	if !t.DirExist() {
		fileMode := os.FileMode(0741)
		err = os.MkdirAll(t.Directory, fileMode)

		if err != nil {
			return err
		}
	}

	err = ioutil.WriteFile(path.Clean(t.Directory+t.FileName),
		[]byte(renderedTemplate),
		0741)

	return err
}
