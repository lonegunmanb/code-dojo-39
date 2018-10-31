package main

import (
	"fmt"
	"strings"
)

type singer struct {
	name             string
	overtureTemplate string
	writer           func(song string)
	sibling          *singer
	deadly           bool
}

type chorus struct {
	singers []*singer
}

func newChorus() *chorus{
	var singers = []*singer{
		newSinger("fly", ""),
		newSinger("spider", "That wriggled and wiggled and tickled inside her."),
		newSinger("bird", "How absurd to swallow a %s."),
		newSinger("cat", "Fancy that to swallow a %s!"),
		newSinger("dog", "What a hog, to swallow a %s!"),
		newSinger("cow", "I don't know how she swallowed a %s!"),
		newSinger("horse", ""),
	}
	c := &chorus{singers}
	c.FormUp()
	return c
}

func (c *chorus) FormUp() {
	for i, s := range c.singers {
		if i > 0 {
			s.sibling = c.singers[i-1]
		}
		if i == len(c.singers)-1 {
			c.singers[i].deadly = true
		}
	}
}

func (c *chorus) Sing() string{
	sb, writer := buildWriter()
	writer("\n")
	c.singers[len(c.singers)-1].Sing(writer)
	writer("\n	")
	return sb.String()
}

func newSinger(name string, overtureTemplate string) *singer {
	return &singer{name: name, overtureTemplate: overtureTemplate}
}

func (s *singer) Sing(writer func(sing string)) {
	s.writer = writer
	s.siblingSing(writer)
	s.selfIntroduceSing()
	if !s.finalSing() {
		s.preludeSing()
		s.chorusSing()
		s.postludeSing()
	}
}

func (s *singer) siblingSing(writer func(sing string)) {
	if s.sibling != nil {
		s.sibling.Sing(writer)
		s.writer("\n")
	}
}

func (s *singer) finalSing() bool {
	if s.deadly {
		s.writer("...She's dead, of course!")
	}
	return s.deadly
}

func (s *singer) selfIntroduceSing() {
	s.writer(fmt.Sprintf("There was an old lady who swallowed a %s", s.name))
	if s.deadly {
		s.writer("...\n")
		return
	}
	delimiter := ";"
	if s.sibling == nil {
		delimiter = "."
	}
	s.writer(delimiter)
	s.writer("\n")
}

func (s *singer) preludeSing() {
	if s.sibling != nil {
		preludeSong := ""
		if strings.Contains(s.overtureTemplate, "%s") {
			preludeSong = fmt.Sprintf(s.overtureTemplate, s.name)
		} else {
			preludeSong = s.overtureTemplate
		}
		s.writer(preludeSong)
		s.writer("\n")
	}
}

func (s *singer) chorusSing() {
	if s.sibling != nil {
		s.writer(fmt.Sprintf("She swallowed the %s to catch the %s", s.name, s.sibling.name))
		delimiter := ","
		if s.sibling.sibling == nil {
			delimiter = ";"
		}
		s.writer(delimiter)
		s.writer("\n")
		s.sibling.chorusSing()
	}
}
func (s *singer) postludeSing() {
	s.writer("I don't know why she swallowed a fly - perhaps she'll die!")
}

func buildWriter() (*strings.Builder, func(sing string)) {
	stringBuilder := &strings.Builder{}
	writer := func(sing string) {
		stringBuilder.WriteString(sing)
	}
	return stringBuilder, writer
}

func main() {
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
	fmt.Printf(song)
}
