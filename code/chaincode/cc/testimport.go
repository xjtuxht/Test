package cc

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math"
	"math/big"
	mrand "math/rand"
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

type Globalstatus struct {
	Name string
	Data string
}

type User struct {
	Workplace string
	Major     string
	Gender    string
	Age       int //ciph
	//Phonenumber string
}

type User_info struct {
	Id          string
	Rep         int
	Value       int
	Phonenumber string
}

type LLPI struct {
	Workplace string
	Major     string
	Gender    string
}

type MLPI struct {
	Token []string
	Ciph  []string
}

type Up_info struct {
	Pubkey string
	Id     string
}

// InitLedger adds a base set of Users to the ledger
func (s *SmartContract) InitLedger(ctx contractapi.TransactionContextInterface) error {
	/*globalid := 0
	var idmap map[string]string
	idmap = make(map[string]string)
	idmap["id"] = strconv.Itoa(globalid)
	idStr, _ := json.Marshal(idmap)
	err := ctx.GetStub().PutState("idnumber", idStr)
	if err != nil {
		return fmt.Errorf("failed to put to world state. %v", err)
	}*/

	/*secuparam := 10
	mkey := Genmkey(secuparam)
	var parammap map[string]string
	parammap = make(map[string]string)
	parammap["mkey"] = strconv.Itoa(mkey)
	mapStr, _ := json.Marshal(parammap)
	err1 := ctx.GetStub().PutState("mkey", mapStr)
	if err1 != nil {
		return fmt.Errorf("failed to put to world state. %v", err1)
	}*/
	count := 10000
	newNode := &Node{1, nil, nil}
	newTree := BST{newNode}
	for i := 2; i <= count+1; i++ {
		newTree.Insert(i)
	}

	strtree := newTree.Serialize()
	fmt.Println("wtf", strtree)

	/*var treemap map[string]string
	treemap = make(map[string]string)
	treemap["tree"] = strtree*/
	bststatus := Globalstatus{
		Name: "bst",
		Data: strtree,
	}

	jsonStr, err := json.Marshal(bststatus)
	if err != nil {
		return err
	}

	err2 := ctx.GetStub().PutState(bststatus.Name, jsonStr)
	if err2 != nil {
		return fmt.Errorf("failed to put to world state. %v", err2)
	}

	return nil
}

/*func (s *SmartContract) Registry(ctx contractapi.TransactionContextInterface, Location string, Gender string, Age int) (string, error) {
	assetJSON, err := ctx.GetStub().GetState("idnumber")

}*/
func (s *SmartContract) Updatedata(ctx contractapi.TransactionContextInterface, data string) (bool, error) {

	up_info := Up_info{
		Pubkey: data,
		Id:     "1",
	}
	jsonup_info, err := json.Marshal(up_info)
	if err != nil {
		return false, err
	}

	err2 := ctx.GetStub().PutState("up_info", jsonup_info)
	if err2 != nil {
		return false, err2
	}

	return true, nil

}

func (s *SmartContract) Readstatus(ctx contractapi.TransactionContextInterface, key string) (*Globalstatus, error) {
	//start := time.Now() // 获取当前时间
	statusJSON, err := ctx.GetStub().GetState(key)
	if err != nil {
		return nil, fmt.Errorf("failed to read from world state: %v", err)
	}
	if statusJSON == nil {
		return nil, fmt.Errorf("the status %s does not exist", key)
	}

	var globalstatus Globalstatus
	err = json.Unmarshal(statusJSON, &globalstatus)
	if err != nil {
		return nil, err
	}
	//elapsed := time.Since(start)
	//fmt.Println("该函数执行完成耗时：", elapsed)

	return &globalstatus, nil
}

func (s *SmartContract) Registry(ctx contractapi.TransactionContextInterface, workplace string, gender string, major string, age int) (string, error) {
	llpi := LLPI{
		Workplace: workplace,
		Major:     gender,
		Gender:    major,
	}
	//str_age := strconv.Itoa(age)
	//mkey := strconv.Itoa(Genmkey(16, ctx.GetStub().GetTxTimestamp()))
	//mkey := "57946"
	//token, ciph := DerEnc(str_age, 10, 16, mkey)
	token:= []string{"ab", "bc", "bcd"}
	ciph:= []string{"app", "bna", "cjy"}
	mlpi := MLPI{
		Token: token,
		Ciph:  ciph,
	}

	jsonllpi, err := json.Marshal(llpi)
	if err != nil {
		return "", err
	}

	err2 := ctx.GetStub().PutState("llpi", jsonllpi)
	if err2 != nil {
		return "", err2
	}

	jsonmlpi, err := json.Marshal(mlpi)
	if err != nil {
		return "", err
	}

	err3 := ctx.GetStub().PutState("mlpi", jsonmlpi)
	if err3 != nil {
		return "", err3
	}

	id := "100"
	return id, nil

}

