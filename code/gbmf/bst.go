package main

import (
	"encoding/json"
	"fmt"
	"math"
	"sort"
	"strconv"
	"strings"
)

type Node_h struct {
	freq  int
	char  rune
	left  *Node_h
	right *Node_h
}

// 建立霍夫曼树
func buildHuffmanTree(freq map[rune]int) *Node {
	var nodes []*Node_h
	for char, f := range freq {
		nodes = append(nodes, &Node_h{
			freq: f,
			char: char,
		})
	}

	for len(nodes) > 1 {
		sort.Slice(nodes, func(i, j int) bool {
			return nodes[i].freq < nodes[j].freq
		})

		left := nodes[0]
		right := nodes[1]
		parent := &Node_h{
			freq:  left.freq + right.freq,
			left:  left,
			right: right,
		}

		nodes = append(nodes[2:], parent)
	}

	return nodes[0]
}

// 记录霍夫曼编码路径
func recordHuffmanPath(node *Node, prefix string, path map[rune]string) {
	if node == nil {
		return
	}

	if node.char != 0 {
		path[node.char] = prefix
	}

	recordHuffmanPath(node.left, prefix+"0", path)
	recordHuffmanPath(node.right, prefix+"1", path)
}

// Node 定义节点.
type Node struct {
	value int   //
	left  *Node // 左子节点
	right *Node // 右子节点
}

// BST 是一个节点的值为int类型的二叉搜索树.
type BST struct {
	root *Node
}

// Insert 插入一个元素.
func (bst *BST) Insert(value int) {
	newNode := &Node{value, nil, nil}
	// 如果二叉树为空，那么这个节点就当作跟节点
	if bst.root == nil {
		bst.root = newNode
	} else {
		insertNode(bst.root, newNode)
	}
}

// 从根节点依次比较
func insertNode(root, newNode *Node) {
	if newNode.value < root.value { // 应该放到根节点的左边
		if root.left == nil {
			root.left = newNode
		} else {
			insertNode(root.left, newNode)
		}
	} else if newNode.value > root.value { // 应该放到根节点的右边
		if root.right == nil {
			root.right = newNode
		} else {
			insertNode(root.right, newNode)
		}
	}
	// 否则等于根节点
}

// Search 搜索元素(检查元素是否存在)
/*func (bst *BST) Search(value int) bool {
	return search(bst.root, value)
}
func search(n *Node, value int) bool {
	if n == nil {
		return false
	}
	if value < n.value {
		return search(n.left, value)
	}
	if value > n.value {
		return search(n.right, value)
	}
	count++
	return true
}*/

func search(root *Node, target int) (bool, int) {
	if root == nil {
		return false, 0
	}
	if root.value == target {
		return true, 1
	} else if root.value > target {
		found, count := search(root.left, target)
		return found, count + 1
	} else {
		found, count := search(root.right, target)
		return found, count + 1
	}
}

// Min 二叉搜索树中的最小值
func (bst *BST) Min() (int, bool) {
	return min(bst.root)
}
func min(node *Node) (int, bool) {
	if node == nil {
		return 0, false
	}
	n := node
	// 从左边找
	for {
		if n.left == nil {
			return n.value, true
		}
		n = n.left
	}
}

// Max 二叉搜索树中的最大值
func (bst *BST) Max() (int, bool) {
	return max(bst.root)
}
func max(node *Node) (int, bool) {
	if node == nil {
		return 0, false
	}
	n := node
	// 从右边找
	for {
		if n.right == nil {
			return n.value, true
		}
		n = n.right
	}
}

// PreOrderTraverse 前序遍历
func (bst *BST) PreOrderTraverse(a []int) {
	preOrderTraverse(bst.root, a)
}
func preOrderTraverse(n *Node, a []int) {
	if n != nil {
		a = append(a, n.value)
		fmt.Println(a) // 前
		preOrderTraverse(n.left, a)
		preOrderTraverse(n.right, a)
	}
}

//序列化
func (bst *BST) Serialize() string {
	return serialize(bst.root)

}
func serialize(root *Node) string {
	arr := []string{}
	var postOrder func(*Node)
	postOrder = func(node *Node) {
		if node == nil {
			return
		}
		postOrder(node.left)
		postOrder(node.right)
		arr = append(arr, strconv.Itoa(node.value))
	}
	postOrder(root)
	return strings.Join(arr, " ")
}

func deserialize(data string) *Node {
	if data == "" {
		return nil
	}
	arr := strings.Split(data, " ")
	var construct func(int, int) *Node
	construct = func(lower, upper int) *Node {
		if len(arr) == 0 {
			return nil
		}
		val, _ := strconv.Atoi(arr[len(arr)-1])
		if val < lower || val > upper {
			return nil
		}
		arr = arr[:len(arr)-1]
		return &Node{value: val, right: construct(val, upper), left: construct(lower, val)}
	}
	return construct(math.MinInt32, math.MaxInt32)
}

type Globalstatus struct {
	Name string
	Data string
}

func main() {
	newNode := &Node{5, nil, nil}
	newTree := BST{newNode}
	newTree.Insert(10)
	newTree.Insert(1)
	newTree.Insert(20)
	newTree.Insert(21)
	newTree.Insert(22)
	newTree.Insert(8)
	newTree.Insert(2)
	newTree.Insert(100)
	for i := 101; i < 200; i++ {
		newTree.Insert(i)

	}

	strtree := newTree.Serialize()
	fmt.Println(strtree)
	anotree := BST{deserialize(strtree)}
	fmt.Println(anotree.Serialize())
	fmt.Println(anotree)

	_, count := search(newTree.root, 105)
	fmt.Printf("比较了 %d 次\n", count)

	var treemap map[string]string
	treemap = make(map[string]string)
	treemap["tree"] = strtree
	var str5 string
	jsonStr, _ := json.Marshal(treemap)

	fmt.Println("string", string(jsonStr))
	var a = []byte(jsonStr)
	_ = json.Unmarshal(a, &str5)

	fmt.Println(jsonStr)
	fmt.Println(str5)

	var globalstatus Globalstatus
	bststatus := Globalstatus{
		Name: "bst",
		Data: "abc",
	}

	jsonStr2, _ := json.Marshal(bststatus)

	fmt.Println(string(jsonStr2))

	_ = json.Unmarshal(jsonStr2, &globalstatus)
	fmt.Println("gs", globalstatus)

}
