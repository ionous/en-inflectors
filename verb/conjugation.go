package verb

type VerbType int

//go:generate stringer -type=VerbType
const (
	VBP VerbType = iota
	VBD
	VBN
	VBZ
	VBG
)

type conjugationObject map[string][]string

var vbpTable = make(conjugationObject)
var vbdTable = make(conjugationObject)
var vbnTable = make(conjugationObject)
var vbzTable = make(conjugationObject)
var vbgTable = make(conjugationObject)
