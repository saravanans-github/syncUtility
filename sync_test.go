package main

import (
	"testing"
)

func TestGetParamsFromCommandLine(t *testing.T) {
	clParams := clParametersType{"file://DUMVPVODOS03/mshe/SS/FTDUMVP/624/458/T00109103101.m3u8", "file://DUMVPVODOS04/mshe/SS/FTDUMVP/624/458/T00109103101.m3u8"}

	getParamsFromCommandLine(&clParams)
}
