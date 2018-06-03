package aca

import (
	"strings"
	"testing"
)

func Test_Hit_1(t *testing.T) {
	tree := NewTree()
	tree.AddKeyWords("he", "his", "she", "hers")
	tree.BuildTree()

	ts := []string{
		"shis",
		"shise",
		"hise",
		"ccavsjisgwhihiswe",
	}

	for _, key := range ts {
		if !tree.Hit(key) {
			t.Error(key, "should hit")
		}
	}

}

func Benchmark_Hit_1(b *testing.B) {
	b.StopTimer()

	tree := NewTree()
	tree.AddKeyWords("his")
	tree.BuildTree()

	b.StartTimer()
	key := "ccavsjisgwhihiswe"
	for i := 0; i < b.N; i++ {
		tree.Hit(key)
	}

}

func Benchmark_Hit_2(b *testing.B) {
	key := "ccavsjisgwhihiswe"
	for i := 0; i < b.N; i++ {
		strings.Contains(key, "his")
	}
}
