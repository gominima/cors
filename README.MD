<p align="center">
  <a href="https://gominima.studio">
  <img alt="Minima" src="./assets/logo.png" />
</a>
</p>

### This is package is wrapper based on [rs/cors](https://github.com/rs/cors) package made for minima.

___

# Geting Started
### Install the package using `go get github.com/gominima/cors` and call it in your main function
<br>

```go
package main

import (
	"github.com/gominima/cors"
	"github.com/gominima/minima"
)

func main() {
	app := minima.New()
	crs := cors.New()
	app.Get("/", func(res *minima.Response, req *minima.Request) {
		res.OK().Send("Hello World")
		res.CloseConn()
	})
	app.Use(crs.AllowAll())
	app.Listen(":3000")
}
```
### Try curling the localhost the result would be something like this

<br>

```sh
$ curl -D - -H 'Origin: http://abc.com' http://localhost:3000/
HTTP/1.1 200 OK
Access-Control-Allow-Origin: *
Vary: Origin
Date: Wed, 09 Feb 2022 03:42:42 GMT
Content-Length: 11
Content-Type: text/plain; charset=utf-8
```
## Custom Cors

### Using your own custom cors config is as simple as it gets thanks to [rs/cors](https://github.com/rs/cors)
<br>

```go
app := minima.New()
crs := cors.New()
c := crs.NewCors(cors.Options{
    AllowedOrigins: []string{"http://ur_url.com"},
    AllowCredentials: true,
    // Enable Debugging for testing, consider disabling in production
    Debug: true,
})
app.Use(c())
```
