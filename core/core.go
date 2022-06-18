package core

import (
	"runtime"
	"github.com/xtls/xray-core/common/serial"
)

var (
        version  = "1.0"
        build    = "Custom"
        codename = "Xray-Lite, Penetrates Everything."
        intro    = "A unified platform for anti-censorship."
)

func Version() string {
	return version
}

func VersionStatement() []string {
	return []string{
		serial.Concat("Xray-Lite ", Version(), " (", codename, ") ", build, " (", runtime.Version(), " ", runtime.GOOS, "/", runtime.GOARCH, ")"),
		intro,
	}
}
