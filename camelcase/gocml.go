/**
 * @Author: jiehu
 * @Description:
 * @Project: test
 * @File:  gocml
 * @CopyRight: fuxi
 * @Date: 2020/7/23 11:32 上午
 */
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
	"unicode"
)

const (
	LOWER_CAMEL_CASE = iota //小驼峰
	UPPER_CAMEL_CASE        //大驼峰
	LINE_CAMEL_CASE         //下划线
)

// 内嵌bytes.Buffer，支持连写
type Buffer struct {
	*bytes.Buffer
}

func NewBuffer() *Buffer {
	return &Buffer{Buffer: new(bytes.Buffer)}
}

func (b *Buffer) Append(i interface{}) *Buffer {
	switch val := i.(type) {
	case int:
		b.append(strconv.Itoa(val))
	case int64:
		b.append(strconv.FormatInt(val, 10))
	case uint:
		b.append(strconv.FormatUint(uint64(val), 10))
	case uint64:
		b.append(strconv.FormatUint(val, 10))
	case string:
		b.append(val)
	case []byte:
		b.Write(val)
	case rune:
		b.WriteRune(val)
	}
	return b
}

func (b *Buffer) append(s string) *Buffer {
	defer func() {
		if err := recover(); err != nil {
			log.Println("*****内存不够了！******")
		}
	}()
	b.WriteString(s)
	return b
}

// 驼峰式写法转为下划线写法
func Camel2Case(name string) string {
	buffer := NewBuffer()
	for i, r := range name {
		if unicode.IsUpper(r) {
			if i != 0 {
				buffer.Append('_')
			}
			buffer.Append(unicode.ToLower(r))
		} else {
			buffer.Append(r)
		}
	}
	return buffer.String()
}

// 下划线写法转为驼峰写法
func Case2Camel(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}

// 首字母大写
func Ucfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToUpper(v)) + str[i+1:]
	}
	return ""
}

// 首字母小写
func Lcfirst(str string) string {
	for i, v := range str {
		return string(unicode.ToLower(v)) + str[i+1:]
	}
	return ""
}

func readFile(filename string) []byte {
	// 使用ioutil一次性读取文件
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("read file err:", err.Error())
		return nil
	}

	// 打印文件内容
	//fmt.Println(string(data))

	return data
}

func readFileLine(inname string, outname string) {
	outfile, err := os.OpenFile(outname, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		fmt.Println(err)
	}
	defer outfile.Close()

	// 读取一个文件的内容
	file, err := os.Open(inname)
	if err != nil {
		fmt.Println("open file err:", err.Error())
		return
	}

	// 处理结束后关闭文件
	defer file.Close()

	// 使用bufio读取
	r := bufio.NewReader(file)

	for {
		// 分行读取文件  ReadLine返回单个行，不包括行尾字节(\n  或 \r\n)
		//data, _, err := r.ReadLine()

		// 以分隔符形式读取,比如此处设置的分割符是\n,则遇到\n就返回,且包括\n本身 直接返回字符串
		//str, err := r.ReadString('\n')

		// 以分隔符形式读取,比如此处设置的分割符是\n,则遇到\n就返回,且包括\n本身 直接返回字节数数组
		data, err := r.ReadBytes('\n')
		// 打印出内容
		fmt.Printf("%v", string(data))
		// 读取到末尾退出
		if err == io.EOF {
			break
		}
		// 读取到末尾退出
		if err == io.EOF {
			break
		}

		if bytes.Index(data, []byte("json:")) == -1 {
			log.Printf("-- %v", string(data))
			if _, err = outfile.Write(data); err != nil {
				fmt.Println(err)
			}
		} else {
			t := replace(data)
			log.Printf("++ %v", string(t))
			if _, err = outfile.Write(t); err != nil {
				fmt.Println(err)
			}
		}
	}
}

func AppendBytes(pBytes ...[]byte) []byte {
	len := len(pBytes)

	var buffer bytes.Buffer
	for i := 0; i < len; i++ {
		buffer.Write(pBytes[i])
	}

	return buffer.Bytes()
}

func replace(data []byte) []byte {
	start := bytes.Index(data, []byte("json:\""))
	//第一部分的数据
	firstP := data[:start+6]
	//log.Printf("replace start: %d", start)

	//第二部分的数据，需要转化
	tmp := data[start+6:]
	//log.Printf("replace tmp: %s", string(tmp))
	endp := bytes.Index(tmp, []byte("\""))
	if endp == -1 {
		//说明输入的数据有错误，不符合要求
		return data
	}
	//log.Printf("replace endp: %d", endp)

	secondP := tmp[:endp]
	//log.Printf("replace secondP: %v", string(secondP))

	var mid []byte
	//第二部分有可能存在其他数据，需要继续过滤
	tI := bytes.Index(secondP, []byte(","))
	if tI != -1 {
		mid = secondP[tI:]
		secondP = tmp[:tI]
	}

	secondPR := gocml(LOWER_CAMEL_CASE, string(secondP))

	//第三部分的数据
	thrdP := tmp[endp:]

	//合并数据
	return AppendBytes(firstP, []byte(secondPR), mid, thrdP)
}

//func GetStruct(body []byte) {
//	leftCnt := bytes.Count(body, []byte("{"))
//	RightCnt := bytes.Count(body, []byte("}"))
//	log.Printf("------leftCnt: %d, rightCnt:%d", leftCnt, RightCnt)
//
//	strs := make([]string, leftCnt)
//
//	head := body
//	for i := 0; i < leftCnt; i++ {
//		h := bytes.Index(head, []byte("{"))
//		e := bytes.Index(head, []byte("}"))
//		log.Printf("-----h:%d, e:%d", h, e)
//
//		strs[i] = string(head[h:e+1])
//
//		head = body[e+2:]
//		//log.Printf("----- head:%s", string(head))
//	}
//
//
//	for k, v := range strs{
//		log.Printf("k: %d, v: %s", k, v)
//	}
//	//var structs []string
//	//for i:=0; i< len(body);i++{
//	//	strings.Index()
//	//}
//}

func gocml(format int, code string) string {
	switch format {
	case LOWER_CAMEL_CASE:
		return Lcfirst(Case2Camel(code))
	case UPPER_CAMEL_CASE:
		return Ucfirst(Case2Camel(code))
	case LINE_CAMEL_CASE:
	//todo
	default:
		//todo
	}

	return ""
}

func main() {
	if len(os.Args) < 3 {
		fmt.Print("Usage of gocml\n")
		fmt.Print("gocml [inputFile] [outputFile]\n")
		return
	}
	//println(gocml(LOWER_CAMEL_CASE, "aaa_bbb_cc"))
	//println(gocml(UPPER_CAMEL_CASE, "aaa_bbb_cc"))

	//var src string
	//flag.StringVar(&src, "src", "nihao", "source file")
	//
	//level := flag.Int("level", 0, "debug level")
	//
	//var memo string
	//flag.StringVar(&memo, "memo", "", "the memory")
	//
	//log.Printf("------src: %s", src)
	//log.Printf("------level: %d", *level)
	//log.Printf("------momo: %s", memo)

	log.Printf("==========args: %s", os.Args[1])
	//body := readFile(os.Args[1])
	//GetStruct(body)
	readFileLine(os.Args[1], os.Args[2])

	//readFile()
	//GetStruct();
}
