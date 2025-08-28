package errHandler

import (
	"errors"
	"fmt"
	"ghid/output"
	"log"
	"os"
	"syscall"

	"github.com/fatih/color"
)

func Show(err error) {
	output.PrintColorText(&output.Text{
		Text:           err.Error(),
		ColorAttribute: color.FgRed,
		Style:          []color.Attribute{color.Bold},
	})
}

type IsERROR struct {
	Err error
	Msg string
}

func IsError(err *IsERROR) {
	if os.IsPermission(err.Err) {
		Show(fmt.Errorf("Permission denied for file: %s", err.Msg))
	}
	if os.IsNotExist(err.Err) {
		Show(fmt.Errorf("The file does not exist: %s", err.Msg))
	}
	if errors.Is(err.Err, syscall.EISDIR) {
		Show(fmt.Errorf("Expected file but found directory: %s", err.Msg))
	}
	if errors.Is(err.Err, syscall.EINVAL) {
		Show(fmt.Errorf("Invalid type"))
	}
	log.Fatal("Error. App not work")
}
