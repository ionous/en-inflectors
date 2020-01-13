package verb

import (
	"regexp"
	"strings"

	"github.com/reiver/go-porterstemmer"
)

func tryPresent(verb string) (ret string) {
	// Gerunds
	if strings.HasSuffix(verb, "ing") {
		if a, ok := VBP[replace(verb, `ing$`, "")]; ok {
			ret = a[0]
		} else if b, ok := VBP[replace(verb, `ing$`, "e")]; ok {
			ret = b[0]
		} else if c, ok := VBP[replace(verb, `.ing$`, "")]; ok {
			ret = c[0]
		} else if d, ok := VBP[replace(verb, `ying$`, "ie")]; ok {
			ret = d[0]
		}
	} else if strings.HasSuffix(verb, "s") {
		// VBZs
		if a, ok := VBP[replace(verb, `s$`, "")]; ok {
			ret = a[0]
		} else if b, ok := VBP[replace(verb, `es$`, "")]; ok {
			ret = b[0]
		} else if c, ok := VBP[replace(verb, `ies$`, "y")]; ok {
			ret = c[0]
		}
	} else if strings.HasSuffix(verb, "ed") {
		// VBNs & VBDs
		if a, ok := VBP[replace(verb, `ed$`, "")]; ok {
			ret = a[0]
		} else if b, ok := VBP[replace(verb, `d$`, "")]; ok {
			ret = b[0]
		} else if c, ok := VBP[replace(verb, `ied$`, "y")]; ok {
			ret = c[0]
		} else if d, ok := VBP[replace(verb, `ed$`, "y")]; ok {
			ret = d[0]
		} else if e, ok := VBP[replace(verb, `ed$`, "e")]; ok {
			ret = e[0]
		} else if f, ok := VBP[replace(verb, `.ed$`, "")]; ok {
			ret = f[0]
		}
	} else if a, ok := VBP[porterstemmer.StemString(verb)]; ok {
		// finally defaults to porter stemmer
		ret = a[0]
	}
	return
}

func replace(str, reg, rep string) string {
	return regexp.MustCompile(reg).ReplaceAllLiteralString(str, rep)
}
