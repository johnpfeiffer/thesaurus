package words

var TwoLetterWords = []string{
	"aa", "ab", "ad", "ae", "ag", "ah", "ai", "al", "am", "an", "ar", "as", "at", "aw", "ax", "ay", "ba", "be", "bi", "bo", "by",
	"de", "do", "ed", "ef", "eh", "el", "em", "en", "er", "es", "et", "ex", "fa", "go", "ha", "he", "hi", "hm", "ho",
	"id", "if", "in", "is", "it", "jo", "ka", "la", "li", "lo", "ma", "me", "mi", "mm", "mo", "mu", "my", "na", "ne",
	"no", "nu", "od", "oe", "of", "oh", "oi", "om", "on", "op", "or", "os", "ow", "ox", "oy", "pa", "pe", "pi", "re",
	"sh", "si", "so", "ta", "ti", "to", "uh", "um", "un", "up", "us", "ut", "we", "wo", "xi", "xu", "ya", "ye", "yo",
}

var SynonymMap = map[string]string{
	"yes":       "si",
	"ourselves": "us",
	"we":        "us",
	"is":        "be",
	"are":       "be",
	"am":        "be",
}

func IsSynonymOfTwoLetterWord(word string) (string, bool) {
	if synonym, ok := SynonymMap[word]; ok {
		return synonym, true
	}
	return "", false
}
