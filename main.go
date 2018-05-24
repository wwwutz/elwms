package elwms

import (
	"fmt"
	"os"
	"runtime"
)

func Exiton(err error, msg string) {
	r := "panic exit."
	if err == nil {
		return
	}
	pc, file, no, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)
	d := ""
	if details != nil {
		d = " in " + details.Name() + "()"
	}
	if ok {
		r = fmt.Sprintf("// %s failed%s with err=%v\n// %s#%d", msg, d, err, file, no)
	} else {
		r = fmt.Sprintf("// %s failed%s with err=%v", msg, d, err)
	}
	fmt.Println(r)
	os.Exit(1)
}

func Exists(filename string) bool {
	if _, err := os.Stat(filename); err == nil {
		fmt.Println("### EXISTS: " + filename + " exists.")
		return true
	}
	return false
}
