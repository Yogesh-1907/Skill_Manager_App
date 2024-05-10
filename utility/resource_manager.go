package utility

import (
	"errors"
	"path/filepath"
	"runtime"

	"github.com/magiconair/properties"
)

var appProp = filepath.Join(resourceManager.GetProjectLocation(), "resources", "application.properties")
var msgProp = filepath.Join(resourceManager.GetProjectLocation(), "resources", "messages.properties")
var propertyFiles = []string{appProp, msgProp}
var Properties, _ = properties.LoadFiles(propertyFiles, properties.UTF8, true)

type ResourceManager struct {
}

func (ResourceManager) GetProperty(propertyKey string) (string, error) {
	propertyValue, propertyFound := Properties.Get(propertyKey)
	if !propertyFound {
		return "", errors.New(Properties.MustGet("no.such.property") + propertyKey)
	}
	return propertyValue, nil
}

func (ResourceManager) GetProjectLocation() string {
	_, filename, _, _ := runtime.Caller(0)
	filename = filepath.Join(filename, "..", "..")
	return filename
}
