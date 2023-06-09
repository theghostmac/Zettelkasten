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
5. We can allow only POST (for example), using:
```go
func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POSt" {
		w.Header().Set("Allow", "POST")
		w.WrieHeader(405)
		w.Write([]byte{"method is not allowed..."})
		return
    }
	w.Write([]byte("doing this function thing now..."))
}
```

6. GO's `http.FileServer()` sanitizes all request path, and also supports [range requests](https://benramsey.com/blog/2008/05/206-partial-content-and-range-requests/).
7. The files are served from RAM and not from disc, with performance as a tradeoff.
8. 