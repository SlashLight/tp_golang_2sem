package main

import (
	"bufio"
	"flag"
	uniq "github.com/SlashLight/tp_golang_2sem/stringUniq"
	"io"
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

func main() {
	var uniqConfig uniq.Config
	flag.BoolVar(&uniqConfig.Count, "c", false, "count uniq files")
	flag.BoolVar(&uniqConfig.Duplicates, "d", false, "show only duplicates")
	flag.BoolVar(&uniqConfig.Unique, "u", false, "show only unique files")
	flag.IntVar(&uniqConfig.SkipFields, "f", 0, "do not count first f fields")
	flag.IntVar(&uniqConfig.SkipChars, "c", 0, "do not count first c chars")

}
