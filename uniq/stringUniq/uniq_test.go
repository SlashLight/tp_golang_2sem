package stringUniq

import (
	"reflect"
	"testing"
)

type testString struct {
	inputSlice  []string
	inputStr    string
	inputConfig Config
	resultSlice []string
	resultStr   string
}

func TestGetNewString(t *testing.T) {
	inputData := []testString{
		{
			inputStr:    "We love music.",
			inputConfig: Config{SkipFields: 1},
			resultStr:   "love music.",
		},
		{
			inputStr:    "We love music.",
			inputConfig: Config{SkipChars: 1},
			resultStr:   "e love music.",
		},
		{
			inputStr:    "We lOvE mUsiC.",
			inputConfig: Config{Register: true},
			resultStr:   "we love music.",
		},
		{
			inputStr: "We love music.",
			inputConfig: Config{
				SkipFields: 1,
				SkipChars:  1,
			},
			resultStr: "ove music.",
		},
		{
			inputStr: "We lOvE mUsiC.",
			inputConfig: Config{
				SkipFields: 1,
				SkipChars:  1,
				Register:   true,
			},
			resultStr: "ove music.",
		},
	}

	for _, data := range inputData {
		result := getNewString(data.inputStr, &data.inputConfig)
		if result != data.resultStr {
			t.Errorf("getNewString returned %s, expected %s", result, data.resultStr)
		}
	}
}

func TestUniq(t *testing.T) {
	inputData := []testString{
		{
			inputSlice: []string{
				"I love music.",
				"I love music.",
				"I love music.",
				"",
				"I love music of Katrik",
				"I love music of Katrik",
				"Thanks",
				"I love music of Katrik",
				"I love music of Katrik",
			},
			resultSlice: []string{
				"I love music.",
				"",
				"I love music of Katrik",
				"Thanks",
				"I love music of Katrik",
			},
		},
		{
			inputSlice: []string{
				"We love music.",
				"I love music.",
				"They love music.",
				"",
				"I love music of Katrik",
				"We love music of Katrik",
				"Thanks",
			},
			inputConfig: Config{SkipFields: 1},
			resultSlice: []string{
				"We love music.",
				"",
				"I love music of Katrik",
				"Thanks",
			},
		},
		{
			inputSlice: []string{
				"I love music.",
				"A love music.",
				"C love music.",
				"",
				"I love music of Katrik",
				"We love music of Katrik",
				"Thanks",
			},
			inputConfig: Config{SkipChars: 1},
			resultSlice: []string{
				"I love music.",
				"",
				"I love music of Katrik",
				"We love music of Katrik",
				"Thanks",
			},
		},
		{
			inputSlice: []string{
				"I LOVE MUSIC.",
				"I love music.",
				"I LoVe MuSiC.",
				"",
				"I love MuSIC of Katrik",
				"I love music of Katrik",
				"Thanks",
				"I love music of Katrik",
				"I Love muSIC of Katrik",
			},
			inputConfig: Config{Register: true},
			resultSlice: []string{
				"I LOVE MUSIC.",
				"",
				"I love MuSIC of Katrik",
				"Thanks",
				"I love music of Katrik",
			},
		},
	}

	for _, data := range inputData {
		result := uniq(data.inputSlice, &data.inputConfig)
		if !reflect.DeepEqual(result, data.resultSlice) {
			t.Errorf("uniq returned %s, expected %s", result, data.resultSlice)
		}
	}
}

func TestCount(t *testing.T) {
	inputData := []testString{
		{
			inputSlice: []string{
				"I love music.",
				"I love music.",
				"I love music.",
				"",
				"I love music of Katrik",
				"I love music of Katrik",
				"Thanks",
				"I love music of Katrik",
				"I love music of Katrik",
			},
			resultSlice: []string{
				"3 I love music.",
				"1",
				"2 I love music of Katrik",
				"1 Thanks",
				"2 I love music of Katrik",
			},
		},
		{
			inputSlice: []string{
				"We love music.",
				"I love music.",
				"They love music.",
				"",
				"I love music of Katrik",
				"We love music of Katrik",
				"Thanks",
			},
			inputConfig: Config{SkipFields: 1},
			resultSlice: []string{
				"3 We love music.",
				"1",
				"2 I love music of Katrik",
				"1 Thanks",
			},
		},
		{
			inputSlice: []string{
				"I love music.",
				"A love music.",
				"C love music.",
				"",
				"I love music of Katrik",
				"We love music of Katrik",
				"Thanks",
			},
			inputConfig: Config{SkipChars: 1},
			resultSlice: []string{
				"3 I love music.",
				"1",
				"1 I love music of Katrik",
				"1 We love music of Katrik",
				"1 Thanks",
			},
		},
		{
			inputSlice: []string{
				"I LOVE MUSIC.",
				"I love music.",
				"I LoVe MuSiC.",
				"",
				"I love MuSIC of Katrik",
				"I love music of Katrik",
				"Thanks",
				"I love music of Katrik",
				"I Love muSIC of Katrik",
			},
			inputConfig: Config{Register: true},
			resultSlice: []string{
				"3 I LOVE MUSIC.",
				"1",
				"2 I love MuSIC of Katrik",
				"1 Thanks",
				"2 I love music of Katrik",
			},
		},
	}

	for _, data := range inputData {
		result := count(data.inputSlice, &data.inputConfig)
		if !reflect.DeepEqual(result, data.resultSlice) {
			t.Errorf("count returned %s, expected %s", result, data.resultSlice)
		}
	}
}

