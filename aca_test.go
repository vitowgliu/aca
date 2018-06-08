package aca

import (
	"strings"
	"testing"
)

func Test_Hit_1(t *testing.T) {
	tree := NewTree("he", "his", "she", "hers")
	// tree.AddKeyWords()
	// tree.BuildTree()

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

func Test_Hit_2(t *testing.T) {
	tree := NewTree("测试2", "测试", "测试3", "测试4", "亲爱", "第")

	ts := []string{
		"第一！！！",
		"一天一口一个亲爱的对方",
		"儿子爸爸也能拿第一",
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
