package main

import (
	"fmt"
	"github.com/jonny7/bazel-golang/logging"
	"github.com/jonny7/bazel-golang/names"
	"github.com/jonny7/bazel-golang/version"
)

func main() {
	random := names.Name()
	logging.LogMessage(fmt.Sprintf("I am binary B, my version is: %s, and I was given a random name of :%s", version.Version, random))
}
