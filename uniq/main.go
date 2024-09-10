package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
	uniq "uniq/stringUniq"
)

func process(input io.Reader, output io.Writer, cfg *uniq.Config) {
	var str []string
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		str = append(str, scanner.Text())
	}

	stringAfterUniq := uniq.UniqCMD(&str, cfg)
	outputString := strings.Join(stringAfterUniq, "\n")
	fmt.Println(' ')
	output.Write([]byte(outputString))
}

func main() {
	var uniqConfig uniq.Config
	flag.BoolVar(&uniqConfig.Count, "c", false, "count uniq files")
	flag.BoolVar(&uniqConfig.Duplicates, "d", false, "show only duplicates")
	flag.BoolVar(&uniqConfig.Unique, "u", false, "show only unique files")
	flag.IntVar(&uniqConfig.SkipFields, "f", 0, "do not count first f fields")
	flag.IntVar(&uniqConfig.SkipChars, "s", 0, "do not count first s chars")

	flag.Parse()

	process(os.Stdin, os.Stdout, &uniqConfig)
	return
}
