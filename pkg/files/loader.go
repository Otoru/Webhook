package files

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/go-yaml/yaml"
	"github.com/spf13/viper"
)

// GetSpecifications turns a list of documents into specifications
func GetSpecifications(documents []*Document) ([]*Specification, error) {
	result := make([]*Specification, 0)

	for _, document := range documents {
		spec := new(Specification)
		spec.Path = document.Path

		if err := yaml.Unmarshal(document.Raw, spec); err != nil {
			return result, err
		}

		result = append(result, spec)
	}

	return result, nil
}

// GetDocuments splits received yaml's into a list of documents
func GetDocuments(files []string) ([]*Document, error) {
	result := make([]*Document, 0)

	for _, file := range files {
		content, err := os.ReadFile(file)

		if err != nil {
			return result, err
		}

		reader := bytes.NewReader(content)
		decoder := yaml.NewDecoder(reader)

		for {
			var value interface{}
			err := decoder.Decode(&value)

			if err == io.EOF {
				break
			}

			if err != nil {
				return nil, err
			}

			raw, err := yaml.Marshal(value)

			if err != nil {
				return nil, err
			}

			instance := &Document{Path: file, Raw: raw}
			result = append(result, instance)
		}
	}

	return result, nil
}

// GetYamlFiles returns the list with all yaml's present in the workdir
func GetYamlFiles() ([]string, error) {
	result := make([]string, 0)

	workdir := viper.GetString("workdir")

	if _, err := os.Stat(workdir); err != nil {
		return result, errors.New(fmt.Sprintf("Directory '%s' not found.\n", workdir))
	}

	handler := func(path string, spec os.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if !spec.IsDir() {
			switch filepath.Ext(path) {
			case ".yml", ".yaml":
				result = append(result, path)
				return nil
			default:
				return nil
			}
		}

		return nil
	}

	err := filepath.WalkDir(workdir, handler)

	if err != nil {
		return result, err
	}

	return result, nil
}
