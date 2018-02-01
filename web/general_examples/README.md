# General Web Examples

These examples showcase the nearly bare-minimum code to get a web server up and running.

There are four different ways showcased in these examples:
* Using `http.DefaultServeMux`
    * Found in `mux/default.go`
* Using `http.ServeMux`
    * Found in `mux/servemux.go`
* Using no multiplexer and analyzing the path instead
    * Found in `no-mux/main.go`
* Using `http.Server`
    * Found in `server/main.go`