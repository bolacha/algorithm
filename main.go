package main

import (
	"algorithm/linkedlist"
	"algorithm/minheap"
	"algorithm/queue"
	stack "algorithm/stack"
	tries "algorithm/tries"
	"algorithm/tries_rune"
	"fmt"
)

func main() {

	linkedTest()
	triesTest()
	triesRuneTest()

	stackTest()

	queueTest()

	minHeapTest()
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

func stackTest() {
	fmt.Println("Stack Test")

	var st stack.Stack

	st.Push("bolacha")
	st.Push("banana")
	st.Push("Dulce")

	for !st.IsEmpty() {
		value, removed := st.Pop()

		if removed {
			fmt.Println(value)
		}
	}

	fmt.Println("---------------------")
}

func queueTest() {
	fmt.Println("Queue Test")

	var qu queue.Queue

	qu.Enqueue("bolacha")
	qu.Enqueue("banana")
	qu.Enqueue("Dulce")

	for !qu.IsEmpty() {
		value, removed := qu.Dequeue()

		if removed {
			fmt.Println(value)
		}
	}

	fmt.Println("---------------------")
}

func minHeapTest() {
	fmt.Println("Min Heap Test")

	mh := &minheap.MinHeap{}

	mh.Insert(35)
	mh.Insert(11)
	mh.Insert(17)
	mh.Insert(23)

	mh.Print()

	removed := mh.Remove()
	fmt.Printf("\nRemoved %d\n", removed)
	mh.Print()

	mh.Insert(6)
	fmt.Println("\nInserted 6")
	mh.Print()

	fmt.Println("---------------------")
}
