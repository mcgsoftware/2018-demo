package helpers

import (
	"github.com/google/uuid"
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

