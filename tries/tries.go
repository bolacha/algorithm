package tries

const Alphabet_Size = 26

type TrieNode struct {
	children   [Alphabet_Size]*TrieNode
	endOfWords bool
}

func GetNode() *TrieNode {
	node := &TrieNode{}
	node.endOfWords = false

	for i := 0; i < Alphabet_Size; i++ {
		node.children[i] = nil
	}

	return node
}

func Insert(root *TrieNode, key string) {
	temp := root

	for i := 0; i < len(key); i++ {
		index := key[i] - 'a' // 'a' is a rune , decimal 97

		if temp.children[index] == nil {
			temp.children[index] = GetNode()
		}

		temp = temp.children[index]
	}

	temp.endOfWords = true
}

func Search(root *TrieNode, key string) bool {
	temp := root

	for i := 0; i < len(key); i++ {
		index := key[i] - 'a'
		if temp.children[index] != nil {
			temp = temp.children[index]
		} else {
			return false
		}
	}

	return temp != nil && temp.endOfWords
}
