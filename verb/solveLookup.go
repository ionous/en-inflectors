package verb

func lookup(input string) (ret []string) {
	if x, ok := VBP[input]; ok {
		ret = x
	} else if x, ok := VBD[input]; ok {
		ret = x
	} else if x, ok := VBN[input]; ok {
		ret = x
	} else if x, ok := VBZ[input]; ok {
		ret = x
	} else if x, ok := VBG[input]; ok {
		ret = x
	}
	return
}

func dirIndex(dir string) int {
	at := -1
	for i, el := range []string{"VBP", "VBD", "VBN", "VBZ", "VBG"} {
		if el == dir {
			at = i
			break
		}
	}
	if at < 0 {
		panic("not found " + dir)
	}
	return at
}

func conjugate(input, direction string) string {
	i := dirIndex(direction)
	return lookup(input)[i]
}

func init() {
	for _, present := range regular {
		past := solveRegex(present, "VBD")
		pastParticiple := solveRegex(present, "VBN")
		present3rd := solveRegex(present, "VBZ")
		gerund := solveRegex(present, "VBG")
		entry := []string{
			present,
			past,
			pastParticiple,
			present3rd,
			gerund,
		}
		VBP[present] = entry
		VBD[past] = entry
		VBN[pastParticiple] = entry
		VBZ[present3rd] = entry
		VBG[gerund] = entry
	}
}