func TestDuplicates(t *testing.T) {
	inputData := []testString{
		{
			inputSlice: []string{
				"I love music.",
				"I love music.",
				"I love music.",
				"",
				"I love music of Katrik",
				"I love music of Katrik",
				"Thanks",
				"I love music of Katrik",
				"I love music of Katrik",
			},
			resultSlice: []string{
				"I love music.",
				"I love music of Katrik",
				"I love music of Katrik",
			},
		},
		{
			inputSlice: []string{
				"We love music.",
				"I love music.",
				"They love music.",
				"",
				"I love music of Katrik",
				"We love music of Katrik",
				"Thanks",
			},
			inputConfig: Config{SkipFields: 1},
			resultSlice: []string{
				"We love music.",
				"I love music of Katrik",
			},
		},
		{
			inputSlice: []string{
				"I love music.",
				"A love music.",
				"C love music.",
				"",
				"I love music of Katrik",
				"We love music of Katrik",
				"Thanks",
			},
			inputConfig: Config{SkipChars: 1},
			resultSlice: []string{
				"I love music.",
			},
		},
		{
			inputSlice: []string{
				"I LOVE MUSIC.",
				"I love music.",
				"I LoVe MuSiC.",
				"",
				"I love MuSIC of Katrik",
				"I love music of Katrik",
				"Thanks",
				"I love music of Katrik",
				"I Love muSIC of Katrik",
			},
			inputConfig: Config{Register: true},
			resultSlice: []string{
				"I LOVE MUSIC.",
				"I love MuSIC of Katrik",
				"I love music of Katrik",
			},
		},
	}

	for _, data := range inputData {
		result := duplicates(data.inputSlice, &data.inputConfig)
		if !reflect.DeepEqual(result, data.resultSlice) {
			t.Errorf("duplicates returned %s, expected %s", result, data.resultSlice)
		}
	}
}

func TestUnique(t *testing.T) {
	inputData := []testString{
		{
			inputSlice: []string{
				"I love music.",
				"I love music.",
				"I love music.",
				"",
				"I love music of Katrik",
				"I love music of Katrik",
				"Thanks",
				"I love music of Katrik",
				"I love music of Katrik",
			},
			resultSlice: []string{
				"",
				"Thanks",
			},
		},
		{
			inputSlice: []string{
				"We love music.",
				"I love music.",
				"They love music.",
				"",
				"I love music of Katrik",
				"We love music of Katrik",
				"Thanks",
			},
			inputConfig: Config{SkipFields: 1},
			resultSlice: []string{
				"",
				"Thanks",
			},
		},
		{
			inputSlice: []string{
				"I love music.",
				"A love music.",
				"C love music.",
				"",
				"I love music of Katrik",
				"We love music of Katrik",
				"Thanks",
			},
			inputConfig: Config{SkipChars: 1},
			resultSlice: []string{
				"",
				"I love music of Katrik",
				"We love music of Katrik",
				"Thanks",
			},
		},
		{
			inputSlice: []string{
				"I LOVE MUSIC.",
				"I love music.",
				"I LoVe MuSiC.",
				"",
				"I love MuSIC of Katrik",
				"I love music of Katrik",
				"Thanks",
				"I love music of Katrik",
				"I Love muSIC of Katrik",
			},
			inputConfig: Config{Register: true},
			resultSlice: []string{
				"",
				"Thanks",
			},
		},
	}

	for idx, data := range inputData {
		result := unique(data.inputSlice, &data.inputConfig)
		if !reflect.DeepEqual(result, data.resultSlice) {
			t.Errorf("unique returned %s, expected %s on test %v ", result, data.resultSlice, idx)
		}
	}
}

func TestUniqCMD(t *testing.T) {
	inputData := []testString{
		{
			inputSlice: []string{
				"I love music.",
				"I love music.",
				"I love music.",
				"",
				"I love music of Katrik",
				"I love music of Katrik",
				"Thanks",
				"I love music of Katrik",
				"I love music of Katrik",
			},
			resultSlice: []string{
				"I love music.",
				"",
				"I love music of Katrik",
				"Thanks",
				"I love music of Katrik",
			},
		},
		{
			inputSlice: []string{
				"I love music.",
				"I love music.",
				"I love music.",
				"",
				"I love music of Katrik",
				"I love music of Katrik",
				"Thanks",
				"I love music of Katrik",
				"I love music of Katrik",
			},
			inputConfig: Config{Count: true},
			resultSlice: []string{
				"3 I love music.",
				"1",
				"2 I love music of Katrik",
				"1 Thanks",
				"2 I love music of Katrik",
			},
		},
		{
			inputSlice: []string{
				"I love music.",
				"I love music.",
				"I love music.",
				"",
				"I love music of Katrik",
				"I love music of Katrik",
				"Thanks",
				"I love music of Katrik",
				"I love music of Katrik",
			},
			inputConfig: Config{Duplicates: true},
			resultSlice: []string{
				"I love music.",
				"I love music of Katrik",
				"I love music of Katrik",
			},
		},
		{
			inputSlice: []string{
				"I love music.",
				"I love music.",
				"I love music.",
				"",
				"I love music of Katrik",
				"I love music of Katrik",
				"Thanks",
				"I love music of Katrik",
				"I love music of Katrik",
			},
			inputConfig: Config{Unique: true},
			resultSlice: []string{
				"",
				"Thanks",
			},
		},
	}

	for _, data := range inputData {
		result := UniqCMD(data.inputSlice, &data.inputConfig)
		if !reflect.DeepEqual(result, data.resultSlice) {
			t.Errorf("UniqCMD returned %s, expected %s", result, data.resultSlice)
		}
	}
}
