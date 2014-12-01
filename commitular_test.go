package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFormatSubject(t *testing.T) {
	var subject = FormatSubject("type", "scope", "description")
	assert.Equal(t, "type(scope): description", subject)
}
