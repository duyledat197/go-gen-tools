package registry

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Consul(t *testing.T) {
	client, err := NewClient("localhost:8500")
	assert.NoError(t, err)
	assert.NotNil(t, client)
}
