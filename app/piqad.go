package app

type piqad struct {
	character string
	unicode   int
}

// this map was created based on the article from https://www.kli.org/about-klingon/sounds-of-klingon
var (
	englishToPiqadMap = map[string]piqad{
		"a": piqad{character: "a", unicode: 0xF8D0},
		"b": piqad{character: "b", unicode: 0xF8D1},
		"c": piqad{character: "ch", unicode: 0xF8D2},
		"d": piqad{character: "D", unicode: 0xF8D3},
		"e": piqad{character: "e", unicode: 0xF8D4},
		"g": piqad{character: "gh", unicode: 0xF8D5},
		"h": piqad{character: "H", unicode: 0xF8D6},
		"i": piqad{character: "I", unicode: 0xF8D7},
		"j": piqad{character: "j", unicode: 0xF8D8},
		"l": piqad{character: "L", unicode: 0xF8D9},
		"m": piqad{character: "m", unicode: 0xF8DA},
		"n": piqad{character: "n", unicode: 0xF8DB},
		"f": piqad{character: "ng", unicode: 0xF8DC},
		"o": piqad{character: "o", unicode: 0xF8DD},
		"p": piqad{character: "p", unicode: 0xF8DE},
		"k": piqad{character: "q", unicode: 0xF8DF},
		"q": piqad{character: "Q", unicode: 0xF8E0},
		"r": piqad{character: "r", unicode: 0xF8E1},
		"s": piqad{character: "S", unicode: 0xF8E2},
		"t": piqad{character: "t", unicode: 0xF8E3},
		"x": piqad{character: "tlh", unicode: 0xF8E4},
		"u": piqad{character: "u", unicode: 0xF8E5},
		"v": piqad{character: "v", unicode: 0xF8E6},
		"w": piqad{character: "w", unicode: 0xF8E7},
		"y": piqad{character: "y", unicode: 0xF8E8},
		"z": piqad{character: "'", unicode: 0xF8E9},
		"0": piqad{character: "0", unicode: 0xF8F0},
		"1": piqad{character: "1", unicode: 0xF8F1},
		"2": piqad{character: "2", unicode: 0xF8F2},
		"3": piqad{character: "3", unicode: 0xF8F3},
		"4": piqad{character: "4", unicode: 0xF8F4},
		"5": piqad{character: "5", unicode: 0xF8F5},
		"6": piqad{character: "6", unicode: 0xF8F6},
		"7": piqad{character: "7", unicode: 0xF8F7},
		"8": piqad{character: "8", unicode: 0xF8F8},
		"9": piqad{character: "9", unicode: 0xF8F9},
		",": piqad{character: ",", unicode: 0xF8FE},
		".": piqad{character: ".", unicode: 0xF8FD},
		" ": piqad{character: " ", unicode: 0x0020},
	}
)
