package main

import "fmt"

// Product: The object being built
type Request struct {
    method string
    url    string
}

// Builder: Defines the step-by-step construction
type RequestBuilder struct {
    method string
    url    string
}

func (rb *RequestBuilder) SetMethod(method string) {
    rb.method = method
}

func (rb *RequestBuilder) SetURL(url string) {
    rb.url = url
}

func (rb *RequestBuilder) Build() Request {
    return Request{
        method: rb.method,
        url:    rb.url,
    }
}

func main() {
    builder := RequestBuilder{}  // Create a builder instance

    builder.SetMethod("GET")     // Configure the request
    builder.SetURL("https://example.com")
    req := builder.Build()       // Construct the final object
    fmt.Println(req.method, req.url) // Output: GET https://example.com
}
