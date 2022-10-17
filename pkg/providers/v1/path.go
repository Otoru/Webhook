package v1

type Path struct {
	Name string
	Path string
}

func (p *Path) Validate() error {
	return nil
}
