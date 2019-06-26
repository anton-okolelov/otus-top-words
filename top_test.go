package top_words

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTop10Words(t *testing.T) {
	words := GetTop10Words("привет, привет, пока")
	assert.Equal(t, []string{"привет", "пока"}, words)
}

func TestLongText(t *testing.T) {
	words := GetTop10Words("привет, привет, пока, asdfasdf, asdfasdf, ,asdfasdfasdf,s df0jf0a9sjdfmasd, f,as dfa sd,fa sdf df df df ")
	assert.Equal(t,
		[]string{"df", "привет", "asdfasdf", "пока", "sjdfmasd", "sdf", "sd", "s", "jf", "fa"},
		words)
}
