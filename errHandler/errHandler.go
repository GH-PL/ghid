package errHandler

import (
	"errors"
	"fmt"
	"ghid/output"
	"log"
	"os"
	"syscall"
)

type IsERROR struct {
	Err error
	Msg string
}

func IsError(err *IsERROR) {
	if os.IsPermission(err.Err) {
		output.PrintError(fmt.Errorf("Permission denied for file: %s", err.Msg))
	}
	if os.IsNotExist(err.Err) {
		output.PrintError(fmt.Errorf("The file does not exist: %s", err.Msg))
	}
	if errors.Is(err.Err, syscall.EISDIR) {
		output.PrintError(fmt.Errorf("Expected file but found directory: %s", err.Msg))
	}
	if errors.Is(err.Err, syscall.EINVAL) {
		output.PrintError(fmt.Errorf("Invalid type"))
	}
	log.Fatal("Error. App not work")
}
