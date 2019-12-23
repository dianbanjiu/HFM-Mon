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
