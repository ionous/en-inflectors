package verb

import "regexp"

type rule struct {
	test      *regexp.Regexp
	transform func(string, string) string
}

func solveRegex(input, to string) (ret string) {
	for _, rule := range rules {
		if rule.test.MatchString(input) {
			ret = rule.transform(input, to)
			break
		}
	}
	return
}

var rules = []rule{
	// [long vowel][consonant]
	{
		test: regexp.MustCompile(`([uao]m[pb]|[oa]wn|ey|elp|[ei]gn|ilm|o[uo]r|[oa]ugh|igh|ki|ff|oubt|ount|awl|o[alo]d|[iu]rl|` +
			`upt|[oa]y|ight|oid|empt|act|aud|e[ea]d|ound|[aeiou][srcln]t|ept|dd|[eia]n[dk]|[ioa][xk]|[oa]rm|[ue]` +
			`rn|[ao]ng|uin|eam|ai[mr]|[oea]w|[eaoui][rscl]k|[oa]r[nd]|ear|er|[^aieouyfm]it|[aeiouy]ir|[^aieouyfm]` +
			`et|ll|en|vil|om|[^rno].mit|rup|bat|val|.[^skxwb][rvmchslwngb]el)$`),
		transform: func(vb, to string) (ret string) {
			if to == "VBZ" {
				ret = vb + "s"
			} else if to == "VBG" {
				ret = vb + "ing"
			} else if to == "VBN" || to == "VBD" {
				ret = vb + "ed"
			} else {
				ret = vb
			}
			return
		},
	},

	// [consonant][y]
	{
		test: regexp.MustCompile(`[^aeiou]y$`),
		transform: func(vb, to string) (ret string) {
			base := vb[0 : len(vb)-1]
			if to == "VBZ" {
				ret = base + "ies"
			} else if to == "VBG" {
				ret = vb + "ing"
			} else if to == "VBN" || to == "VBD" {
				ret = base + "ied"
			} else {
				ret = vb
			}
			return
		},
	},

	// [consonant][e]
	{
		test: regexp.MustCompile(`[^aeiouy]e$`),
		transform: func(vb, to string) (ret string) {
			base := vb[0 : len(vb)-1]
			if to == "VBZ" {
				ret = vb + "s"
			} else if to == "VBG" {
				ret = base + "ing"
			} else if to == "VBN" || to == "VBD" {
				ret = base + "ed"
			} else {
				ret = vb
			}
			return
		},
	},

	// [short vowel][consonant]
	{
		test: regexp.MustCompile(`([^dtaeiuo][aeiuo][ptlgnm]|ir|cur|ug|[hj]ar|[^aouiey]ep|[^aeiuo][oua][db])$`),
		transform: func(vb, to string) (ret string) {
			if to == "VBZ" {
				ret = vb + "s"
			} else if to == "VBG" {
				ret = vb + vb[len(vb)-1:] + "ing"
			} else if to == "VBN" || to == "VBD" {
				ret = vb + vb[len(vb)-1:] + "ed"
			} else {
				ret = vb
			}
			return
		},
	},

	// [sibilant]
	{
		test: regexp.MustCompile(`([ieao]ss|[aeiouy]zz|[aeiouy]ch|nch|rch|[aeiouy]sh|[iae]tch|ax)$`),
		transform: func(vb, to string) (ret string) {
			if to == "VBZ" {
				ret = vb + "es"
			} else if to == "VBG" {
				ret = vb + "ing"
			} else if to == "VBN" || to == "VBD" {
				ret = vb + "ed"
			} else {
				ret = vb
			}
			return
		},
	},

	// [e][e]
	{
		test: regexp.MustCompile(`(ee)$`),
		transform: func(vb, to string) (ret string) {
			if to == "VBZ" {
				ret = vb + "s"
			} else if to == "VBG" {
				ret = vb + "ing"
			} else if to == "VBN" || to == "VBD" {
				ret = vb + "d"
			} else {
				ret = vb
			}
			return
		},
	},

	// [i][e]
	{
		test: regexp.MustCompile(`(ie)$`),
		transform: func(vb, to string) (ret string) {
			if to == "VBZ" {
				ret = vb + "s"
			} else if to == "VBG" {
				ret = vb[:len(vb)-2] + "ying"
			} else if to == "VBN" || to == "VBD" {
				ret = vb + "d"
			} else {
				ret = vb
			}
			return
		},
	},

	// [u][e]
	{
		test: regexp.MustCompile(`(ue)$`),
		transform: func(vb, to string) (ret string) {
			if to == "VBZ" {
				ret = vb + "s"
			} else if to == "VBG" {
				ret = vb[:len(vb)-1] + "ing"
			} else if to == "VBN" || to == "VBD" {
				ret = vb + "d"
			} else {
				ret = vb
			}
			return
		},
	},

	// default (regular)
	{
		test: regexp.MustCompile(`.`),
		transform: func(vb, to string) (ret string) {
			if to == "VBZ" {
				ret = vb + "s"
			} else if to == "VBG" {
				ret = vb + "ing"
			} else if to == "VBN" || to == "VBD" {
				ret = vb + "ed"
			} else {
				ret = vb
			}
			return
		},
	},
}
