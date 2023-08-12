package validate

import "github.com/knvi/bale/internal/validate/regex"

func UnixPath(path string) bool {
	return regex.UnixPath.MatchString(path)
}

func WindowsPath(path string) bool {
	return regex.WinPath.MatchString(path)
}

func Alphanumeric(name string) bool {
	return regex.AlphaNum.MatchString(name)
}