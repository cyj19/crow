/**
 * @Author: cyj19
 * @Date: 2021/11/23 10:33
 */

// Package word 单词格式转换包
package word

import (
	"strings"
	"unicode"
)

// ToUpper 单词全部转为大写
func ToUpper(s string) string {
	return strings.ToUpper(s)
}

// ToLower 单词全部转为小写
func ToLower(s string) string {
	return strings.ToLower(s)
}

// UnderscoreToUpperCamelCase 下划线单词转为大写驼峰
func UnderscoreToUpperCamelCase(s string) string {
	// 把下划线替换为空格字符
	s = strings.Replace(s, "_", " ", -1)
	// 每个单词的首字母转为大写
	s = strings.Title(s)
	// 把空格字符替换为空
	return strings.Replace(s, " ", "", -1)
}

// UnderscoreToLowerCamelCase 下划线单词转为小写驼峰
func UnderscoreToLowerCamelCase(s string) string {
	// 把单词转为大写驼峰
	s = UnderscoreToUpperCamelCase(s)
	// 首字母转为小写
	return string(unicode.ToLower(rune(s[0]))) + s[1:]
}

// CamelCaseToUnderscore 驼峰转下划线
func CamelCaseToUnderscore(s string) string {
	var output []rune
	for i, r := range s {
		// 第一个字母转为小写
		if i == 0 {
			output = append(output, unicode.ToLower(r))
			continue
		}
		// 如果字母是大写，增加下划线
		if unicode.IsUpper(r) {
			output = append(output, '_')
		}
		// 字母转为小写
		output = append(output, unicode.ToLower(r))
	}

	return string(output)
}
