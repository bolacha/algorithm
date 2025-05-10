package main

import (
	"algorithm/linkedlist"
	tries "algorithm/tries"
	"algorithm/tries_rune"
	"fmt"
)

func main() {

	linkedTest()
	triesTest()
	triesRuneTest()
}

func linkedTest() {
	fmt.Println("Linked Test")
	linkedList := linkedlist.LinkedList{}

	linkedList.Append(10)
	linkedList.Append(1)
	linkedList.PrintList()
	linkedList.Append(10)
	linkedList.Append(2)
	linkedList.PrintList()
	linkedList.Append(10)
	linkedList.Append(3)
	linkedList.PrintList()

	linkedList.DeleteWithValue(1)
	linkedList.PrintList()
	linkedList.DeleteWithValue(2)
	linkedList.PrintList()
	linkedList.DeleteWithValue(3)
	linkedList.PrintList()

	fmt.Println("---------------------")
}

func triesTest() {
	fmt.Println("Tries Test")
	words := []string{"a", "chocolate", "potato", "mango"}

	root := tries.GetNode()

	for i := 0; i < len(words); i++ {
		tries.Insert(root, words[i])
	}

	fmt.Println("contains the work chocolate", tries.Search(root, "chocolate"))
	fmt.Println("contains the work man", tries.Search(root, "man"))
	fmt.Println("contains the work potato", tries.Search(root, "chocolate"))
	fmt.Println("---------------------")
}

func triesRuneTest() {
	fmt.Println("Tries Rune Test")

	words := []string{"a", "chocolate", "potato", "mango"}

	root := tries_rune.GetNode()

	for i := 0; i < len(words); i++ {
		tries_rune.Insert(root, words[i])
	}

	fmt.Println("contains the work chocolate", tries_rune.Search(root, "chocolate"))
	fmt.Println("contains the work man", tries_rune.Search(root, "man"))
	fmt.Println("contains the work potato", tries_rune.Search(root, "chocolate"))

	fmt.Println("---------------------")
}
