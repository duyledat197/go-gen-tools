package models

import (
	"github.com/cucumber/messages-go/v16"
)

// Feature is an internal object to group together
// the parsed gherkin document, the pickles and the
// raw content.
type Feature struct {
	*messages.GherkinDocument
	Pickles []*messages.Pickle
	Content []byte
}

// FindScenario ...
func (f Feature) FindScenario(astScenarioID string) *messages.Scenario {
	for _, child := range f.GherkinDocument.Feature.Children {
		if sc := child.Scenario; sc != nil && sc.Id == astScenarioID {
			return sc
		}
	}
	return nil
}
