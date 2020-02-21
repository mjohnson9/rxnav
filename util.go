package rxnav

import (
	"context"
	"encoding/xml"
)

// NotConnectedError wraps any error that oocurs when doing a connection check using CheckConnection
type NotConnectedError struct {
	display string
	err     error
}

func (e *NotConnectedError) Error() string {
	errorStr := "rxnav: unable to connect: " + e.display
	if e.err != nil {
		errorStr += ": " + e.err.Error()
	}

	return errorStr
}

// Unwrap returns the error that caused this error. If this error was the root cause, Unwrap will return nil
func (e *NotConnectedError) Unwrap() error {
	return e.err
}

type versionResponse struct {
	XMLName xml.Name `xml:"rxnormdata"`
	Version string   `xml:"version"`
}

// CheckConnection tests for a connection to the API by retrieving the API version.
func CheckConnection(ctx context.Context) error {
	response := &versionResponse{}
	err := getXML(ctx, "version", response)
	if err != nil {
		return &NotConnectedError{
			display: "http request failed",
			err:     err,
		}
	}

	if response.Version == "" {
		return &NotConnectedError{
			display: "got back invalid version",
		}
	}

	return nil
}
