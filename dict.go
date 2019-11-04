package roops

var dict map[rune]rune

func init() {
	// Values, used to fill up dictionary
	ru := " 1234567890!ЙЦУКЕНГШЩЗХЪФЫВАПРОЛДЖЭЯЧСМИТЬБЮйцукенгшщзхъфывапролджэячсмитьбю"
	en := ` 1234567890!QWERTYUIOP{}ASDFGHJKL:"ZXCVBNM<>qwertyuiop[]asdfghjkl;'zxcvbnm,.`

	enRunes := []rune(en)
	// Filling dictionary
	dict = map[rune]rune{}
	for i, j := range []rune(ru) {
		dict[j] = enRunes[i]
	}
}
