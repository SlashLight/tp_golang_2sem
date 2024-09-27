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
}

func getNewString(s string, conf *Config) string {
	fields, chars, reg := conf.SkipFields, conf.SkipChars, conf.Register

	listOfFields := strings.Fields(s)
	if len(listOfFields) <= fields {
		return ""
	}
	newStringIndex := strings.Index(s, listOfFields[fields])
	stringWithoutFields := s[newStringIndex:]
	if reg {
		stringWithoutFields = strings.ToLower(stringWithoutFields)
	}
	return stringWithoutFields[chars:]
}

func UniqCMD(s []string, conf *Config) []string {
	switch {
	case conf.Count:
		return count(s, conf)
	case conf.Duplicates:
		return duplicates(s, conf)
	case conf.Unique:
		return unique(s, conf)
	default:
		return uniq(s, conf)
	}
}

func uniq(s []string, conf *Config) []string {
	var (
		prev string
		ans  []string
	)

	for _, elem := range s {
		str := getNewString(elem, conf)
		if prev == str {
			continue
		}
		ans = append(ans, strings.Trim(elem, " "))
		prev = str
	}

	return ans
}

func count(s []string, conf *Config) []string {
	var (
		prev    = getNewString(s[0], conf)
		prevIdx = 0
		ans     []string
		counter int
	)

	for idx, elem := range s {
		str := getNewString(elem, conf)
		if prev == str {
			counter++
			continue
		}
		ans = append(ans, strings.Trim(strconv.Itoa(counter)+" "+s[prevIdx], " "))
		counter = 1
		prev = str
		prevIdx = idx
	}
	ans = append(ans, strings.Trim(strconv.Itoa(counter)+" "+s[prevIdx], " "))
	return ans
}

func duplicates(s []string, conf *Config) []string {
	var (
		prev          string
		prevIdx       int
		ans           []string
		duplicateFlag bool
	)

	for idx, elem := range s {
		str := getNewString(elem, conf)
		if prev == str {
			if duplicateFlag {
				continue
			}
			duplicateFlag = true
			ans = append(ans, strings.Trim(s[prevIdx], " "))
		} else {
			duplicateFlag = false
			prev = str
			prevIdx = idx
		}
	}

	return ans
}

func unique(s []string, conf *Config) []string {
	var (
		prev     string
		ans      []string
		uniqFlag = true
	)

	for idx, elem := range s {
		str := getNewString(elem, conf)
		if prev == str {
			uniqFlag = true
		} else if uniqFlag {
			uniqFlag = false
			prev = str
		} else {
			ans = append(ans, strings.Trim(s[idx-1], " "))
			prev = str
		}
	}

	if !uniqFlag {
		ans = append(ans, strings.Trim(s[len(s)-1], " "))
	}

	return ans
}
