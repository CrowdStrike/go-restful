package restful

import (
	"testing"
)

func TestParameter_AddExtension(t *testing.T) {
	p := &Parameter{&ParameterData{Name: "name", Description: "desc"}}
	extKey := "x-testing-add-extension"
	p.AddExtension(extKey, "better work")

	if _, ok := p.data.Extensions[extKey]; !ok {
		t.Error("AddExtension should result in value being stored within the ParameterData Extensions map")
	}
}
