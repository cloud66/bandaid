package bandaid

import (
	"fmt"
	"net/http"
	"net/http/httputil"
)

func PrintRequest(r *http.Request) {
	dump, _ := httputil.DumpRequestOut(r, true)
	fmt.Println(string(dump))
}

func PrintResponse(r *http.Response) {
	dump, _ := httputil.DumpResponse(r, true)
	fmt.Println(string(dump))
}
