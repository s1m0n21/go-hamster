package build

import (
	"fmt"
)

const (
	version  = "0.0.1"
	revision = "200312"
	status 	 = "beta"
)

func CurrentVersion() string {
	return fmt.Sprintf("%v.%v_%v", version, revision, status)
}
