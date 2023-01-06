package main

import (
	"testing"

	"github.com/duyledat197/go-gen-tools/features"
	"github.com/duyledat197/go-gen-tools/utils/pathutils"

	"github.com/cucumber/godog"
)

func TestFeatures(t *testing.T) {
	pkgDir := pathutils.GetPkgDir()

	suite := godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Format:   "pretty",
			Paths:    []string{pkgDir},
			TestingT: t, // Testing instance that will run subtests.
		},
	}

	if suite.Run() != 0 {
		t.Fatal("non-zero status returned, failed to run feature tests")
	}
}

func InitializeScenario(sc *godog.ScenarioContext) {
	s := &features.Suite{}
	s.RegisterStep(sc)
}
