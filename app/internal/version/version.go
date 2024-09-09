package version

var Version string

func GetVersion() string {
	if Version == "" {
		return "development"
	}
	return Version
}
