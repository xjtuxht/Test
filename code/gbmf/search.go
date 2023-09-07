package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func search(root *Node, target int) (bool, int) {
	if root == nil {
		return false, 0
	}
	if root.Val == target {
		return true, 1
	} else if root.Val > target {
		found, count := search(root.Left, target)
		return found, count + 1
	} else {
		found, count := search(root.Right, target)
		return found, count + 1
	}
}

func main() {
	root := &Node{Val: 5, Left: &Node{Val: 3, Left: &Node{Val: 2}, Right: &Node{Val: 4}}, Right: &Node{Val: 7, Left: &Node{Val: 6}, Right: &Node{Val: 8}}}
	target := 4
	found, count := search(root, target)
	if found {
		fmt.Printf("找到了，比较了 %d 次\n", count)
	} else {
		fmt.Println("没找到")
	}
}
