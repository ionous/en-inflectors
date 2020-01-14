package verb

// in go, this guards indexing an empty array
func (to VerbType) lookup(input string) (ret string, okay bool) {
	i := int(to)
	if slice := lookup(input); i >= 0 && i < len(slice) {
		if str := slice[i]; len(str) > 0 {
			ret, okay = str, true
		}
	}
	return
}

// work backwards from an input to find a conjugation table row
func lookup(input string) (ret []string) {
	if x, ok := vbpTable[input]; ok {
		ret = x
	} else if x, ok := vbdTable[input]; ok {
		ret = x
	} else if x, ok := vbnTable[input]; ok {
		ret = x
	} else if x, ok := vbzTable[input]; ok {
		ret = x
	} else if x, ok := vbgTable[input]; ok {
		ret = x
	}
	return
}
