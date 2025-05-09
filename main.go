package main

import (
	"algorithm/linkedlist"
)

func main() {

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
	
}
