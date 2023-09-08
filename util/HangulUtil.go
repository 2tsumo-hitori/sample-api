package util

import "strings"

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
