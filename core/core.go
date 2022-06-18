package core

import (
	"runtime"
	"github.com/xtls/xray-core/common/serial"
)

func Version() string {
	return version
}

func VersionStatement() []string {
	return []string{
		serial.Concat("Xray ", Version(), " (", codename, ") ", build, " (", runtime.Version(), " ", runtime.GOOS, "/", runtime.GOARCH, ")"),
		intro,
	}
}
