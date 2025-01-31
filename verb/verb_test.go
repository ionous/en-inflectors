package verb_test

import (
	"testing"

	"github.com/ionous/en-inflectors/verb"
)

func testConjugations(t *testing.T, entry string, expected []string) {
	for i, to := range []verb.VerbType{verb.VBP, verb.VBD, verb.VBN, verb.VBZ, verb.VBG} {
		result := to.Conjugate(entry)
		if x := expected[i]; x == result {
			t.Log("ok:", entry, to, result)
		} else {
			t.Log("ng:", entry, to, result, "expected:", x)
			t.Fail()
		}
	}
}

func TestVerbList(t *testing.T) {
	for k, v := range verbList {
		testConjugations(t, k, v)
	}
}

// from: test/verbs/verbs.ts
var verbList = map[string][]string{
	"sublet":           {"sublet", "sublet", "sublet", "sublets", "subletting"},
	"sunburn":          {"sunburn", "sunburnt", "sunburnt", "sunburns", "sunburning"},
	"telecast":         {"telecast", "telecast", "telecast", "telecasts", "telecasting"},
	"typecast":         {"typecast", "typecast", "typecast", "typecasts", "typecasting"},
	"typewrite":        {"typewrite", "typewrote", "typewritten", "typewrites", "typewriting"},
	"unbending":        {"unbend", "unbent", "unbent", "unbends", "unbending"},
	"unbind":           {"unbind", "unbound", "unbound", "unbinds", "unbinding"},
	"unclothe":         {"unclothe", "unclad", "unclad", "unclothes", "unclothing"},
	"undercuts":        {"undercut", "undercut", "undercut", "undercuts", "undercutting"},
	"underfeeding":     {"underfeed", "underfed", "underfed", "underfeeds", "underfeeding"},
	"underwent":        {"undergo", "underwent", "undergone", "undergoes", "undergoing"},
	"whelps":           {"whelp", "whelped", "whelped", "whelps", "whelping"},
	"threatening":      {"threaten", "threatened", "threatened", "threatens", "threatening"},
	"prospers":         {"prosper", "prospered", "prospered", "prospers", "prospering"},
	"toe":              {"toe", "toed", "toed", "toes", "toeing"},
	"toenailing":       {"toenail", "toenailed", "toenailed", "toenails", "toenailing"},
	"fluctuated":       {"fluctuate", "fluctuated", "fluctuated", "fluctuates", "fluctuating"},
	"vacillate":        {"vacillate", "vacillated", "vacillated", "vacillates", "vacillating"},
	"discards":         {"discard", "discarded", "discarded", "discards", "discarding"},
	"convulsed":        {"convulse", "convulsed", "convulsed", "convulses", "convulsing"},
	"hampers":          {"hamper", "hampered", "hampered", "hampers", "hampering"},
	"loathed":          {"loathe", "loathed", "loathed", "loathes", "loathing"},
	"abominated":       {"abominate", "abominated", "abominated", "abominates", "abominating"},
	"execrated":        {"execrate", "execrated", "execrated", "execrates", "execrating"},
	"disfavored":       {"disfavor", "disfavored", "disfavored", "disfavors", "disfavoring"},
	"despised":         {"despise", "despised", "despised", "despises", "despising"},
	"disinherited":     {"disinherit", "disinherited", "disinherited", "disinherits", "disinheriting"},
	"eschewed":         {"eschew", "eschewed", "eschewed", "eschews", "eschewing"},
	"apostatized":      {"apostatize", "apostatized", "apostatized", "apostatizes", "apostatizing"},
	"preponderated":    {"preponderate", "preponderated", "preponderated", "preponderates", "preponderating"},
	"outweighed":       {"outweigh", "outweighed", "outweighed", "outweighs", "outweighing"},
	"overbalance":      {"overbalance", "overbalanced", "overbalanced", "overbalances", "overbalancing"},
	"outbalance":       {"outbalance", "outbalanced", "outbalanced", "outbalances", "outbalancing"},
	"ousted":           {"oust", "ousted", "ousted", "ousts", "ousting"},
	"scoffing":         {"scoff", "scoffed", "scoffed", "scoffs", "scoffing"},
	"flouting":         {"flout", "flouted", "flouted", "flouts", "flouting"},
	"spoofing":         {"spoof", "spoofed", "spoofed", "spoofs", "spoofing"},
	"burlesquing":      {"burlesque", "burlesqued", "burlesqued", "burlesques", "burlesquing"},
	"suffusing":        {"suffuse", "suffused", "suffused", "suffuses", "suffusing"},
	"misbehaving":      {"misbehave", "misbehaved", "misbehaved", "misbehaves", "misbehaving"},
	"requiting":        {"requite", "requited", "requited", "requites", "requiting"},
	"crippling":        {"cripple", "crippled", "crippled", "cripples", "crippling"},
	"lame":             {"lame", "lamed", "lamed", "lames", "laming"},
	"stultifying":      {"stultify", "stultified", "stultified", "stultifies", "stultifying"},
	"tousle":           {"tousle", "tousled", "tousled", "tousles", "tousling"},
	"knots":            {"knot", "knotted", "knotted", "knots", "knotting"},
	"unscrambles":      {"unscramble", "unscrambled", "unscrambled", "unscrambles", "unscrambling"},
	"unpack":           {"unpack", "unpacked", "unpacked", "unpacks", "unpacking"},
	"moors":            {"moor", "moored", "moored", "moors", "mooring"},
	"wharfs":           {"wharf", "wharfed", "wharfed", "wharfs", "wharfing"},
	"levitates":        {"levitate", "levitated", "levitated", "levitates", "levitating"},
	"declaims":         {"declaim", "declaimed", "declaimed", "declaims", "declaiming"},
	"autographs":       {"autograph", "autographed", "autographed", "autographs", "autographing"},
	"overuses":         {"overuse", "overused", "overused", "overuses", "overusing"},
	"thrones":          {"throne", "throned", "throned", "thrones", "throning"},
	"overvalues":       {"overvalue", "overvalued", "overvalued", "overvalues", "overvaluing"},
	"overestimates":    {"overestimate", "overestimated", "overestimated", "overestimates", "overestimating"},
	"overrates":        {"overrate", "overrated", "overrated", "overrates", "overrating"},
	"retools":          {"retool", "retooled", "retooled", "retools", "retooling"},
	"revises":          {"revise", "revised", "revised", "revises", "revising"},
	"footnotes":        {"footnote", "footnoted", "footnoted", "footnotes", "footnoting"},
	"tickles":          {"tickle", "tickled", "tickled", "tickles", "tickling"},
	"seeps":            {"seep", "seeped", "seeped", "seeps", "seeping"},
	"palpitates":       {"palpitate", "palpitated", "palpitated", "palpitates", "palpitating"},
	"teeters":          {"teeter", "teetered", "teetered", "teeters", "teetering"},
	"totters":          {"totter", "tottered", "tottered", "totters", "tottering"},
	"toddles":          {"toddle", "toddled", "toddled", "toddles", "toddling"},
	"satirizes":        {"satirize", "satirized", "satirized", "satirizes", "satirizing"},
	"lampoons":         {"lampoon", "lampooned", "lampooned", "lampoons", "lampooning"},
	"frazzles":         {"frazzle", "frazzled", "frazzled", "frazzles", "frazzling"},
	"addles":           {"addle", "addled", "addled", "addles", "addling"},
	"muddles":          {"muddle", "muddled", "muddled", "muddles", "muddling"},
	"malingers":        {"malinger", "malingered", "malingered", "malingers", "malingering"},
	"skulks":           {"skulk", "skulked", "skulked", "skulks", "skulking"},
	"scrounges":        {"scrounge", "scrounged", "scrounged", "scrounges", "scrounging"},
	"forages":          {"forage", "foraged", "foraged", "forages", "foraging"},
	"strolls":          {"stroll", "strolled", "strolled", "strolls", "strolling"},
	"tingles":          {"tingle", "tingled", "tingled", "tingles", "tingling"},
	"disinclines":      {"disincline", "disinclined", "disinclined", "disinclines", "disinclining"},
	"reimburses":       {"reimburse", "reimbursed", "reimbursed", "reimburses", "reimbursing"},
	"transliterates":   {"transliterate", "transliterated", "transliterated", "transliterates", "transliterating"},
	"transcribes":      {"transcribe", "transcribed", "transcribed", "transcribes", "transcribing"},
	"unsublet":         {"unsublet", "unsublet", "unsublet", "unsublets", "unsubletting"},
	"insunburn":        {"insunburn", "insunburnt", "insunburnt", "insunburns", "insunburning"},
	"clone-telecast":   {"clone-telecast", "clone-telecast", "clone-telecast", "clone-telecasts", "clone-telecasting"},
	"offline-typecast": {"offline-typecast", "offline-typecast", "offline-typecast", "offline-typecasts", "offline-typecasting"},
	"force-typewrite":  {"force-typewrite", "force-typewrote", "force-typewritten", "force-typewrites", "force-typewriting"},
	"geo-unbending":    {"geo-unbend", "geo-unbent", "geo-unbent", "geo-unbends", "geo-unbending"},
	"deimitated":       {"deimitate", "deimitated", "deimitated", "deimitates", "deimitating"},
	"co-immolate":      {"co-immolate", "co-immolated", "co-immolated", "co-immolates", "co-immolating"},
}
