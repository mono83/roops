package roops

// ToLatin attempts to convert provided source string into
// latin characters using conversion dictionary, based on
// russian keyboard layout.
//
// For example "Ащщ" will be converted into "Foo"
//
// Conversion will be applied only if all characters inside
// string can be converted, so this function does not work
// in partial mode.
//
// If coversion is successful, this function will return
// result and true. In other cases this func will return
// source string and false.
//
// PS. Works on standard Windows keyboard layouts.
func ToLatin(src string) (res string, replaced bool) {
	if len(src) == 0 {
		return src, false
	}

	var result []rune
	for _, j := range []rune(src) {
		replacement, ok := dict[j]
		if !ok {
			// At least one character is not replaceable
			return src, false
		}
		result = append(result, replacement)
	}

	return string(result), true
}
