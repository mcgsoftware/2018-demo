package structlog



import (
	"fmt"
	"os"
	"runtime"
	"strings"
)





// return the source filename after the last slash
func chopPath(original string) string {
	i := strings.LastIndex(original, "/")
	if i == -1 {
		return original
	} else {
		return original[i+1:]
	}
}

// Gets place where caller is in source code. File name, line number and current func.
func GetSrcLocation() string {
	// Get location of code 1 level up the stack so we aren't getting it from here.
	function, file, line, _ := runtime.Caller(1)
	str := fmt.Sprintf("%s:%d, func: %s ", chopPath(file), line, runtime.FuncForPC(function).Name())
	return str
}

// Get stack trace and error out of Go error.
func ErrorToString(err error) string {
	return fmt.Sprintf("%+v", err)
}


// Get hostname
func GetHostname() string {
	hostname, err := os.Hostname();
	if err != nil {
		return "unknown"
	}
	return hostname
}

