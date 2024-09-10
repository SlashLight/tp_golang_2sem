package stringUniq

import (
	"strconv"
	"strings"
)

type Config struct {
	Count      bool
	Duplicates bool
	Unique     bool
	SkipFields int
	SkipChars  int
	Register   bool
	InputFile  *string
	OutputFile *string
}

func getNewString(s string, fields, chars int) string {
	listOfFields := strings.Fields(s)
	if len(listOfFields) <= fields {
		return ""
	}
	newStringIndex := strings.Index(s, listOfFields[fields])
	stringWithoutFields := s[newStringIndex:]
	return stringWithoutFields[chars:]
}

func UniqCMD(s *[]string, conf *Config) []string {
	if conf.Count {
		return count(s, conf.SkipFields, conf.SkipChars)
	}
	if conf.Duplicates {
		return duplicates(s, conf.SkipFields, conf.SkipChars)
	}
	if conf.Unique {
		return unique(s, conf.SkipFields, conf.SkipChars)
	}
	return uniq(s, conf.SkipFields, conf.SkipChars)
}

func uniq(s *[]string, fields, chars int) []string {
	var prev string
	var ans []string

	for _, elem := range *s {
		str := getNewString(elem, fields, chars)
		if prev == str {
			continue
		}
		ans = append(ans, elem)
		prev = str
	}

	return ans
}

func count(s *[]string, fields, chars int) []string {
	var (
		prev    = getNewString((*s)[0], fields, chars)
		prevIdx = 0
		ans     []string
		counter int
	)

	for idx, elem := range *s {
		str := getNewString(elem, fields, chars)
		if prev == str {
			counter++
			continue
		}
		ans = append(ans, strconv.Itoa(counter)+" "+(*s)[prevIdx])
		counter = 1
		prev = str
		prevIdx = idx
	}
	ans = append(ans, strconv.Itoa(counter)+" "+(*s)[prevIdx])
	return ans
}

func duplicates(s *[]string, fields, chars int) []string {
	var (
		prev          string
		prevIdx       int
		ans           []string
		duplicateFlag bool
	)

	for idx, elem := range *s {
		str := getNewString(elem, fields, chars)
		if prev == str {
			if duplicateFlag {
				continue
			}
			duplicateFlag = true
			ans = append(ans, (*s)[prevIdx])
		} else {
			duplicateFlag = false
			prev = str
			prevIdx = idx
		}
	}

	return ans
}

func unique(s *[]string, fields, chars int) []string {
	var (
		prev     string
		ans      []string
		uniqFlag = true
	)

	for idx, elem := range *s {
		str := getNewString(elem, fields, chars)
		if prev == str {
			uniqFlag = true
		} else if uniqFlag {
			uniqFlag = false
			prev = str
		} else {
			ans = append(ans, (*s)[idx-1])
			prev = str
		}
	}

	if !uniqFlag {
		ans = append(ans, (*s)[len(*s)-1])
	}

	return ans
}
