# Discoveries from this little project

## Network addresses
1. The TCP network address passed to `http.ListenAndServe()` should be in the format `"host:port"`. If I omit the host and put just
`":8082"`, the server will listen to all network interfaces available on the computer.
2. In other Go projects or documentation you might sometimes see network addresses written using named ports 
like `":http"` or `":http-alt"` instead of a number. If you use a named port then Go will attempt to look up 
the relevant port number from your `/etc/services` file when starting the server, or will return an error if a match can’t be found.

