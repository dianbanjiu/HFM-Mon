# 数据结构实践报告

## 一、需求分析

1. 构建哈夫曼编/译器  
2. 要求：  
   1. 初始化：根据键盘获取的数据初始化哈夫曼树，并将结果存储到 hfmTree
   2. 编码：利用构建好的哈夫曼树队 ToBeTran 中的数据进行编码，结果存储到 CodeFile 当中  
   3. 译码：利用构建好的哈夫曼树队 CodeFile 中的数据进行译码，结果存储在 TextFile 当中  
   4. 打印代码文件：将CodeFile 中的数据以每行 50 个字符的长度打印在终端，并将结果存储在 CodePrin 文件  
   5. 打印哈夫曼树：将创建好的哈夫曼树打印在终端，并将结果存储在 TreePrint 文件当中  
3. 测试数据：
   1. 待编译数据：THIS PROGRAM IS MY FAVRIATE
   2. 字符集及频度见 测试结果  


## 二、概要设计

1. 抽象数据类型设计  
ADT hfmTree {
    数据对象：D = {d,w,l,r,p | $d\in string, w,l,r,p\in int$}  
    数据关系：R = {<$d_1,w_1,l_1,r_1,p_1$>...}

    func InitHFMTree()  
    // 根据键盘获取的数据创建哈夫曼树，  
    // 并将结果写入文件  

    func CreHFMTreeFF() HFMTree   
    // 通过 hfmTree 还原对应的哈夫曼树  

    func Encoding()  
    //  利用构建好的哈夫曼树对数据进行编码  

    func Index(hfmtree HFMTree, c string) int  
    // 获取某个元素在哈夫曼树中的索引  

    func Print()  
    // 打印 CodeFile  

    func TreePrint()  
    // 打印 HFMTree  

    func Decoding()  
    // 使用已有的 HFMTree 对 CodeFile 进行译码  
}
2. 主函数设计  

func main(){  
    依次调用编译码的各个函数
}
## 三、详细设计

