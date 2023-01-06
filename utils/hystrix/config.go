package hystrix

import (
	"time"

	"github.com/afex/hystrix-go/hystrix"
)

// HystrixConfig ...
func HystrixConfig() hystrix.CommandConfig {
	return hystrix.CommandConfig{
		// How long to wait for command to complete, in milliseconds
		Timeout: int(15 * time.Second.Milliseconds()),

		// MaxConcurrent is how many commands of the same type
		// can run at the same time
		MaxConcurrentRequests: 20000,

		// VolumeThreshold is the minimum number of requests
		// needed before a circuit can be tripped due to health
		RequestVolumeThreshold: 100,

		// SleepWindow is how long, in milliseconds,
		// to wait after a circuit opens before testing for recovery
		SleepWindow: int(5 * time.Second.Milliseconds()),

		// ErrorPercentThreshold causes circuits to open once
		// the rolling measure of errors exceeds this percent of requests
		ErrorPercentThreshold: 80,
	}
}
