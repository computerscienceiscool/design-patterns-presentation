package builder

import "fmt"

type RequestBuilder struct {
	method string
	url    string
}

func (rb *RequestBuilder) SetMethod(method string) *RequestBuilder {
	rb.method = method
	return rb
}

func (rb *RequestBuilder) SetURL(url string) *RequestBuilder {
	rb.url = url
	return rb
}

func (rb *RequestBuilder) Build() string {
	return fmt.Sprintf("%s %s", rb.method, rb.url)
}

func DemoBuilder() {
	req := &RequestBuilder{}
	req.SetMethod("GET")
	req.SetURL("https://example.com")
	fmt.Println(req.Build())
}
