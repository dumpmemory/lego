// Code generated by 'internal/releaser'; DO NOT EDIT.

package useragent

import (
	"fmt"
	"net/http"
	"runtime"
)

const (
	// ourUserAgent is the User-Agent of this underlying library package.
	ourUserAgent = "goacme-lego/4.25.1"

	// ourUserAgentComment is part of the UA comment linked to the version status of this underlying library package.
	// values: detach|release
	// NOTE: Update this with each tagged release.
	ourUserAgentComment = "detach"
)

// Get builds and returns the User-Agent string.
func Get() string {
	return fmt.Sprintf("%s (%s; %s; %s)", ourUserAgent, ourUserAgentComment, runtime.GOOS, runtime.GOARCH)
}

// SetHeader sets the User-Agent header.
func SetHeader(h http.Header) {
	h.Set("User-Agent", Get())
}