```go
package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

// 哈夫曼树的节点结构
type Node struct {
	Data                    string
	Weight                  int
	LChild, RChild, Parents int
}

// 哈夫曼树结构体
type HFMTree []Node

// 下面三个函数是为了对哈夫曼树以权值进行排序而需要实现的 sort.Sort 接口
// 获取切片的长度
func (h HFMTree) Len() int {
	return len(h)
}

// 索引为 i 的权值小于索引为 j 的权值
func (h HFMTree) Less(i, j int) bool {
	return h[i].Weight < h[j].Weight
}

// 交换两个元素的值
func (h HFMTree) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// 根据键盘获取的数据创建哈夫曼树，
// 并将结果写入文件
func InitHFMTree() {
	fmt.Println("InitHFMTree")
	fmt.Println("待输入字符集的大小为：")
	var n int
	fmt.Scanln(&n)

	t := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// 根据给定的字符集大小，初始化其数据项及权重
	var hfmTree = make(HFMTree, n)
	for i := 0; i < n; i++ {
		fmt.Printf("请输入第 %d 个点的字符数据及权重", i+1)
		var d string
		var w int
		fmt.Scanln(&d, &w)
		// 在获取字符“ ” 及其权值的时候，因为语言的设计，该项的 d 会等于 w， 而 w 会为 0,
		//所以添加下面的判断，当 d 不存在于 26 个字母中时，通过下面的转换，让 d 等于 “ ”，w 等于 d
		if !strings.Contains(t, d) {
			w, _ = strconv.Atoi(d)
			d = " "
		}
		var node = Node{
			Data:    d,
			Weight:  w,
			LChild:  -1,
			RChild:  -1,
			Parents: -1,
		}
		hfmTree[i] = node
	}

	// 根据权重对获取的元素进行排序
	sort.Sort(hfmTree)

	// 对获取元素的父节点索引进行赋值
	hfmTree[0].Parents = n
	for i := 0; i < n-1; i++ {
		hfmTree[i+1].Parents = n + i
	}

	// 创建哈夫曼树其余的 n-1 个节点
	node := Node{
		Data:    "#",
		Weight:  hfmTree[0].Weight + hfmTree[1].Weight,
		LChild:  0,
		RChild:  1,
		Parents: n + 1,
	}
	hfmTree = append(hfmTree, node)
	for i := 1; i < n-1; i++ {
		// w,l,r,p 分别代表新增节点的权重，左右孩子索引，双亲索引
		// 新增节点的数据项默认以一个 "#" 进行填充
		var w, l, r, p int

		w = hfmTree[n+i-1].Weight + hfmTree[i+1].Weight
		if hfmTree.Less(n+i-1, i+1) {
			l = n + i - 1
			r = i + 1
		} else {
			l = i + 1
			r = n + i - 1
		}
		if i != n-2 {
			p = n + i + 1
		} else {
			p = -1
		}
		node = Node{
			Data:    "#",
			Weight:  w,
			LChild:  l,
			RChild:  r,
			Parents: p,
		}
		hfmTree = append(hfmTree, node)
	}

	// 若 hfmTree 文件不存在，则新建
	var newFile *os.File
	_, err := os.Stat("hfmTree")
	if os.IsNotExist(err) {
		newFile, err = os.Create("hfmTree")
		if err != nil {
			log.Fatal(err)
		}
	}
	defer newFile.Close()

	// 将创建的哈夫曼树写入该文件
	for _, v := range hfmTree {
		_, err = newFile.Write([]byte(fmt.Sprintf("%v", v)))
		newFile.Write([]byte("\n"))
		if err != nil {
			log.Fatal(err)
		}
	}
}

// 通过 hfmTree 还原对应的哈夫曼树
func CreHFMTreeFF() HFMTree {
	hfmTreeFile, err := os.Open("hfmTree")
	if err != nil {
		log.Fatal(err)
	}
	defer hfmTreeFile.Close()
	t := "ABCDEFGHIJKLMNOPQRSTUVWXYZ#"

	// 逐行读取文件中的内容，每一行来对应于一个节点
	hfmTreeReader := bufio.NewReader(hfmTreeFile)
	data, _, err := hfmTreeReader.ReadLine()
	hfmtree := make(HFMTree, 0)
	for err == nil {
		var (
			d          string
			w, l, r, p int
		)
		// 类似于在创建哈夫曼树时的情形，空格项所在的所有元素可能会前移一个
		// 导致 d,w,l,r,p=w,l,r,p,0
		// 所以通过下面的判断将该项恢复到本来的形式
		fmt.Sscanf(string(data), "{%s %d %d %d %d}", &d, &w, &l, &r, &p)
		if !strings.Contains(t, d) {
			w, _ = strconv.Atoi(d)
			d = " "
			p, r = r, l
		}
		node := Node{
			Data:    d,
			Weight:  w,
			LChild:  l,
			RChild:  r,
			Parents: p,
		}
		hfmtree = append(hfmtree, node)
		data, _, err = hfmTreeReader.ReadLine()
	}
	return hfmtree
}

//  利用构建好的哈夫曼树对数据进行编码
func Encoding() {
	fmt.Println("Encoding")
	hfmtree := CreHFMTreeFF()
	encodeFile, err := os.Open("ToBeTran")
	if err != nil {
		log.Fatal(err)
	}
	defer encodeFile.Close()

	dataByte, err := ioutil.ReadAll(encodeFile)
	if err != nil {
		log.Fatal(err)
	}
	// 默认文件中仅有一行数据，如果有多行，可以逐行获取，然后创建一个 []string 来存储每一行的内容，下面的内容稍作修改即可
	data := string(dataByte)
	var encoding []int // 存储该行数据对应的哈夫曼编码
	// 外层循环遍历该行数据
	// 内层循环自底向上将每个字符编译为对应的哈夫曼编码
	for i := 0; i < len(data); i++ {
		c := string(data[i])
		index := Index(hfmtree, c)
		p := hfmtree[index].Parents
		letterCode := []int{}
		for p != -1 {
			if hfmtree[p].LChild == index {
				letterCode = append([]int{0}, letterCode...)
			} else {
				letterCode = append([]int{1}, letterCode...)
			}
			index = p
			p = hfmtree[index].Parents
		}
		encoding = append(encoding, letterCode...)
	}

	// 将编译好的数据写入到 CodeFile 中去
	s := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(encoding)), ""), "[]")
	sbyte := []byte(s)
	file, err := os.Create("CodeFile")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	file.Write(sbyte)
}

// 获取某个元素在哈夫曼树中的索引
func Index(hfmtree HFMTree, c string) int {
	var index int
	for i, v := range hfmtree {
		if v.Data == c {
			index = i
			break
		}
	}
	return index
}

// 打印 CodeFile
func Print() {
	fmt.Println("Print")
	file, err := os.Open("CodeFile")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	printFile, err := os.Create("CodePrin")
	if err != nil {
		log.Fatal(err)
	}
	defer printFile.Close()

	dataByte, _ := ioutil.ReadAll(file)
	data := string(dataByte)

	// 根据题目要求，每行 50 个字符元素
	d := len(data) % 50 // 判断所有的元素数量是否为 50 的整数倍
	n := len(data) / 50 // 均为 50 个字符的行数
	// 如果不是 50 的整数倍，则行数需要 +1
	if d != 0 {
		n += 1
	}
	for i := 0; i < n; i++ {
		var s, e int
		s = 50 * i
		e = 50*i + 50
		if e > len(data) {
			e = len(data)
		}
		dataSlice := []byte(data[s:e] + "\n")
		fmt.Println(data[s:e])
		printFile.Write(dataSlice)
	}
}

// 打印 HFMTree
func TreePrint() {
	fmt.Println("TreePrint")
	file, err := os.Open("hfmTree")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	treePrint, err := os.Create("TreePrint")
	if err != nil {
		log.Fatal(err)
	}
	defer treePrint.Close()

	hfmtree := CreHFMTreeFF()
	cur := make([]int, 0)
	cur = append(cur, len(hfmtree)-1)
	// 哈夫曼树的深度与非生成节点的数量相同
	// 生成节点的元素名为 #
	// 下为层序遍历，条件可以是深度，也可以以 next 数组的长度来代替，
	// 当 next 数组长度为空时终止循环
	for i := 0; i < (len(hfmtree)+1)/2; i++ {
		next := []int{}
		for j := 0; j < len(cur); j++ {
			if hfmtree[cur[j]].LChild != -1 && hfmtree[cur[j]].RChild != -1 { // 若该节点的左右孩子均存在，则将其加入到 next 数组
				next = append(next, hfmtree[cur[j]].LChild, hfmtree[cur[j]].RChild)
			} else if hfmtree[cur[j]].LChild == -1 && hfmtree[cur[j]].RChild != -1 { // 若仅右孩子存在，则将右孩子添加到 next 数组
				next = append(next, hfmtree[cur[j]].RChild)
			} else if hfmtree[cur[j]].LChild != -1 && hfmtree[cur[j]].RChild == -1 { // 若仅左孩子存在，则将左孩子加入到 next 数组
				next = append(next, hfmtree[cur[j]].LChild)
			}
			// 逐层打印哈夫曼树，并逐层将其添加到 TreePrint 文件当中
			fmt.Print("(", cur[j], " ", hfmtree[cur[j]].Weight, ")    ")
			s := "(" + string(cur[j]) + " " + string(hfmtree[cur[j]].Weight) + ")     "
			treePrint.Write([]byte(s))
		}
		fmt.Println()
		cur = next
		treePrint.Write([]byte("\n")) // 每层打印完成之后换行
	}
}

// 使用已有的 HFMTree 对 CodeFile 进行译码
func Decoding() {
	fmt.Println("Decoding")
	hfmtree := CreHFMTreeFF()
	encodeFile, err := os.Open("CodeFile")
	if err != nil {
		log.Fatal(err)
	}
	defer encodeFile.Close()
	encodeData, _ := ioutil.ReadAll(encodeFile)
	encodeDataString := string(encodeData)
	//encodeDataString = encodeDataString[1:]

	// 根据自顶向下的原则，从哈夫曼树的根开始与编码数据进行比较
	// 成功译码一个元素之后，将指向哈夫曼树的指针重新指向根节点
	var decodeData string
	i := 0
	for i < len(encodeDataString) {
		p := len(hfmtree) - 1
		for hfmtree[p].LChild != -1 && hfmtree[p].RChild != -1 && i < len(encodeDataString) {
			switch c, _ := strconv.Atoi(string(encodeDataString[i])); c {
			case 0:
				p = hfmtree[p].LChild
			case 1:
				p = hfmtree[p].RChild
			}
			i++
		}
		decodeData = decodeData + hfmtree[p].Data
	}
	decodeFile, err := os.Create("TextFile")
	if err != nil {
		log.Fatal(err)
	}
	defer decodeFile.Close()

	decodeFile.Write([]byte(decodeData))
}

func main(){
    InitHFMTree()
	TreePrint()
	Encoding()
	Print()
	Decoding()
}
```

