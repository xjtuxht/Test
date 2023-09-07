package main

import (
	"encoding/json"
	"fmt"
	"math"
	"strconv"
	"strings"

	"github.com/hyperledger/fabric-contract-api-go/contractapi"
)

// SmartContract provides functions for managing an Asset
type SmartContract struct {
	contractapi.Contract
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

// InitLedger adds a base set of Users to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface, value int) error {
	newNode := &Node{value, nil, nil}
	newTree := BST{newNode}
	newTree.Insert(10)
	strtree := newTree.Serialize()
	var treemap map[string]string
	treemap = make(map[string]string)
	treemap["tree"] = strtree
	jsonStr, _ := json.Marshal(treemap)
	err := ctx.GetStub().PutState("bst", jsonStr) //problem
	if err != nil {
		return fmt.Errorf("failed to put to world state. %v", err)
	}
	return nil
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

func main() {}
