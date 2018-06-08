// Aho–Corasick Algorithms
package aca

import (
	"fmt"
)

type ACATree struct {
	root *Node
	size int //树节点总数
	seq  int
}

type Node struct {
	subtree map[rune]*Node
	fail    *Node
	isEnd   bool //是否是可输出状态

	id         int
	value      string
	value_list string
}

func (a *ACATree) addNode() *Node {
	defer func() {
		a.seq++
	}()

	return &Node{
		subtree: make(map[rune]*Node),
		id:      a.seq,
	}
}

func NewTree(words_list ...string) (tree *ACATree) {
	tree = &ACATree{}
	tree.size = 1
	tree.root = tree.addNode()
	tree.AddKeyWords(words_list...)
	return
}

// 建立trie树
func (a *ACATree) AddKeyWords(words_list ...string) {
	for _, word := range words_list {
		pNode := a.root

		tmp := []rune(word)
		for i, ch := range tmp {
			value_list := pNode.value_list
			if pNode.subtree[ch] == nil {
				pNode.subtree[ch] = a.addNode()
				a.size += 1
			}

			pNode = pNode.subtree[ch]

			pNode.value = string(ch)
			pNode.value_list = value_list + pNode.value
			if i == len(tmp)-1 {
				pNode.isEnd = true // 单词末节点：设置为可输出状态
				fmt.Println(pNode.value_list)
			}
		}
	}
	fmt.Println("\nTotal Nodes =", a.size)
	a.BuildTree()
}

// 建里fail表
func (a *ACATree) BuildTree() {
	if a.size == 1 {
		// 说明只有一个root节点
		return
	}

	child_num := a.size
	list := make(chan *Node, a.size)
	list <- a.root
	child_num--

	for pCur := range list {
		// 广度优先遍历
		for _, son := range pCur.subtree {
			list <- son
			child_num--
			if child_num == 0 {
				close(list)
			}
		}

		for key, son := range pCur.subtree {
			fail := pCur.fail // 初始化为：父节点的fail指针
			for ; fail != nil; fail = fail.fail {
				if fail.subtree[key] != nil {
					son.fail = fail.subtree[key]
					break
				}
			}

			if fail == nil {
				son.fail = a.root
			}
			fmt.Printf("node[%d: %s: %v].fail = node[%d:%s]\n", son.id, son.value, son.isEnd, son.fail.id, son.fail.value)
		}
	}

}

// 命中可输出节点
func (a *ACATree) Hit(content string) bool {
	pNode := a.root
	for _, ch := range content {
		for pNode != nil {
			if pNode.subtree[ch] != nil {
				pNode = pNode.subtree[ch]
				if pNode.isEnd {
					// fmt.Printf(` content("%s")  `, pNode.value_list)
					return true
				}
				break
			} else {
				pNode = pNode.fail
			}
		}

		if pNode == nil {
			pNode = a.root
		}
	}
	return false
}
