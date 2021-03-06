package jsonpath

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStackPush(t *testing.T) {
	as := assert.New(t)
	s := newIntStack()

	s.push(5)
	as.Equal(s.len(), 1)

	s.push(12)
	as.Equal(s.len(), 2)
}

func TestStackPop(t *testing.T) {
	as := assert.New(t)
	s := newIntStack()

	s.push(99)
	as.Equal(s.len(), 1)

	v, ok := s.pop()
	as.True(ok)
	as.Equal(99, v)

	as.Equal(s.len(), 0)
}

func TestStackPeek(t *testing.T) {
	as := assert.New(t)
	s := newIntStack()

	s.push(99)
	v, ok := s.peek()
	as.True(ok)
	as.Equal(99, v)

	s.push(54)
	v, ok = s.peek()
	as.True(ok)
	as.Equal(54, v)

	s.pop()
	v, ok = s.peek()
	as.True(ok)
	as.Equal(99, v)

	s.pop()
	_, ok = s.peek()
	as.False(ok)
}
