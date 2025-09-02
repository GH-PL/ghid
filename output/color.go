package output

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

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

	outLine := getColor(attributes)
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

func PrintBlueText(msg string) {
	PrintColorText(&Text{
		Text:           msg,
		ColorAttribute: color.FgBlue,
		Style:          []color.Attribute{color.Bold},
	})

}
func PrintGreenText(msg string) {
	PrintColorText(&Text{
		Text:           msg,
		ColorAttribute: color.FgGreen,
		Style:          []color.Attribute{color.Bold},
	})

}

var colorCache = make(map[string]*color.Color)

func getColor(attribute []color.Attribute) *color.Color {
	key := makeKey(attribute)
	if cache, ok := colorCache[key]; ok {
		return cache
	}
	cache := color.New(attribute...)
	colorCache[key] = cache
	return cache
}

func makeKey(attribute []color.Attribute) string {
	attr := make([]color.Attribute, len(attribute))
	copy(attr, attribute)

	sort.Slice(attr, func(i, j int) bool {
		return attr[i] < attr[j]
	})

	var key strings.Builder
	key.Grow(len(attr)*4 - 1)
	for i, a := range attr {
		if i > 0 {
			key.WriteByte('-')
		}
		key.WriteString(strconv.Itoa(int(a)))
	}
	return key.String()
}
