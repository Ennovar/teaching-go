# http.Server

The main advantage of using `http.Server` as a means of a web server is the configuration that can't be changed when using `http.ListenAndServe`. For a full list of configurable fields on the `http.Server` struct click [click here](https://golang.org/pkg/net/http/#Server).

When utilizing a server you can still use multiplexers. The `http.Server.Handler` field of the struct can take a multiplexer, as a multiplexer implements the `ServeHTTP` method. Just like when setting the handler argument of `http.ListenAndServe` to nil, if you set the `http.Server.Handler` field to nil it will utilize the `http.DefaultServeMux`.