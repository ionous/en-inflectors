package verb

// warning: many early returns
func (to VerbType) Conjugate(input string) string {
	// direct dictionary lookup
	if str, ok := to.lookup(input); ok {
		return str
	} else if r := vbpTable[input]; len(r) > 0 {
		// passed VBP => regex solution
		return to.solveRegex(input)
	} else if stripped := strip(input); len(stripped) > 0 {
		// try to strip the verb
		// "foreseeing" => "seeing"
		// the stripping result is in the dictionary (doesn't have to be VBP)
		if lookupStripped, ok := to.lookup(stripped); ok {
			return rebuild(input, stripped, lookupStripped)
		} else if el := vbpTable[stripped]; len(el) > 0 {
			// the stripped verb is in VBP form (we can apply regexes)
			return rebuild(input, stripped, to.solveRegex(input))
			// Note: we don't need to try to convert to VBP
			// since this step takes characters from the beginning
			// so it's best to convert the original (not the stripped)
			// to present
		}
	}

	// try to convert it to VBP form
	if stem := tryPresent(input); len(stem) > 0 {
		if el, ok := to.lookup(stem); ok {
			return el
		} else {
			input = stem
		}
	}

	// final shot (will give wrong results if passed a verb that is not a VBP)
	return to.solveRegex(input)
}

// ToPresent provides the base form of the word
// "typewrite"
func ToPresent(input string) string {
	return VBP.Conjugate(input)
}

// ToPast transforms the input into the simple past.
// ex. "typewrote"
func ToPast(input string) string {
	return VBD.Conjugate(input)
}

// ToPastParticiple transforms the input into the past participle;
// for regular verbs this is the same as the simple past.
// "typewritten"
func ToPastParticiple(input string) string {
	return VBN.Conjugate(input)
}

// ToPresentSignular transforms the input into the third-person singular present.
// "typewrites"; Third person singular present
func ToPresentSingular(input string) string {
	return VBZ.Conjugate(input)
}

// ToGerund is also used for the present participle,
// it nouns a verb, ending it with ing.
// "typewriting"
func ToGerund(input string) string {
	return VBG.Conjugate(input)
}
