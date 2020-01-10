package verb

// warning: many early returns
func xform(input, to string) string {
	// direct dictionary lookup
	look := lookup(input)
	toIndex := dirIndex(to)

	if str := look[toIndex]; len(str) > 0 {
		return look[toIndex]
	} else if r := VBP[input]; len(r) > 0 {
		// passed VBP => regex solution
		return solveRegex(input, to)
	} else if stripped := strip(input); len(stripped) > 0 {
		// try to strip the verb
		// "foreseeing" => "seeing"
		// the stripping result is in the dictionary (doesn't have to be VBP)
		if lookupStripped := lookup(stripped)[toIndex]; len(lookupStripped) > 0 {
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
		if el := lookup(stem)[toIndex]; len(el) > 0 {
			return el
		} else {
			return solveRegex(stem, to)
		}
	}

	// final shot (will give wrong results if passed a verb that is not a VBP)
	return solveRegex(input, to)
}

func ToPresent(input string) string {
	return xform(input, "VBP")
}
func ToPast(input string) string {
	return xform(input, "VBD")
}
func ToPastParticiple(input string) string {
	return xform(input, "VBN")
}
func ToPresentS(input string) string {
	return xform(input, "VBZ")
}
func ToGerund(input string) string {
	return xform(input, "VBG")
}
