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

func PrintWarning(msg string) {
	PrintColorText(&Text{
		Text:           msg,
		ColorAttribute: color.BgYellow,
		Style:          []color.Attribute{color.Bold},
	})

}
func PrintError(err error) {
	PrintColorText(&Text{
		Text:           err.Error(),
		ColorAttribute: color.FgRed,
		Style:          []color.Attribute{color.Bold},
	})
}
