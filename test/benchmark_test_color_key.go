package test

import (
	"sort"
	"strconv"
	"strings"
	"testing"

	"github.com/fatih/color"
)

func newMakeKey(attribute []color.Attribute) string {
	attr := make([]color.Attribute, len(attribute))
	copy(attr, attribute)

	sort.Slice(attr, func(i, j int) bool {
		return attr[i] < attr[j]
	})

	var parts []string
	for _, a := range attr {
		parts = append(parts, strconv.Itoa(int(a)))
	}
	return strings.Join(parts, "-")
}

var testAttributes = []color.Attribute{
	color.FgBlue,
	color.Bold,
	color.BgYellow,
	color.FgRed,
}

func BenchmarkNewMakeKey(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = newMakeKey(testAttributes)
	}
}
