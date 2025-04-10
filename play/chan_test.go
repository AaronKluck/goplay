package play

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSelectorChan(t *testing.T) {
	var results []string

	var ch1 = make(chan string, 99)
	var ch2 = make(chan string, 99)

	// Test the SimpleChan function
	SelectorChan(
		NewPerson,
		NewCar,
		func(val Stringable) {
			ch1 <- val.Repr()
		},
		func(val Stringable) {
			ch2 <- val.Repr()
		},
	)
	close(ch1)
	close(ch2)
	for val := range ch1 {
		results = append(results, val)
	}
	for val := range ch2 {
		results = append(results, val)
	}
	assert.Equal(t, []string{
		"My name is Name-0 and I am 30 years old",
		"My name is Name-1 and I am 31 years old",
		"My name is Name-2 and I am 32 years old",
		"My name is Name-3 and I am 33 years old",
		"My name is Name-4 and I am 34 years old",
		"My name is Name-5 and I am 35 years old",
		"My name is Name-6 and I am 36 years old",
		"My name is Name-7 and I am 37 years old",
		"My name is Name-8 and I am 38 years old",
		"My name is Name-9 and I am 39 years old",
		"My 2010 Make-10 Model-10 goes brrrrr",
		"My 2011 Make-11 Model-11 goes brrrrr",
		"My 2012 Make-12 Model-12 goes brrrrr",
		"My 2013 Make-13 Model-13 goes brrrrr",
		"My 2014 Make-14 Model-14 goes brrrrr",
		"My 2015 Make-15 Model-15 goes brrrrr",
		"My 2016 Make-16 Model-16 goes brrrrr",
		"My 2017 Make-17 Model-17 goes brrrrr",
		"My 2018 Make-18 Model-18 goes brrrrr",
		"My 2019 Make-19 Model-19 goes brrrrr",
	}, results)
}
