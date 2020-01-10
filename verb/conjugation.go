package verb

type conjugationObject map[string][]string

var VBP = make(conjugationObject)
var VBD = make(conjugationObject)
var VBN = make(conjugationObject)
var VBZ = make(conjugationObject)
var VBG = make(conjugationObject)

func init() {
	for _, entry := range irregular {
		VBP[entry[0]] = entry
		//
		if len(entry[1]) > 0 {
			VBD[entry[1]] = entry
		}
		if len(entry[2]) > 0 {
			VBN[entry[2]] = entry
		}
		if len(entry[3]) > 0 {
			VBZ[entry[3]] = entry
		}
		if len(entry[4]) > 0 {
			VBG[entry[4]] = entry
		}
	}

}
