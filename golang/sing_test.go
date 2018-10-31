package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_SingFly(t *testing.T) {
	song := `There was an old lady who swallowed a fly.
I don't know why she swallowed a fly - perhaps she'll die!`

	stringBuilder := &strings.Builder{}
	writer := func(sing string) {
		stringBuilder.WriteString(sing)
	}
	fly := NewSinger("fly")
	fly.writer = writer
	fly.Sing()

	assert.Equal(t, song, stringBuilder.String())
}
