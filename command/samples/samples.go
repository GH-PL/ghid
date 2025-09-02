package samples

import (
	"strings"
	"sync"

	"ghid/data"
	"ghid/errHandler"
	"ghid/flags"
	"ghid/output"
	"ghid/utils"

	"github.com/spf13/cobra"
)

func Commands() []*cobra.Command {

	var samplesCmd = &cobra.Command{
		Use:           "samples [NAME HASH]",
		Short:         "Display hash samples for the given type",
		Long:          "Display hash samples for the given type",
		SilenceErrors: true,
		SilenceUsage:  true,
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) < 1 {
				output.PrintError(errHandler.ErrEmptyArgument)
				cmd.Usage()
			} else {
				loadSamplesIndex()
				samples(args[0])
			}
		},
	}

	flags.AddCommandFlags(samplesCmd, data.CMD_SAMPLES)

	return []*cobra.Command{samplesCmd}
}

var (
	cacheOne     sync.Once
	samplesIndex map[string][]string
)

func loadSamplesIndex() {
	cacheOne.Do(func() {
		hash := utils.ParseJson(data.WAY_DATA_JSON)
		samplesIndex = make(map[string][]string)
		for _, hashValue := range hash {
			for _, mode := range hashValue.Modes {
				samplesIndex[strings.ToUpper(mode.Name)] = mode.Samples
			}
		}
	})
}

func samples(str string) {

	key := strings.ToUpper(str)
	sample, ok := samplesIndex[key]

	if !ok {
		output.PrintError(errHandler.ErrNotFoundName)
		return
	}
	if len(sample) == 0 {
		output.PrintError(errHandler.ErrNotExampleFound)
		return
	}

	var out strings.Builder
	for _, sampl := range sample {
		out.WriteString(sampl + "\n")
	}
	output.PrintBlueText(out.String())
}
