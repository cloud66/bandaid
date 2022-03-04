package bandaid

import (
	"context"
	"net/http"
	"time"
)

func HttpGetWithTimeout(url string, timeout time.Duration, headers map[string]string) (*http.Response, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(req.Context(), timeout)
	defer cancel()

	req = req.WithContext(ctx)

	if headers != nil {
		for k, v := range headers {
			req.Header.Set(k, v)
		}
	}

	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
