package tries_rune

type TrieNode struct {
	children   map[rune]*TrieNode
	endOfWords bool
}

func GetNode() *TrieNode {
	return &TrieNode{
		children:   make(map[rune]*TrieNode),
		endOfWords: false,
	}
}

func Insert(root *TrieNode, key string) {
	temp := root

	for _, ch := range key {
		if child, ok := temp.children[ch]; ok {
			temp = child
		} else {
			newChild := &TrieNode{children: make(map[rune]*TrieNode)}
			temp.children[ch] = newChild
			temp = newChild
		}
	}

	temp.endOfWords = true
}

func Search(root *TrieNode, key string) bool {
	temp := root

	for _, ch := range key {
		if child, ok := temp.children[ch]; ok {
			temp = child
			continue
		}
		return false
	}
	return temp.endOfWords
}
