# BuilderHttpClient

BuilderHttpClient is a Go package that provides a simple and flexible HTTP client builder for making HTTP requests. It allows you to easily construct and customize HTTP requests using a fluent interface.

## Installation

To use BuilderHttpClient in your Go project, you can simply run the following command:

```shell
go get github.com/catnovelapi/BuilderHttpClient
```

## Usage

Import the package in your Go file:

```go
import "github.com/catnovelapi/BuilderHttpClient"
```

### Making GET Requests

To make a GET request, you can use the `Get` function:

```go
response := BuilderHttpClient.Get(url, options...)
```

Here, `url` is the URL of the endpoint you want to send the GET request to, and `options` is an optional list of request options.

### Making POST Requests

To make a POST request, you can use the `Post` function:

```go
response := BuilderHttpClient.Post(url, options...)
```

Here, `url` is the URL of the endpoint you want to send the POST request to, and `options` is an optional list of request options.

### Making PUT Requests

To make a PUT request, you can use the `Put` function:

```go
response := BuilderHttpClient.Put(url, options...)
```

Here, `url` is the URL of the endpoint you want to send the PUT request to, and `options` is an optional list of request options.

### Making DELETE Requests

To make a DELETE request, you can use the `Delete` function:

```go
response := BuilderHttpClient.Delete(url, options...)
```

Here, `url` is the URL of the endpoint you want to send the DELETE request to, and `options` is an optional list of request options.

### Making PATCH Requests

To make a PATCH request, you can use the `Patch` function:

```go
response := BuilderHttpClient.Patch(url, options...)
```

Here, `url` is the URL of the endpoint you want to send the PATCH request to, and `options` is an optional list of request options.

### Request Options

You can customize the request by providing various options. The available options are:

- `Method(method string)`: Sets the HTTP method for the request.
- `ApiPath(URL string)`: Sets the URL of the API endpoint.
- `Header(header map[string]interface{})`: Sets the headers for the request.
- `Body(dataBody interface{})`: Sets the request body data.

Here's an example that demonstrates how to use request options:

```go
options := []BuilderHttpClient.Option{
    BuilderHttpClient.Method("POST"),
    BuilderHttpClient.ApiPath("https://api.example.com/users"),
    BuilderHttpClient.Header(map[string]interface{}{
        "Content-Type": "application/json",
        "Authorization": "Bearer token",
    }),
    BuilderHttpClient.Body(map[string]interface{}{
        "name": "John Doe",
        "email": "john@example.com",
    }),
}

response := BuilderHttpClient.Post(url, options...)
```

### Response Handling

The response from the HTTP request is returned as a `ResponseInterfaceBuilder`. You can use the following methods to handle the response:

- `Code() int`: Returns the HTTP status code of the response.
- `Status() string`: Returns the HTTP status message of the response.
- `Json(v interface{}) error`: Parses the response body as JSON and stores it in the provided struct or map.
- `Text() string`: Returns the response body as a string.
- `Gjson() gjson.Result`: Parses the response body as JSON and returns a `gjson.Result` object.
- `Debug() *ResponseBuilder`: Prints debug information about the request and response.
- `Cookies() string`: Returns the cookies from the response.
- `Bytes() []byte`: Returns the response body as a byte array.
Here's an example that demonstrates how to handle the response:

```go
response := BuilderHttpClient.Get(url, options...)

fmt.Println("Status Code:", response.Code())
fmt.Println("Status:", response.Status())
fmt.Println("Response Body:", response.Text())

var data map[string]interface{}
err := response.Json(&data)
if err != nil {
    fmt.Println("Error parsing JSON:", err)
} else {
    fmt.Println("Parsed JSON:", data)
}
``` 
