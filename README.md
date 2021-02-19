# go_web

## Intro
run `Hello World` on go web

## Server Creation
use `ListenAndServe(port, ServeMux)`

`http.HandleFunc` will transfer to a Handler with the method.
- To create with `http.Handle` need to realise `ServeHTTP` interface
    - [create handler](src/webapp/chapter02/web01/main.go)
- `http.Server` to config server
    - [create server](src/webapp/chapter02/web02/main.go)
- `http.NewServeMux` to create mux
    - [create serveMux](src/webapp/chapter02/web03/main.go)

## HTTP
`net/http`

1. Open a TCP connection
2. Send an HTTP message
3. Read the response sent by the server
4. Close or reuse the connection for further requests

http request
- [Get Message](src/webapp/chapter03_http/main.go)
- [Post Message](src/webapp/chapter03_http/index.html)

http response
```
1xx informational response – the request was received, continuing process
2xx successful – the request was successfully received, understood, and accepted
3xx redirection – further action needs to be taken in order to complete the request  
4xx client error – the request contains bad syntax or cannot be fulfilled
5xx server error – the server failed to fulfil an apparently valid request
```

[HTTP Reference](https://developer.mozilla.org/en-US/docs/Web/HTTP)

## Database
`database/sql`