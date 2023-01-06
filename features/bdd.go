package features

import "github.com/cucumber/godog"

type Suite struct {
}

// godogsCtxKey is the key used to store the available godogs in the context.Context.
type godogsCtxKey struct{}

func (s *Suite) GetSteps() map[string]interface{} {
	return map[string]interface{}{
		// example:
		// `^there are (\d+) godogs$`: s.thereAreGodogs,

		/*generate_key*/
	}
}
func (s *Suite) RegisterStep(sc *godog.ScenarioContext) {
	steps := s.GetSteps()
	for step, fn := range steps {
		sc.Step(step, fn)
	}
}

func aBackground() error {
	return godog.ErrPending
}

func aSignedIn(arg1 string) error {
	return godog.ErrPending
}
