package main

import "strings"

type BuildFlag string

const (
	FrontendBuildFlag BuildFlag = "frontend"
	AllBuildFlag      BuildFlag = "all"
)

func getUsageBuildFlagString() string {
	return strings.Join([]string{string(FrontendBuildFlag), string(AllBuildFlag)}, " | ")
}
