package stringUniq

import "strconv"

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

func UniqCMD(s *[]string, conf *Config) []string {
	if conf.Count {
		return count(s)
	}
	if conf.Duplicates {
		return duplicates(s)
	}
	if conf.Unique {
		return unique(s)
	}
	return uniq(s)
}

func uniq(s *[]string) []string {
	var prev string
	var ans []string

	for _, str := range *s {
		if prev == str {
			continue
		}
		ans = append(ans, str)
		prev = str
	}

	return ans
}

func count(s *[]string) []string {
	var (
		prev    = (*s)[0]
		ans     []string
		counter int
	)

	for _, str := range *s {
		if prev == str {
			counter++
			continue
		}
		ans = append(ans, strconv.Itoa(counter)+" "+prev)
		counter = 1
		prev = str
	}
	ans = append(ans, strconv.Itoa(counter)+" "+prev)
	return ans
}

func duplicates(s *[]string) []string {
	var (
		prev          string
		ans           []string
		duplicateFlag bool
	)

	for _, str := range *s {
		if prev == str {
			if duplicateFlag {
				continue
			}
			duplicateFlag = true
			ans = append(ans, str)
		} else {
			duplicateFlag = false
			prev = str
		}
	}

	return ans
}

func unique(s *[]string) []string {
	var (
		prev     string
		ans      []string
		uniqFlag = true
	)

	for _, str := range *s {
		if prev == str {
			uniqFlag = true
		} else if uniqFlag {
			uniqFlag = false
			prev = str
		} else {
			ans = append(ans, prev)
			prev = str
		}
	}

	return ans
}
