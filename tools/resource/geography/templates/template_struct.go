// Package templates contains utilities for working with available templates
package templates

import (
	"bytes"
	"errors"
	"generator/backend-go/tools/resource"
	"generator/backend-go/tools/resource/geography/templates/templateUtils"
	"os"
	"text/template"
)

var ErrPayload = errors.New("template playload is not set")

//Template struct holds information about template
type Template struct {
	resource.Resource
	Payload   *template.Template
	Variables templateUtils.TemplateVariables
}

//NewTemplateVariables returns new Template
func New(res resource.Resource, tpl *template.Template, rnd templateUtils.TemplateVariables) Template {
	return Template{
		Resource:  res,
		Payload:   tpl,
		Variables: rnd,
	}
}

//Render returns rendered template as string
func (t Template) render() (string, error) {
	var tpl bytes.Buffer

	if t.Payload == nil {
		return "", ErrPayload
	}

	err := t.Payload.Execute(&tpl, t.Variables)

	return tpl.String(), err
}

//RenderAndWrite renders given template and put it in given place in filesystem.
func (t Template) RenderAndWrite() error {
	renderedTemplate, err := t.render()

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

	_, err = t.Write([]byte(renderedTemplate))

	return err
}