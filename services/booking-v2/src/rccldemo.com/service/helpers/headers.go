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
/*
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
			 fmt.Printf("Copying header %s with value %s\n", headerKey, headerValue)
		 }
	 }
}
*/

//
// Used to copy tracing headers received from caller to an outgoing service call.
// This is for service to service calls. If we don't do this, traces break at each
// service in the call chain.
//
const HdrRequestId = 	"x-request-id"
const HdrTraceId =		"x-b3-traceid"
const HdrSpanId =		"x-b3-spanid"
const HdrParentSpanId =	"x-b3-parentspanid"
const HdrSampled =		"x-b3-sampled"
const HdrFlags =		"x-b3-flags"
const HdrSpanCxt=		"x-ot-span-context"

type TraceHeaders struct {
	RequestId 		string
	TraceId   		string
	SpanId			string
	ParentSpanId	string
	Sampled			string
	Flags			string
	SpanCxt			string

}

//
// Sets trace headers upon the given request for propagating the trace
//
func ( tHdr *TraceHeaders) SetHeaders(req *http.Request) {
	if tHdr.RequestId != "" {
		req.Header.Set(HdrRequestId, tHdr.RequestId)
	}

	if tHdr.TraceId != "" {
		req.Header.Set(HdrTraceId, tHdr.TraceId)
	}

	if tHdr.SpanId != "" {
		req.Header.Set(HdrSpanId, tHdr.SpanId)
	}

	if tHdr.ParentSpanId != "" {
		req.Header.Set(HdrParentSpanId, tHdr.ParentSpanId)
	}

	if tHdr.Sampled != "" {
		req.Header.Set(HdrSampled, tHdr.Sampled)
	}

	if tHdr.Flags != "" {
		req.Header.Set(HdrFlags, tHdr.Flags)
	}

	if tHdr.SpanCxt != "" {
		req.Header.Set(HdrSpanCxt, tHdr.SpanCxt)
	}
}

// Factor to build TraceHeaders from a HTTP request's header values
// returns nil if we didn't find any headers.
func BuildFromRequestHeader(r *http.Request) (*TraceHeaders) {

	var this *TraceHeaders = &TraceHeaders{}
	found := false

	if hdrValue := r.Header.Get(HdrRequestId); hdrValue != "" {
		this.RequestId = hdrValue
		found = true
	}

	if hdrValue := r.Header.Get(HdrTraceId); hdrValue != "" {
		this.TraceId = hdrValue
		found = true
	}

	if hdrValue := r.Header.Get(HdrSpanId); hdrValue != "" {
		this.SpanId = hdrValue
		found = true
	}

	if hdrValue := r.Header.Get(HdrParentSpanId); hdrValue != "" {
		this.ParentSpanId = hdrValue
		found = true
	}

	if hdrValue := r.Header.Get(HdrSampled); hdrValue != "" {
		this.Sampled = hdrValue
	}

	if hdrValue := r.Header.Get(HdrFlags); hdrValue != "" {
		this.Flags = hdrValue
	}

	if hdrValue := r.Header.Get(HdrSpanCxt); hdrValue != "" {
		this.SpanCxt = hdrValue
	}

	if found {
		return this
	} else {
		return nil
	}
}