func (s *SmartContract) Matching(ctx contractapi.TransactionContextInterface, reputation int, age_l int, age_u int) error {
	s.Readstatus(ctx, "bst")
	num1 := 30
	num2 := 50
	if reputation < 100 {
		return nil
	}

	mkey := "57946"
	str_age_l := strconv.Itoa(age_l)
	str_age_u := strconv.Itoa(age_u)

	token_l, ciph_l := DerEnc(str_age_l, 10, 16, mkey)
	token_u, ciph_u := DerEnc(str_age_u, 10, 16, mkey)

	globalstatus, _ := s.Readstatus(ctx, "bst")
	bst_str := globalstatus.Data
	bst := BST{deserialize(bst_str)}
	_, time1 := search(bst.root, num1)
	_, time2 := search(bst.root, num2)

	for i := 0; i < time1; i++ {
		Cmp(ciph_l, token_l, ciph_u, 16)
	}
	for i := 0; i < time2; i++ {
		Cmp(ciph_u, token_u, ciph_l, 16)
	}

	return nil

}

func (s *SmartContract) SecondMatching(ctx contractapi.TransactionContextInterface, workplace string, gender string, major string, age int) (string, error) {
	mkey := "57946"
	llpi := LLPI{
		Workplace: workplace,
		Major:     gender,
		Gender:    major,
	}
	str_age := strconv.Itoa(age)
	str_age2 := strconv.Itoa(age + 5)

	//mkey := strconv.Itoa(Genmkey(16, ctx.GetStub().GetTxTimestamp()))
	token, ciph := DerEnc(str_age, 10, 16, mkey)
	_, ciph2 := DerEnc(str_age2, 10, 16, mkey)

	mlpi := MLPI{
		Token: token,
		Ciph:  ciph,
	}

	jsonllpi, err := json.Marshal(llpi)
	if err != nil {
		return "", err
	}

	err2 := ctx.GetStub().PutState("llpi", jsonllpi)
	if err2 != nil {
		return "", err2
	}

	jsonmlpi, err := json.Marshal(mlpi)
	if err != nil {
		return "", err
	}

	err3 := ctx.GetStub().PutState("mlpi", jsonmlpi)
	if err3 != nil {
		return "", err3
	}

	Cmp(ciph, token, ciph2, 16)

	pr := 100 * 0.8 / (100 + 0.8)
	var result string
	if pr > 50 {
		result = "yes"
	} else {
		result = "no"
	}

	return result, nil

}

func (s *SmartContract) Contact(ctx contractapi.TransactionContextInterface, pubkey string, id string) (bool, error) {


	/*downJSON, err := ctx.GetStub().GetState("bst")
	var down_info Up_info
	err = json.Unmarshal(downJSON, &down_info)
	if err != nil {
		return false, err
	}*/

	user_info := User_info{
		Id:          "1002",
		Rep:         100,
		Value:       1000,
		Phonenumber: "18612345678",
	}

	jsonuser_info, err := json.Marshal(user_info)
	if err != nil {
		return false, err
	}

	err3 := ctx.GetStub().PutState("user_info", jsonuser_info)
	if err3 != nil {
		return false, err3
	}

	return true, nil

}

func (s *SmartContract) Uppk(ctx contractapi.TransactionContextInterface, pubkey string, id string) (bool, error) {
	up_info := Up_info{
		Pubkey: pubkey,
		Id:     id,
	}
	jsonup_info, err := json.Marshal(up_info)
	if err != nil {
		return false, err
	}

	err2 := ctx.GetStub().PutState("up_info", jsonup_info)
	if err2 != nil {
		return false, err2
	}
	return true, nil

}


