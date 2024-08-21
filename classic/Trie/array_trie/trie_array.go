package arraytrie

// TrieNode 前缀树节点
type TrieNode struct {
	pass  int
	end   int
	nexts [26]*TrieNode // 固定数组 实现路
}

// NewTrie 创建前缀树节点
func NewTrie() *TrieNode {
	return &TrieNode{
		pass:  0,
		end:   0,
		nexts: [26]*TrieNode{}, // make([]*TrieNode, 26),
	}
}

var root = NewTrie()

func (t *TrieNode) insert(word string) {
	node := root
	node.pass++

	// 从左往右遍历字符
	for i := 0; i < len(word); i++ {
		// 'a' -> 'a' - 'a' -> 0
		// 'z' -> 'z' - 'a' -> 25
		path := int(byte(word[i]) - 'a')
		if node.nexts[path] == nil {
			node.nexts[path] = NewTrie()
		}
		node = node.nexts[path]
		node.pass++
	}
	node.end++
}

// search 查询前缀树中，word单词出现了几次
func (t *TrieNode) search(word string) int {
	node := root

	for i := 0; i < len(word); i++ {
		path := int(byte(word[i]) - 'a')
		if node.nexts[path] == nil {
			return 0
		}
		node = node.nexts[path]
	}
	return node.end
}

// prefixNumber 查询前缀树中，有多少单词以 prefix开头
func (t *TrieNode) prefixNumber(prefix string) int {
	node := root

	for i := 0; i < len(prefix); i++ {
		path := int(byte(prefix[i]) - 'a')

		if node.nexts[path] == nil {
			return 0
		}
		node = node.nexts[path]
	}
	return node.pass
}

// 如果之前word插入过前缀树，那么此时删掉一次
// 如果之前word没有插入过前缀树，那么什么都不做
func (t *TrieNode) delete(word string) {
	if t.search(word) > 0 {
		node := root
		node.pass--
		for i := 0; i < len(word); i++ {
			path := int(byte(word[i]) - 'a')
			node.nexts[path].pass--
			if node.nexts[path].pass == 0 {
				node.nexts[path] = nil
				return
			}
			node = node.nexts[path]
		}
		node.end--
	}
}
