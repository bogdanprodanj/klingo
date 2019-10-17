package app

type piqad struct {
	character string
	unicode   int
}

// this map was created based on the article from https://www.kli.org/about-klingon/sounds-of-klingon
var (
	englishToPiqadMap = map[string]piqad{
		"a": {character: "a", unicode: 0xF8D0},
		"b": {character: "b", unicode: 0xF8D1},
		"c": {character: "ch", unicode: 0xF8D2},
		"d": {character: "D", unicode: 0xF8D3},
		"e": {character: "e", unicode: 0xF8D4},
		"g": {character: "gh", unicode: 0xF8D5},
		"h": {character: "H", unicode: 0xF8D6},
		"i": {character: "I", unicode: 0xF8D7},
		"j": {character: "j", unicode: 0xF8D8},
		"l": {character: "L", unicode: 0xF8D9},
		"m": {character: "m", unicode: 0xF8DA},
		"n": {character: "n", unicode: 0xF8DB},
		"f": {character: "ng", unicode: 0xF8DC},
		"o": {character: "o", unicode: 0xF8DD},
		"p": {character: "p", unicode: 0xF8DE},
		"k": {character: "q", unicode: 0xF8DF},
		"q": {character: "Q", unicode: 0xF8E0},
		"r": {character: "r", unicode: 0xF8E1},
		"s": {character: "S", unicode: 0xF8E2},
		"t": {character: "t", unicode: 0xF8E3},
		"x": {character: "tlh", unicode: 0xF8E4},
		"u": {character: "u", unicode: 0xF8E5},
		"v": {character: "v", unicode: 0xF8E6},
		"w": {character: "w", unicode: 0xF8E7},
		"y": {character: "y", unicode: 0xF8E8},
		"z": {character: "'", unicode: 0xF8E9},
		"0": {character: "0", unicode: 0xF8F0},
		"1": {character: "1", unicode: 0xF8F1},
		"2": {character: "2", unicode: 0xF8F2},
		"3": {character: "3", unicode: 0xF8F3},
		"4": {character: "4", unicode: 0xF8F4},
		"5": {character: "5", unicode: 0xF8F5},
		"6": {character: "6", unicode: 0xF8F6},
		"7": {character: "7", unicode: 0xF8F7},
		"8": {character: "8", unicode: 0xF8F8},
		"9": {character: "9", unicode: 0xF8F9},
		",": {character: ",", unicode: 0xF8FE},
		".": {character: ".", unicode: 0xF8FD},
		" ": {character: " ", unicode: 0x0020},
	}
)
