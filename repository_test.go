package wire

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhenDependenciesAreRegisteredTheyCanBeResolved(t *testing.T) {

	key := "some-key"
	value := "some-value"

	p := SingleDependencyRepository(key, func() interface{} {
		return value
	})

	result := p.Get(key)
	assert.NotNil(t, result)
	assert.Exactly(t, value, result)
}

func TestWhenCreatingASingleDependencyProviderOnlyThatObjectIsResolvable(t *testing.T) {

	key := "some-key"
	value := "some-value"

	p := SingleDependencyRepository(key, func() interface{} {
		return value
	})
	result := p.Get(key)
	assert.NotNil(t, result)
	assert.Exactly(t, value, result)

	otherResult := p.Get("otherkey")
	assert.Nil(t, otherResult)
}
