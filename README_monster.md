# 数据结构实践报告

## 一、需求分析

1. 根据 ($\theta\delta_1..\delta_n\to\theta\delta_1..\theta\delta_n\theta$) 、$B\to tAdA$ 和 $A\to sae$ 的规则设计一个魔王语言解释器，将魔王语言解释为人可以听懂的语言。  
2. 测试数据：**B(ehnxgz)B** 解释为 **天上一只鹅地上一只鹅鹅追鹅赶鹅下鹅蛋鹅恨鹅天上一只鹅地上一只鹅**  


## 二、概要设计

1. 抽象数据类型设计
ADT stack&quene{
    数据对象：D = {$a_1..a_i$ | $a_i\in letter$, letter=A..Za..z}  
    数据关系：R = {[$D_1, D_n$]}  

    init()  
    程序开始之前，先初始化字母对应的汉字，以 map 存储，且为全局变量。  
    
    MonsterLanguage(lang string)  
    将 lang 转化为汉语  

    remark(q []string) []string  
    将给出的字符按照 ($\theta\delta_1..\delta_n\to\theta\delta_1..\theta\delta_n\theta$) 的规则进行转换。  

    push(q []string,s string)[]string  
    因为在该程序中的栈和队列均是使用切片进行构造的，所以他们的入队或者入栈操作类似，仅是出栈或者出队操作刚好相反，所以使用该函数同时进行入队和入栈操作。  
}

2. 主函数设计  
func main(){  
    MonsterLanguage("B(ehnxgz)B") //调用魔王语言解释程序  
}  

## 三、详细设计

```go
package main

import "fmt"

// 定义 A 和 B 代表的字符串
// 定义每个字符串对应的汉字
var AB map[string]string
var trans map[string]string

//初始化 AB 和 字符串对应的汉字
func init() {
	AB = make(map[string]string)
	AB["A"] = "sae"
	AB["B"] = "tsaedsae"
	trans = make(map[string]string)
	trans["t"] = "天"
	trans["d"] = "地"
	trans["s"] = "上"
	trans["a"] = "一只"
	trans["e"] = "鹅"
	trans["z"] = "追"
	trans["g"] = "赶"
	trans["x"] = "下"
	trans["n"] = "蛋"
	trans["h"] = "恨"
}

//根据传入的字符串构建对应的汉语并打印
func MonsterLanguage(lang string) {
	stack := make([]string, 0)
	qeuene := make([]string, 0)
	for i := 0; i < len(lang); i++ {
		s := string(lang[i])
		if s == ")" {
			var j int
			for j = len(stack) - 1; j >= 0 && string(stack[j]) != "("; j-- {
				qeuene = push(qeuene, stack[j])
			}
			stack = stack[:j]
			qeuene = remark(qeuene)
			for j := 0; j < len(qeuene); j++ {
				stack = push(stack, qeuene[j])
			}
		} else {
			stack = push(stack, s)
		}
	}
	for i := 0; i < len(stack); i++ {
		fmt.Print(trans[stack[i]])
	}
}

// 将括号内的字符按规则进行变换
func remark(qeuene []string) []string {
	if len(qeuene) == 0 {
		return nil
	}
	r := make([]string, 0)
	r = append(r, qeuene[len(qeuene)-1])
	for i := 0; i < len(qeuene)-1; i++ {
		r = append(r, qeuene[i])
		r = append(r, r[0])
	}
	return r
}

// 入队或者入栈操作
func push(qeuene []string, s string) []string {
	if s == "A" {
		for i := 0; i < len(AB[s]); i++ {
			qeuene = append(qeuene, string(AB[s][i]))
		}
	} else if s == "B" {
		for i := 0; i < len(AB[s]); i++ {
			qeuene = append(qeuene, string(AB[s][i]))
		}
	} else {
		qeuene = append(qeuene, s)
	}
	return qeuene
}

func main(){
    MonsterLanguage("B(ehnxgz)B")
}

```
## 五、用户手册

1. 本程序运行的环境为 Linux，go1.13.5。 
2. 打开 20191218 目录下的 main.go 文件，在 main 函数中，仅保留 MonsterLanguage("B(ehnxgz)B")，保存文件。  
3. 打开终端，进入 main.go 文件所在位置，执行 `go run main.go`  
4. 程序执行之后会打印对应的中文。  
5. 因为程序仅对 B()ehnxgz 这些字符做了中文转换，如果在字符串中添加了除上面所述之外的字符，可能会出现打印错误的现象。  

## 七、附录

20191218/
    MonsterLanguage.go  // 程序的主要代码
    go.mod  // 程序的依赖文件，不过因为本程序并未引入外部包，所以该文件可以忽略
    main.go // 程序入口
