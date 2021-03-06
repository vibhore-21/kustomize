/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package error has contextual error types.
package error

import (
	"fmt"

	yaml "k8s.io/apimachinery/pkg/util/yaml"
)

// YamlFormatError represents error with yaml file name where json/yaml format error happens.
type YamlFormatError struct {
	Path     string
	ErrorMsg string
}

func (e YamlFormatError) Error() string {
	return fmt.Sprintf("YAML file [%s] encounters a format error.\n%s\n", e.Path, e.ErrorMsg)
}

// ErrorHandler handles YamlFormatError
func ErrorHandler(e error, path string) error {
	if err, ok := e.(yaml.YAMLSyntaxError); ok {
		return YamlFormatError{
			Path:     path,
			ErrorMsg: err.Error(),
		}
	}
	return e
}
