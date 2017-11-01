package server

// #cgo LDFLAGS: -landroid -llog
// #include <android/log.h>
// #include <string.h>
import "C"
import (
	"log"
	"unsafe"
	"os"
	"bufio"
)

var (
	ctag = C.CString("go-server")
)

type infoWriter struct{}

func (infoWriter) Write(p []byte) (n int, err error) {
	cstr := C.CString(string(p))
	C.__android_log_write(C.ANDROID_LOG_INFO, ctag, cstr)
	C.free(unsafe.Pointer(cstr))
	return len(p), nil
}

func lineLog(f *os.File, priority C.int) {
	const logSize = 1024 // matches android/log.h.
	r := bufio.NewReaderSize(f, logSize)
	for {
		line, _, err := r.ReadLine()
		str := string(line)
		if err != nil {
			str += " " + err.Error()
		}
		cstr := C.CString(str)
		C.__android_log_write(priority, ctag, cstr)
		C.free(unsafe.Pointer(cstr))
		if err != nil {
			break
		}
	}
}

func init() {
	log.SetOutput(infoWriter{})

	// android logcat includes all of log.LstdFlags
	log.SetFlags(log.Flags() &^ log.LstdFlags)

	// handle stdout & stderr
	r, w, err := os.Pipe()
	if err != nil {
		panic(err)
	}
	os.Stderr = w
	go lineLog(r, C.ANDROID_LOG_ERROR)

	r, w, err = os.Pipe()
	if err != nil {
		panic(err)
	}
	os.Stdout = w
	go lineLog(r, C.ANDROID_LOG_INFO)
}
