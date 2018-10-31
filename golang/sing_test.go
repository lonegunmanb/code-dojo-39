package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_SingleSinger(t *testing.T) {
	song := `There was an old lady who swallowed a fly.
I don't know why she swallowed a fly - perhaps she'll die!`
	testSing(t, song, 0)
}

func Test_DoubleSinger(t *testing.T){
	song := `There was an old lady who swallowed a fly.
I don't know why she swallowed a fly - perhaps she'll die!
There was an old lady who swallowed a spider;
That wriggled and wiggled and tickled inside her.
She swallowed the spider to catch the fly;
I don't know why she swallowed a fly - perhaps she'll die!`
	testSing(t, song, 1)
}

func Test_TripleSinger(t *testing.T){
	song := `There was an old lady who swallowed a fly.
I don't know why she swallowed a fly - perhaps she'll die!
There was an old lady who swallowed a spider;
That wriggled and wiggled and tickled inside her.
She swallowed the spider to catch the fly;
I don't know why she swallowed a fly - perhaps she'll die!
There was an old lady who swallowed a bird;
How absurd to swallow a bird.
She swallowed the bird to catch the spider,
She swallowed the spider to catch the fly;
I don't know why she swallowed a fly - perhaps she'll die!`
	testSing(t, song, 2)
}

func Test_ForthSinger(t *testing.T){
	song := `There was an old lady who swallowed a fly.
I don't know why she swallowed a fly - perhaps she'll die!
There was an old lady who swallowed a spider;
That wriggled and wiggled and tickled inside her.
She swallowed the spider to catch the fly;
I don't know why she swallowed a fly - perhaps she'll die!
There was an old lady who swallowed a bird;
How absurd to swallow a bird.
She swallowed the bird to catch the spider,
She swallowed the spider to catch the fly;
I don't know why she swallowed a fly - perhaps she'll die!
There was an old lady who swallowed a cat;
Fancy that to swallow a cat!
She swallowed the cat to catch the bird,
She swallowed the bird to catch the spider,
She swallowed the spider to catch the fly;
I don't know why she swallowed a fly - perhaps she'll die!`
	testSing(t, song, 3)
}

func TestSeventhSinger(t *testing.T){
	song := `There was an old lady who swallowed a fly.
I don't know why she swallowed a fly - perhaps she'll die!
There was an old lady who swallowed a spider;
That wriggled and wiggled and tickled inside her.
She swallowed the spider to catch the fly;
I don't know why she swallowed a fly - perhaps she'll die!
There was an old lady who swallowed a bird;
How absurd to swallow a bird.
She swallowed the bird to catch the spider,
She swallowed the spider to catch the fly;
I don't know why she swallowed a fly - perhaps she'll die!
There was an old lady who swallowed a cat;
Fancy that to swallow a cat!
She swallowed the cat to catch the bird,
She swallowed the bird to catch the spider,
She swallowed the spider to catch the fly;
I don't know why she swallowed a fly - perhaps she'll die!
There was an old lady who swallowed a dog;
What a hog, to swallow a dog!
She swallowed the dog to catch the cat,
She swallowed the cat to catch the bird,
She swallowed the bird to catch the spider,
She swallowed the spider to catch the fly;
I don't know why she swallowed a fly - perhaps she'll die!
There was an old lady who swallowed a cow;
I don't know how she swallowed a cow!
She swallowed the cow to catch the dog,
She swallowed the dog to catch the cat,
She swallowed the cat to catch the bird,
She swallowed the bird to catch the spider,
She swallowed the spider to catch the fly;
I don't know why she swallowed a fly - perhaps she'll die!
There was an old lady who swallowed a horse...
...She's dead, of course!`
	testSing(t, song, 6)
}

func TestChorusSing(t *testing.T){
	song := `
There was an old lady who swallowed a fly.
I don't know why she swallowed a fly - perhaps she'll die!
There was an old lady who swallowed a spider;
That wriggled and wiggled and tickled inside her.
She swallowed the spider to catch the fly;
I don't know why she swallowed a fly - perhaps she'll die!
There was an old lady who swallowed a bird;
How absurd to swallow a bird.
She swallowed the bird to catch the spider,
She swallowed the spider to catch the fly;
I don't know why she swallowed a fly - perhaps she'll die!
There was an old lady who swallowed a cat;
Fancy that to swallow a cat!
She swallowed the cat to catch the bird,
She swallowed the bird to catch the spider,
She swallowed the spider to catch the fly;
I don't know why she swallowed a fly - perhaps she'll die!
There was an old lady who swallowed a dog;
What a hog, to swallow a dog!
She swallowed the dog to catch the cat,
She swallowed the cat to catch the bird,
She swallowed the bird to catch the spider,
She swallowed the spider to catch the fly;
I don't know why she swallowed a fly - perhaps she'll die!
There was an old lady who swallowed a cow;
I don't know how she swallowed a cow!
She swallowed the cow to catch the dog,
She swallowed the dog to catch the cat,
She swallowed the cat to catch the bird,
She swallowed the bird to catch the spider,
She swallowed the spider to catch the fly;
I don't know why she swallowed a fly - perhaps she'll die!
There was an old lady who swallowed a horse...
...She's dead, of course!
	`
	c := newChorus()
	actual := c.Sing()
	assert.Equal(t, song, actual)
}

func testSing(t *testing.T, song string, index int) {
	c := newChorus()
	sb, writer := buildWriter()
	singer := c.singers[index]
	singer.Sing(writer)
	assert.Equal(t, song, sb.String())
}

