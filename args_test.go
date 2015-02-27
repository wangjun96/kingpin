package kingpin

import (
	"github.com/stretchr/testify/assert"

	"testing"
)

func TestArgRemainder(t *testing.T) {
	a := newArgGroup()
	v := a.Arg("test", "").Strings()
	a.init()
	args := []string{"hello", "world"}
	tokens := tokenize(args)
	a.parse(tokens)
	assert.Equal(t, args, *v)
}

func TestArgRemainderErrorsWhenNotLast(t *testing.T) {
	a := newArgGroup()
	a.Arg("test", "").Strings()
	a.Arg("test2", "").String()
	assert.Error(t, a.init())
}

func TestArgMultipleRequired(t *testing.T) {
	a := newArgGroup()
	a.Arg("a", "").Required().String()
	a.Arg("b", "").Required().String()
	a.init()

	err := a.parse(tokenize([]string{}))
	assert.Error(t, err)
	err = a.parse(tokenize([]string{"A"}))
	assert.Error(t, err)
	err = a.parse(tokenize([]string{"A", "B"}))
	assert.NoError(t, err)
}

func TestInvalidArgsDefaultCanBeOverridden(t *testing.T) {
	a := newArgGroup()
	a.Arg("a", "").Default("invalid").Bool()
	assert.NoError(t, a.init())
	tokens := tokenize([]string{})
	err := a.parse(tokens)
	assert.Error(t, err)
}
