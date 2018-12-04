package helpers

import (
	"github.com/google/uuid"
	"net/http"
	"strings"
)


// Headers required by our services
const TRACE_ID_HEADER = "X-B3-TRACEID"
const VDSID_HEADER = "VDSID"



// Generate a trace Id as UUID string with no dashes
func GenerateTraceId() string {
	uuid := uuid.New()
	nodashes := strings.Replace(uuid.String(), "-", "", -1)

	return 	nodashes

}

//
// Copies tracing headers received from caller to outgoing response.
// This is for service to service calls. Otherwise, traces break at each
// service in the call chain.
//
func CopyTracingHeaders(w http.ResponseWriter, r *http.Request) {

	// List of trace related headers
	headers := [...]string{
		"x-request-id",
		"x-b3-traceid",
		"x-b3-spanid",
		"x-b3-parentspanid",
		"x-b3-sampled",
		"x-b3-flags",
		"x-ot-span-context",
	}

	// Iterate over all trace related headers. If request has them, copy header to response.
	for _, headerKey := range headers {
		if headerValue := r.Header.Get(headerKey); headerValue != "" {
			w.Header().Set(headerKey, headerValue)
		}
	}
}