package main

import "fmt"

// 反转字符串 同时还要去掉多余的空格
// 例： are   you   ok?  中文  => 中文 ok? you are
func main() {
	s := reverseString(" are   you   ok?  中文 ")
	fmt.Printf("___%s___", s)
}

var (
	runes []rune // 把字符串转为数组
)

// reverseString 传入待反转的字符串返回反转后的字符转
func reverseString(str string) string {
	beforeIsSpace := true
	for _, char := range []rune(str) {
		if string(char) == " " { // 扫描到空格
			if beforeIsSpace { // 如果之前是空格则直接继续
				continue
			} else { // 之前不是空格
				runes = append(runes, char)
				beforeIsSpace = true
			}
		} else { // 不是空格 添加进slice
			runes = append(runes, char)
			beforeIsSpace = false
		}
	}

	// 去除尾部可能有的空格
	for i := len(runes) - 1; i >= 0; i-- {
		//fmt.Printf("%c\n", runes[i])
		if string(runes[i]) == " " {
			runes = runes[:i]
		} else {
			break
		}
	}

	// 开始反转
	// 先反转整个字符串
	reverse(0, len(runes)-1)

	begin := -1
	for i, char := range runes {
		if string(char) == " " { // 扫描到空格 分词
			reverse(begin+1, i-1)
			begin = i
		}
		if i == len(runes) - 1 { // 到最后 需要反转最后一个单词
			reverse(begin+1, i)
		}
	}

	return string(runes)
}

func reverse(begin, end int) {
	start := begin
	last := end
	for last > start {
		runes[last], runes[start] = runes[start], runes[last]
		start++
		last--
	}
}
