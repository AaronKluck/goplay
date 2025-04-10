package play

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThing(t *testing.T) {
	var results []string

	var ch = make(chan string, 1000)

	// Test the SimpleChan function
	SelectorChan(
		NewPerson,
		NewCar,
		func(val Stringable) {
			ch <- val.Repr()
		},
		func(val Stringable) {
			ch <- val.Repr()
		},
	)
	close(ch)
	for val := range ch {
		results = append(results, val)
	}
	assert.Equal(t, results, []string{
		"My name is Name-0 and I am 30 years old",
		"My 2000 Make-0 Model-0 goes brrrrr",
		"My name is Name-1 and I am 31 years old",
		"My 2001 Make-1 Model-1 goes brrrrr",
		"My name is Name-2 and I am 32 years old",
		"My 2002 Make-2 Model-2 goes brrrrr",
		"My name is Name-3 and I am 33 years old",
		"My 2003 Make-3 Model-3 goes brrrrr",
		"My name is Name-4 and I am 34 years old",
		"My 2004 Make-4 Model-4 goes brrrrr",
		"My name is Name-5 and I am 35 years old",
		"My 2005 Make-5 Model-5 goes brrrrr",
		"My name is Name-6 and I am 36 years old",
		"My 2006 Make-6 Model-6 goes brrrrr",
		"My name is Name-7 and I am 37 years old",
		"My 2007 Make-7 Model-7 goes brrrrr",
		"My name is Name-8 and I am 38 years old",
		"My 2008 Make-8 Model-8 goes brrrrr",
		"My name is Name-9 and I am 39 years old",
		"My 2009 Make-9 Model-9 goes brrrrr",
		"My name is Name-10 and I am 40 years old",
		"My 2010 Make-10 Model-10 goes brrrrr",
	})
}
