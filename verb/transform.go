package verb

// warning: many early returns
func Conjugate(input, to string) string {
	toIndex := dirIndex(to)
	// direct dictionary lookup
	if str, ok := lookupAt(input, toIndex); ok {
		return str
	} else if r := VBP[input]; len(r) > 0 {
		// passed VBP => regex solution
		return solveRegex(input, to)
	} else if stripped := strip(input); len(stripped) > 0 {
		// try to strip the verb
		// "foreseeing" => "seeing"
		// the stripping result is in the dictionary (doesn't have to be VBP)
		if lookupStripped, ok := lookupAt(stripped, toIndex); ok {
			return rebuild(input, stripped, lookupStripped)
		} else if el := VBP[stripped]; len(el) > 0 {
			// the stripped verb is in VBP form (we can apply regexes)
			return rebuild(input, stripped, solveRegex(input, to))
			// Note: we don't need to try to convert to VBP
			// since this step takes characters from the beginning
			// so it's best to convert the original (not the stripped)
			// to present
		}
	}

	// try to convert it to VBP form
	if stem := tryPresent(input); len(stem) > 0 {
		if el, ok := lookupAt(stem, toIndex); ok {
			return el
		} else {
			input = stem
		}
	}

	// final shot (will give wrong results if passed a verb that is not a VBP)
	return solveRegex(input, to)
}

func ToPresent(input string) string {
	return Conjugate(input, "VBP")
}
func ToPast(input string) string {
	return Conjugate(input, "VBD")
}
func ToPastParticiple(input string) string {
	return Conjugate(input, "VBN")
}
func ToPresentS(input string) string {
	return Conjugate(input, "VBZ")
}
func ToGerund(input string) string {
	return Conjugate(input, "VBG")
}
