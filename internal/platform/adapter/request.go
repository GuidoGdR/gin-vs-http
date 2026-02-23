package adapter

type Request struct {
	Params map[string]string // Parámetros de URL si usas Gin/Chi
	Body   []byte            // El JSON crudo
	Method string
}

func newRequest(method string, body []byte, params map[string]string) *Request {
	return &Request{
		Params: params,
		Method: method,
		Body:   body,
	}
}
