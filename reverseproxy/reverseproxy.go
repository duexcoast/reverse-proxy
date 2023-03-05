package reverseproxy

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

type ReverseProxy struct {
	Target url.URL
	proxy  *httputil.ReverseProxy
}

// Start will listen on configured listeners
func (r *ReverseProxy) Start() error {
	r.proxy = &httputil.ReverseProxy{
		Director: r.Director(),
	}

	// hard coding port 80 for now
	srv := &http.Server{
		Addr:    ":80",
		Handler: r.proxy,
	}

	return srv.ListenAndServe()
}
