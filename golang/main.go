package main

import (
	"fmt"
	"strings"
)

var chorus = []*singer{
	NewSinger("fly", ""),
	NewSinger("spider", "That wriggled and wiggled and tickled inside her."),
}

func formUp(){
	for i, s := range chorus {
		if i > 0 {
			s.sibling = chorus[i-1]
		}
	}
}

type singer struct {
	name   string
	overtureTemplate string
	writer func(song string)
	sibling *singer
}

func NewSinger(name string, overtureTemplate string) *singer {
	return &singer{name: name, overtureTemplate:overtureTemplate}
}

func (s *singer) Sing(writer func(sing string)) {
	s.writer = writer
	if s.sibling != nil{
		s.sibling.Sing(writer)
		s.writer("\n")
	}

	s.selfIntroduceSing()
	s.preludeSing()
	s.chorusSing()
	s.postludeSing()
}

func (s *singer) selfIntroduceSing() {
	s.writer(fmt.Sprintf("There was an old lady who swallowed a %s", s.name))
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
		s.writer(fmt.Sprintf("She swallowed the %s to catch the %s;", s.name, s.sibling.name))
		s.writer("\n")
		s.sibling.chorusSing()
	}
}
func (s *singer) postludeSing() {
	s.writer("I don't know why she swallowed a fly - perhaps she'll die!")
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
