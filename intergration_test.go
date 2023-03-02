package main

import (
	"os"
	"testing"

	"github.com/duyledat197/go-gen-tools/features"
	"github.com/duyledat197/go-gen-tools/utils/pathutils"

	"github.com/cucumber/godog"
	"github.com/cucumber/godog/colors"
)

func TestFeatures(t *testing.T) {
	pkgDir := pathutils.GetPkgDir()

	suite := &godog.TestSuite{
		ScenarioInitializer: InitializeScenario,
		Options: &godog.Options{
			Output:        colors.Colored(os.Stdout),
			Format:        "pretty",
			Strict:        true,
			Paths:         []string{pkgDir},
			TestingT:      t, // Testing instance that will run subtests.
			StopOnFailure: true,
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
