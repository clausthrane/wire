package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenFactoriesAreRegisteredWithASingletonProviderTheyAreOnlyInvokedOnce(t *testing.T) {
	p := NewSingletonRepository()

	value := &struct {
		count int
	}{0}

	p.Register("test1", func() interface{} {
		value.count++
		return value
	})

	v1 := p.Get("test1")
	v2 := p.Get("test1")

	assert.Exactly(t, v1, v2)
	assert.Equal(t, 1, value.count)
}
