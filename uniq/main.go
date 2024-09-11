package main

import (
	"bufio"
	"flag"
	uniq "github.com/SlashLight/tp_golang_2sem/stringUniq"
	"io"
	"os"
	"strings"
)

func process(input io.Reader, output io.Writer, cfg *uniq.Config) {
	var str []string
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		str = append(str, scanner.Text())
	}

	stringAfterUniq := uniq.UniqCMD(&str, cfg)
	outputString := strings.Join(stringAfterUniq, "\n")
	output.Write([]byte(outputString))
}

func getInput(inputPath string) (input *os.File) {
	input, err := os.Open(inputPath)
	if err != nil {
		panic(err)
	}
	return
}

func getOutput(outputPath string) (output *os.File) {
	output, err := os.Create(outputPath)
	if err != nil {
		panic(err)
	}
	return
}

func main() {
	var (
		uniqConfig uniq.Config
		input      = os.Stdin
		output     = os.Stdout
	)
	flag.BoolVar(&uniqConfig.Count, "c", false, "count uniq files")
	flag.BoolVar(&uniqConfig.Duplicates, "d", false, "show only duplicates")
	flag.BoolVar(&uniqConfig.Unique, "u", false, "show only unique files")
	flag.IntVar(&uniqConfig.SkipFields, "f", 0, "do not count first f fields")
	flag.IntVar(&uniqConfig.SkipChars, "s", 0, "do not count first s chars")
	flag.BoolVar(&uniqConfig.Register, "i", false, "register uniq files")

	flag.Parse()

	inAndOutArgs := flag.Args()
	if len(inAndOutArgs) > 0 {
		input = getInput(inAndOutArgs[0])
		defer input.Close()

		if len(inAndOutArgs) > 1 {
			output = getOutput(inAndOutArgs[1])
			defer output.Close()
		}
	}

	process(input, output, &uniqConfig)
	return
}