## 五、用户手册

1. 本程序的执行环境为 Linux，go1.13.5。  
2. 进入 20191218,打开 main.go 文件，删除 MonsterLanguage() 所在行，保存退出。  
3. 在此处打开终端，执行 `go run main.go` 命令。  
4. 按照提示依次输入构建哈夫曼树所需的字符及其权重。  
5. 接着程序会创建并打印生成的哈夫曼树，并存储在 hfmTree 文件当中。  
6. 程序之后会对指定文件进行编码，编码文件存储于 CodeFile 文件当中，程序会以每行 50 个字符的长度打印生成的编码。  
7. 程序接着会对编码文件进行译码操作，并将译码生成的数据存储在 TextFile 文件当中。  


## 七、附录

20191218/  
    main.go //主程序入口  
    go.mod // 依赖文件，本程序并未引入外部包，可以忽略  
    Halfman.go // 程序的主要代码  
    ToBeTran    // 待编译的文件

    // 下面的文件均为根据实验要求生成的文件，可以删除，删除之后，再次执行程序后会再次生成  
    CodeFile    // 编译之后生成的编译文件  
    CodePrin    // 以每行 50 个字符长度打印的编译文件  
    hfmTree // 构建好的哈夫曼编码树  
    TextFile    // 译码之后数据的存储文件  
    TreePrin    // 哈夫曼树以层序存储的文件  
