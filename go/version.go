package aitokenizer

import "runtime"

var (
	ServiceName string
	Version     string
	Revision    string
	Branch      string
	BuildDate   string
	OSArch      string
	GoVersion   = runtime.Version()
)
