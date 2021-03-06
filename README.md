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

[db](src/webapp/chapter04_db/utils/db.go)

## Request

`net/http`

[request](src/webapp/chapter05_request/main.go)

## Template

`text/template`

`html/template`

[template](src/webapp/chapter06_template/main.go)

### Handle Static Source

`StripPrefix` - remove prefix and handle use certain handler

`FileServer`, `FileSystem`, `Dir`

### Action

[action](src/webapp/chapter06_template/main.go)
```go
{{if arg}}
    // actions
{{end}}

{{if arg}}
    // action
{{else}}
    // action
{{end}}
```

```go
{{range .}}
traverse {{.}}
{{end}}


{{range .}}
traverse {{.}}
{{else}}
nothing
{{end}}

{{range $k, $v := .}}
{{end}}
```

```go
{{with arg}}
New Value is {{.}}
{{end}}

{{with arg}}
New Value is {{.}}
{{else}}
No New Value {{.}}
{{end}}
```

```go
{{template "name"}}

{{template "name" arg}}
```

```go
{{define "model"}}
    {{template "content"}}
{{end}}
```

```go
{{block arg}}
Show if not find
{{end}}
```

## Session

### Cookie
`net/http`

- Pros
  - Ads
  - Keep Login
- Cons
  - Size Limit
  - Privacy

### Session

1. Server side session w UUID
2. Cookie to save UUID
3. Use UUID Cookie get the session




