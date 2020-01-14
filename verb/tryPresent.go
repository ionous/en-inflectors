package verb

import (
	"regexp"
	"strings"

	"github.com/reiver/go-porterstemmer"
)

func tryPresent(verb string) (ret string) {
	// Gerunds
	if strings.HasSuffix(verb, "ing") {
		if a, ok := vbpTable[replace(verb, `ing$`, "")]; ok {
			ret = a[0]
		} else if b, ok := vbpTable[replace(verb, `ing$`, "e")]; ok {
			ret = b[0]
		} else if c, ok := vbpTable[replace(verb, `.ing$`, "")]; ok {
			ret = c[0]
		} else if d, ok := vbpTable[replace(verb, `ying$`, "ie")]; ok {
			ret = d[0]
		}
	} else if strings.HasSuffix(verb, "s") {
		// VBZs
		if a, ok := vbpTable[replace(verb, `s$`, "")]; ok {
			ret = a[0]
		} else if b, ok := vbpTable[replace(verb, `es$`, "")]; ok {
			ret = b[0]
		} else if c, ok := vbpTable[replace(verb, `ies$`, "y")]; ok {
			ret = c[0]
		}
	} else if strings.HasSuffix(verb, "ed") {
		// VBNs & VBDs
		if a, ok := vbpTable[replace(verb, `ed$`, "")]; ok {
			ret = a[0]
		} else if b, ok := vbpTable[replace(verb, `d$`, "")]; ok {
			ret = b[0]
		} else if c, ok := vbpTable[replace(verb, `ied$`, "y")]; ok {
			ret = c[0]
		} else if d, ok := vbpTable[replace(verb, `ed$`, "y")]; ok {
			ret = d[0]
		} else if e, ok := vbpTable[replace(verb, `ed$`, "e")]; ok {
			ret = e[0]
		} else if f, ok := vbpTable[replace(verb, `.ed$`, "")]; ok {
			ret = f[0]
		}
	} else if a, ok := vbpTable[porterstemmer.StemString(verb)]; ok {
		// finally defaults to porter stemmer
		ret = a[0]
	}
	return
}

func replace(str, reg, rep string) string {
	return regexp.MustCompile(reg).ReplaceAllLiteralString(str, rep)
}
