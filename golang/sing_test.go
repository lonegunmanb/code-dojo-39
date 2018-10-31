package main

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func Test_SingleSinger(t *testing.T) {
	song := `There was an old lady who swallowed a fly.
I don't know why she swallowed a fly - perhaps she'll die!`
	formUp()
	sb, writer := buildWriter()
	fly := chorus[0]
	fly.Sing(writer)

	assert.Equal(t, song, sb.String())
}

func Test_DoubleSinger(t *testing.T){
	song := `There was an old lady who swallowed a fly.
I don't know why she swallowed a fly - perhaps she'll die!
There was an old lady who swallowed a spider;
That wriggled and wiggled and tickled inside her.
She swallowed the spider to catch the fly;
I don't know why she swallowed a fly - perhaps she'll die!`
	formUp()
	sb, writer := buildWriter()

	spider := chorus[1]
	spider.Sing(writer)

	assert.Equal(t, song, sb.String())
}

//func Test_TripleSinger(t *testing.T){
//	song := `There was an old lady who swallowed a fly.
//I don't know why she swallowed a fly - perhaps she'll die!
//There was an old lady who swallowed a spider;
//That wriggled and wiggled and tickled inside her.
//She swallowed the spider to catch the fly;
//I don't know why she swallowed a fly - perhaps she'll die!
//There was an old lady who swallowed a bird;
//How absurd to swallow a bird.
//She swallowed the bird to catch the spider,
//She swallowed the spider to catch the fly;
//I don't know why she swallowed a fly - perhaps she'll die!`
//
//}

func buildWriter() (*strings.Builder, func(sing string)) {
	stringBuilder := &strings.Builder{}
	writer := func(sing string) {
		stringBuilder.WriteString(sing)
	}
	return stringBuilder, writer
}
