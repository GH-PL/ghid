package errHandler

import (
	"fmt"
	"ghid/output"
	"log"
)

type FileError struct {
	Operation string
	Path      string
	Err       error
}

func ErrorFile(operation string, path string, err error) {
	txtErr := &FileError{
		Operation: operation,
		Path:      path,
		Err:       err,
	}
	output.PrintError(txtErr.Error())
	log.Fatal("Error. App not work")
}

func (e *FileError) Error() error {
	return fmt.Errorf("error during: %s of file: %s -- %v", e.Operation, e.Path, e.Err)
}
func (e *FileError) Unwrap() error {
	return e.Err
}
