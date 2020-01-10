package verb

import (
	"regexp"
	"strings"

	"github.com/reiver/go-porterstemmer"
)

func tryPresent(verb string) (ret string) {
	// Gerunds
	if strings.HasSuffix(verb, "ing") {
		if a := VBP[replace(verb, `ing$`, "")]; len(a) > 0 {
			ret = a[0]
		} else if b := VBP[replace(verb, `ing$`, "e")]; len(b) > 0 {
			ret = b[0]
		} else if c := VBP[replace(verb, `.ing$`, "")]; len(c) > 0 {
			ret = c[0]
		} else if d := VBP[replace(verb, `ying$`, "ie")]; len(d) > 0 {
			ret = d[0]
		}
	} else if strings.HasSuffix(verb, "s") {
		// VBZs
		if a := VBP[replace(verb, `s$`, "")]; len(a) > 0 {
			ret = a[0]
		} else if b := VBP[replace(verb, `es$`, "")]; len(b) > 0 {
			ret = b[0]
		} else if c := VBP[replace(verb, `ies$`, "y")]; len(c) > 0 {
			ret = c[0]
		}
	} else if strings.HasSuffix(verb, "ed") {
		// VBNs & VBDs
		if a := VBP[replace(verb, `ed$`, "")]; len(a) > 0 {
			ret = a[0]
		} else if b := VBP[replace(verb, `d$`, "")]; len(b) > 0 {
			ret = b[0]
		} else if c := VBP[replace(verb, `ied$`, "y")]; len(c) > 0 {
			ret = c[0]
		} else if d := VBP[replace(verb, `ed$`, "y")]; len(d) > 0 {
			ret = d[0]
		} else if e := VBP[replace(verb, `ed$`, "e")]; len(e) > 0 {
			ret = e[0]
		} else if f := VBP[replace(verb, `.ed$`, "")]; len(f) > 0 {
			ret = f[0]
		}
	} else if a := VBP[porterstemmer.StemString(verb)]; len(a) > 0 {
		// finally defaults to porter stemmer
		ret = a[0]
	}
	return
}

func replace(str, reg, rep string) string {
	return regexp.MustCompile(reg).ReplaceAllLiteralString(str, rep)
}
