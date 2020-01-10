package verb

import "strings"

/**
 *
 * Those two functions are used to strip down a verb to it's nearest
 * possible parent. for example the verb "foresee" will be stripped
 * to "see", so the transformation function can use the list data
 * to conjugate it, then it passes it to the rebuild function
 * to rebuild it.
 *
**/

func strip(verb string) (ret string) {
	if len(verb) >= 2 {
		verb = verb[1:]
		if l := lookup(verb); len(l) > 0 {
			ret = verb
		} else {
			ret = strip(verb)
		}
	}
	return
}

func rebuild(original, stripped, conjugated string) string {
	// fix?
	// original code has superfluous statements:
	// let rebuilt = original.substr(0, original.indexOf(stripped) + stripped.length);
	// rebuilt = original.split(stripped).join(conjugated);
	// return rebuilt;

	split := strings.Split(original, stripped)
	return strings.Join(split, conjugated)
}
