package ver

import (
	"fmt"
)

// VersionName -
var VersionName = "n/a"

// GitCommit -
var GitCommit = "n/a"

// BuildDate -
var BuildDate = "n/a"

// ShowVersion -
func ShowVersion() {
	fmt.Printf("Version: %s\nCommit: %s\nDate: %s\n", VersionName, GitCommit, BuildDate)
}
