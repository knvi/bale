package regex

import "regexp"

const (
	unixPath = `^(.*)\/([^\/]*)$`
	winPath  = `^(.*)\\([^\\]*)$`

	alphaNum = `^[a-zA-Z0-9]+$`
)

var (
	UnixPath = regexp.MustCompile(unixPath)
	WinPath  = regexp.MustCompile(winPath)

	AlphaNum = regexp.MustCompile(alphaNum)
)