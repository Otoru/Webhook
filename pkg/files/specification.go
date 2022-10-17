package files

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/otoru/webhook/internal/tools"
	v1 "github.com/otoru/webhook/pkg/providers/v1"
)

type Metadata struct {
	Name        string            `json:"name"`
	Description string            `json:"description,omitempty"`
	Labels      map[string]string `json:"labels,omitempty"`
}

type Specification struct {
	Path       string                   `json:"path,omitempty"`
	ApiVersion string                   `json:"apiVersion"`
	Kind       string                   `json:"kind"`
	Metadata   Metadata                 `json:"metadata"`
	Specs      []map[string]interface{} `json:"specs"`
}

func validateMapWithStruct(item map[string]interface{}, instance v1.Validator) error {
	payload, err := json.Marshal(item)

	if err != nil {
		return err
	}

	if err := json.Unmarshal(payload, instance); err != nil {
		return err
	}

	if err := instance.Validate(); err != nil {
		return err
	}

	return nil
}

func (spec *Specification) Validate() error {
	validApiVersions := []any{"webhook/v1"}
	validKinds := []any{"Listener", "Path", "Response", "Service"}

	if !tools.Contains(spec.Kind, validKinds) {
		return errors.New(fmt.Sprintf("The kind '%s' is invalid", spec.Kind))
	}

	if !tools.Contains(spec.ApiVersion, validApiVersions) {
		return errors.New(fmt.Sprintf("The ApiVersion '%s' is invalid", spec.ApiVersion))
	}

	if spec.Metadata.Name == "" {
		return errors.New("Every specification needs a name")
	}

	if len(spec.Specs) < 1 {
		return errors.New("Every specification that needs at least one definition")
	}

	for _, item := range spec.Specs {
		switch spec.ApiVersion {
		case "webhook/v1":
			var instance v1.Validator
			switch spec.Kind {
			case "Listener":
				instance = new(v1.Listener)
				break
			case "Path":
				instance = new(v1.Path)
				break
			case "Response":
				instance = new(v1.Response)
				break
			case "Service":
				instance = new(v1.Response)
				break
			}

			if err := validateMapWithStruct(item, instance); err != nil {
				return err
			}
		}
	}

	return nil
}
