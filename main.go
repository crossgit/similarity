package main

import (
	"log"
	"math"
	"strings"

	"github.com/huichen/sego"
)

func main1() {
	str1 := `曾任天津市金融工作局副局长。2015年9月，免去李克强天津市金融工作局副局长职务。`
	str2 := `现任广西玉林市陆川县小城镇建设有限公司副总经理。 	`

	txt1 := []byte(strings.Replace(str1, string(rune(32)), "", -1))
	txt2 := []byte(strings.Replace(str2, string(rune(32)), "", -1))

	m1 := SeGoFunc(txt1)
	m2 := SeGoFunc(txt2)
	log.Println("\r\n", m1, "\r\n", m2)

	f := CosCalculate(m1, m2)
	log.Println(f)
	// ----------------
	var c1 StrCos
	var c2 StrCos
	var c3 StrCos
	c1.Code = "abc"
	c2.Code = "bbb"
	c3.Code = "ccc"
	c1.Cos = 0.38765
	c1.Cos = 0.48765
	c1.Cos = 0.58765

	var arr = []StrCos{c1, c2, c3}
	var tmp StrCos

	for _, v := range arr {
		if v.Cos >= tmp.Cos && v.Cos > 0.8 {
			tmp = v
		}
	}
	log.Println(tmp.Code)
}

func SeGoFunc(txt []byte) map[string]int {
	segments := segmenter.Segment(txt)
	arr := sego.SegmentsToSlice(segments, false)
	mapArr := make(map[string]int)
	for ok, info := range arr {
		log.Println(ok, info)
		mapArr[info]++
	}
	return mapArr
	// log.Println(sego.SegmentsToString(segments, true))
}

var segmenter sego.Segmenter

func init() {
	segmenter.LoadDictionary("./dictionary.txt")
}

// AppendMapKey 将后面参数的map的key和前面的map合并在一起,不存在的按0值补充
func AppendMapKey(t, after map[string]int) map[string]int {
	for k := range after {
		if _, exsit := t[k]; !exsit {
			t[k] = 0
		}
	}
	return t
}

// CosCalculate 两个map的余弦公式计算
func CosCalculate(m1, m2 map[string]int) float64 {
	// 将两个map转换成key一样的map,并只留下value
	m1 = AppendMapKey(m1, m2)

	top := 0
	buttom1 := 0
	buttom2 := 0

	for k, v1 := range m1 {
		buttom1 += v1 * v1
		if _, ex := m2[k]; ex {
			top += v1 * m2[k]
			buttom2 += m2[k] * m2[k]
		}
	}
	return float64(top) / (math.Sqrt(float64(buttom1)) * math.Sqrt(float64(buttom2)))
}

// StrCos 存放姓名编号和相似值
// var arr []StrCos
// var tmp StrCos
// for k,v := range arr{
// 	if v.Cos>tmp.Cos {
//		tmp = v
//}
//}
// return tmp.Code
type StrCos struct {
	Cos  float64
	Code string
}