func (s *SmartContract) Feedback(ctx contractapi.TransactionContextInterface, score int) (bool, error) {
	user_info := User_info{
		Id:          "1002",
		Rep:         100,
		Value:       1000,
		Phonenumber: "18612345678",
	}
	var new_rep int
	if score == 0 {
		new_rep = user_info.Rep
	} else if score == 1 {
		new_rep = 2 * user_info.Rep
	} else {
		new_rep = 3 * user_info.Rep
	}
	new_user_info := User_info{
		Id:          "1002",
		Rep:         new_rep,
		Value:       1000,
		Phonenumber: "18612345678",
	}
	jsonuser_info, err := json.Marshal(new_user_info)
	if err != nil {
		return false, err
	}

	err3 := ctx.GetStub().PutState("user_info", jsonuser_info)
	if err3 != nil {
		return false, err3
	}
	return true, nil

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

//func main() {}
//generate bn, n=32
func RandKey() *big.Int {
	n, err := rand.Int(rand.Reader, big.NewInt(math.MaxInt32))
	if err == nil {
		//fmt.Println("rand.Int：", n, n.BitLen())
		return n
	} else {
		return nil
	}
}

func Genmkey(digit int, time int64) int {
	mrand.Seed(time)
	k := float64(digit) //security param
	Range := int(math.Pow(2, k))
	mkey := mrand.Intn(Range)
	return mkey
}

func HashA(key string, data string, k int) string {
	hmac := hmac.New(md5.New, []byte(key))
	hmac.Write([]byte(data))
	out := hex.EncodeToString(hmac.Sum([]byte(""))[0 : k/8])
	//fmt.Println(out)
	//fmt.Println(hmac.Sum([]byte("")), "hmac")

	return out
}

func HashB(key string, data string, k int) string {
	hmac := hmac.New(sha1.New, []byte(key))
	hmac.Write([]byte(data))
	out := hex.EncodeToString(hmac.Sum([]byte(""))[0 : k/8])
	//fmt.Println(out)
	return out
}

func HashC(key string, data string, k int) string {
	hmac := hmac.New(sha256.New, []byte(key))
	hmac.Write([]byte(data))
	out := hex.EncodeToString(hmac.Sum([]byte(""))[0 : k/8])
	//fmt.Println(out)

	return out
}

func DerEnc(num string, n int, k int, mkey string) ([]string, []string) {
	data1, err := strconv.Atoi(num)
	var token []string
	var flist []string
	if err == nil {
		data2 := int64(data1)
		strdata := strconv.FormatInt(data2, 2) //binary
		length := strings.Count(strdata, "") - 1
		//fmt.Println(strdata, "strdata") //
		b := 0
		//mkey := strconv.Itoa(Genmkey(k)) //k bit mkey
		for i := n; i >= length; i-- {
			d := HashA(mkey, "0", k)
			token = append(token, d)
		}
		//fmt.Println("for111") //
		for i := length - 1; i >= 1; i-- {
			digit, _ := strconv.Atoi(strdata[i : i+1])
			b = b*2 + digit
			d := HashA(mkey, strconv.Itoa(b), k)
			token = append(token, d)
		}
		//fmt.Println("for222") //

		randi := strconv.Itoa(57946)
		for i := n - 1; i >= 0; i-- {
			str1 := HashC(token[n-i-1], randi, k)
			value1, _ := strconv.Atoi(str1)
			str2 := HashB(mkey, token[n-i-1], k)
			value2, _ := strconv.Atoi(str2)

			b = 0
			if i >= n-length {
				//fmt.Println(i) //
				b, _ = strconv.Atoi(strdata[i+length-n : i+length-n+1])
			}
			f := (value1 + value2 + b) % 3
			flist = append(flist, strconv.Itoa(f)) //n-1,0

		}
		//fmt.Println("for333") //

		flist = append(flist, randi) //reversed ciph
		//fmt.Println(flist, "flist")  //

	}
	return token, flist

}

func Cmp(ciph1 []string, token1 []string, ciph2 []string, k int) int {
	n := len(token1)
	c := 0
	for i := 0; i <= n-1; i++ {
		f1, _ := strconv.Atoi(ciph1[i+1])
		f2, _ := strconv.Atoi(ciph2[i+1])
		h1 := HashC(token1[i], ciph1[n], k)
		h1i, _ := strconv.ParseInt(h1, 16, 32)
		h2 := HashC(token1[i], ciph2[n], k)
		h2i, _ := strconv.ParseInt(h2, 16, 32)

		c = (f1 - f2 - int(h1i) + int(h2i)) % 3
		if c != 0 {
			return c

		}
	}
	return c

}
