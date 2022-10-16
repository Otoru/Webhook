package files

type Metadata struct {
	Name        string            `json:"name"`
	Description string            `json:"description,omitempty"`
	Labels      map[string]string `json:"labels,omitempty"`
}

type Specification struct {
	Path       string                   `json:"path,omitempty"`
	ApiVersion string                   `json:"apiVersion"`
	Kind       string                   `json:"kind"`
	Metadata   *Metadata                `json:"metadata"`
	Specs      []map[string]interface{} `json:"specs"`
}

func (spec *Specification) Validate() error {
	return nil
}
