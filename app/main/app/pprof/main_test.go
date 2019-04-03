package main

import (
	"regexp"
	"strings"
	"testing"
)

var haystack = `Lorem ipsum dolor sit amet, consectetur adipicing auctor elit.
Nullam maximus `

// go test -bench=. -cpuprofile cpu.out

func BenchmarkSubstring(b *testing.B) {
	for i := 0; i < b.N; i++ {
		strings.Contains(haystack, "au1ctor")
	}
}

func BenchmarkRegex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		regexp.MatchString("auctor", haystack)
	}
}

//Precompilation

var reg = regexp.MustCompile("auctor")

func BenchmarkRegex2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reg.MatchString(haystack)
	}
}

func BenchmarkGenereatedRegex(b *testing.B) {
	for i := 0; i < b.N; i++ {
		matchAuctor(haystack)
	}
}
