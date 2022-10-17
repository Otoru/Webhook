package v1

type Listener struct {
	Name string
	Host string
	Port int
}

func (l *Listener) Validate() error {
	return nil
}
