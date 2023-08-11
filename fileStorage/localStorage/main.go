package localStorage

import "net/http"

func RunHttpFileServer(path string) error {

	return http.ListenAndServe("0.0.0.0:9090",
		http.FileServer(http.Dir(path)))
}
