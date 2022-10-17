package v1

type Response struct {
	Name    string
	Code    int
	Headers map[string]string
	Body    map[string]interface{}
}

func (r *Response) Validate() error {
	return nil
}
