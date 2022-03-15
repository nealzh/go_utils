package regx_utils

import "regexp"

var (
	LineBreakRegx, lbrErr = regexp.Compile("\r\n|\r|\n")
	BlankRegx, brErr      = regexp.Compile("\\s+")
)
