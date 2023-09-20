package util

import (
	"bytes"
	"fmt"
	"golang.org/x/text/encoding/korean"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"unicode"
)

const (
	pythonPath = "C:/Users/danawa/AppData/Local/Programs/Python/Python311/python.exe"
	pythonFile = "/unicode.py"
)

var hangulMap = map[string]string{
	" ": " ",
	"ㄱ": "ㄱ",
	"ㄴ": "ㄴ",
	"ㄷ": "ㄷ",
	"ㄹ": "ㄹ",
	"ㅁ": "ㅁ",
	"ㅂ": "ㅂ",
	"ㅅ": "ㅅ",
	"ㅇ": "ㅇ",
	"ㅈ": "ㅈ",
	"ㅊ": "ㅊ",
	"ㅋ": "ㅋ",
	"ㅌ": "ㅌ",
	"ㅍ": "ㅍ",
	"ㅎ": "ㅎ",
	"ㄲ": "ㄲ",
	"ㄸ": "ㄸ",
	"ㅃ": "ㅃ",
	"ㅆ": "ㅆ",
	"ㅉ": "ㅉ",
	"ㄳ": "ㄱㅅ",
	"ㄵ": "ㄴㅈ",
	"ㄶ": "ㄴㅎ",
	"ㄺ": "ㄹㄱ",
	"ㄻ": "ㄹㅁ",
	"ㄼ": "ㄹㅂ",
	"ㄽ": "ㄹㅅ",
	"ㄾ": "ㄹㅌ",
	"ㄿ": "ㄹㅍ",
	"ㅀ": "ㄹㅎ",
	"ㅄ": "ㅂㅅ",
}

func InspectSpell(keyword string) (string, string) {
	spells := strings.Split(keyword, "")

	for _, spell := range spells {
		if _, exists := hangulMap[spell]; !exists {
			return "word", keyword
		}
	}

	return "chosung", FixSpell(spells)
}

func FixSpell(spells []string) string {
	var sb strings.Builder

	for i := 0; i < len(spells); i++ {
		sb.WriteString(hangulMap[spells[i]])
	}

	return sb.String()
}

func CombineSplitWords(spell *string) {
	path, _ := os.Getwd()

	cmd := exec.Command(pythonPath, path+pythonFile, *spell)

	output, err := cmd.CombinedOutput()

	if err != nil {
		panic(err)
	}

	decoder := korean.EUCKR.NewDecoder()
	reader := transform.NewReader(bytes.NewReader(output), decoder)
	decodedBytes, err := ioutil.ReadAll(reader)

	if err != nil {
		fmt.Println("오류:", err)
		return
	}

	*spell = string(decodedBytes)
}

func NormalizeUniCode(suggestKeyword *string) {

	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	result, _, err := transform.String(t, *suggestKeyword)

	if err != nil {
		panic(err)
	}

	*suggestKeyword = result
}
