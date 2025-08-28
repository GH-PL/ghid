package output

import (
	"fmt"

	"github.com/fatih/color"
)

type Text struct {
	Text           string
	ColorAttribute color.Attribute
	Style          []color.Attribute
}

func DisableColorOutput() {
	color.NoColor = true
}

func PrintColorText(txt *Text) {
	if color.NoColor {
		fmt.Println(txt.Text)
		return
	}
	attributes := []color.Attribute{txt.ColorAttribute}
	attributes = append(attributes, txt.Style...)

	outLine := color.New(attributes...)
	outLine.Println(txt.Text)
}
