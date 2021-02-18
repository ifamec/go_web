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
