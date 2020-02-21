package rxnav

import (
	"context"
	"encoding/xml"
	"fmt"
	"net/http"
	"net/url"
)

// HTTPStatusError represents an HTTP status code that wasn't a success code.
type HTTPStatusError struct {
	statusCode int
	statusName string
	err        error
}

func (e *HTTPStatusError) Code() int {
	return e.statusCode
}

func (e *HTTPStatusError) CodeString() string {
	return e.statusName
}

func (e *HTTPStatusError) Error() string {
	return fmt.Sprintf("rxnav: received bad HTTP status from RxNav: %d %s", e.statusCode, e.statusName)
}

func (e *HTTPStatusError) Unwrap() error {
	return e.err
}

var apiBaseURL = urlMustParse("https://rxnav.nlm.nih.gov/REST/")

func get(ctx context.Context, path string, accept string) (*http.Response, error) {
	newURL, err := apiBaseURL.Parse(path)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, newURL.String(), nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", accept)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func getXML(ctx context.Context, path string, v interface{}) error {
	resp, err := get(ctx, path, "application/xml, text/xml")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		return &HTTPStatusError{
			statusCode: resp.StatusCode,
			statusName: resp.Status,
		}
	}

	xmlDecoder := xml.NewDecoder(resp.Body)

	err = xmlDecoder.Decode(v)
	if err != nil {
		return err
	}

	return nil
}

func urlMustParse(rawURL string) *url.URL {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		panic(err)
	}

	return parsedURL
}
