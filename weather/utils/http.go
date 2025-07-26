package utils

import (
	"fmt"
	"net/http"
)

func Fetch(url string, headers map[string]string) (*http.Response, error) {
	method := headers["Method"]
	if method == "" {
		method = "GET"
	}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch URL %s: %w", url, err)
	}
	for key, value := range headers {
		req.Header.Set(key, value)
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch URL %s: %w", url, err)
	}
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("received non 200 response: %d", resp.StatusCode)
	}

	return resp, nil
}
