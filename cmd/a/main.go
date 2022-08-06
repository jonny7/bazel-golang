package main

import (
	"fmt"
	"github.com/jonny7/bazel-golang/logging"
	"github.com/jonny7/bazel-golang/version"
)

func main() {
	logging.LogMessage(fmt.Sprintf("I am binary A, my version is: %s", version.Version))
}
