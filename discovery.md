# Discoveries from this little project

## Network addresses
1. The TCP network address passed to `http.ListenAndServe()` should be in the format `"host:port"`. If I omit the host and put just
`":8082"`, the server will listen to all network interfaces available on the computer.
2. In other Go projects or documentation you might sometimes see network addresses written using named ports 
like `":http"` or `":http-alt"` instead of a number. If you use a named port then Go will attempt to look up 
the relevant port number from your `/etc/services` file when starting the server, or will return an error if a match can’t be found.
3. Don't use `DefaultServeMux` for serious applications, as it allows any package to access it and register a route for itself, 
leading to compromise or exposure of malicious handler for the web. Use your own locally scoped `ServeMux` instead.
4. To redirect all HTTP requests to a canonical URL or if your app is acting as backend to multiple sites/services, use host name matching:
```go
mux := http.NewServeMux()
mux.HandleFunc("foo.example.org/", fooHandler)
mux.HandleFunc("foo", fooHandler)
```