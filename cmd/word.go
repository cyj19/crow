/**
 * @Author: vagary
 * @Date: 2021/11/23 9:28
 */

// Package cmd 自定义命令行工具集
// 单词格式转换命令
package cmd

import (
	"github.com/spf13/cobra"
	"github.com/vagaryer/yucli/internal/word"
	"log"
	"strings"
)

// 转换模式枚举值
const (
	ModeUpper                      = iota + 1 // 全部转大写
	ModeLower                                 // 全部转小写
	ModeUnderscoreToUpperCamelCase            // 下划线转大写驼峰
	ModeUnderscoreToLowerCamelCase            // 下划线转小写驼峰
	ModeCamelCaseToUnderscore                 // 驼峰转下划线
)

// 子命令完整说明
var desc = strings.Join([]string{
	"该子命令提供各种单词转换格式，模式如下：",
	"1: 全部转大写",
	"2: 全部转小写",
	"3: 下划线转大写驼峰",
	"4: 下划线转小写驼峰",
	"5: 驼峰转下划线",
}, "\n")

var name string
var mode int8

// 单词转换命令
var wordCmd = &cobra.Command{
	Use:   "word",
	Short: "单词格式转换",
	Long:  desc,
	Run: func(cmd *cobra.Command, args []string) {
		var content string
		switch mode {
		case ModeUpper:
			content = word.ToUpper(name)
		case ModeLower:
			content = word.ToLower(name)
		case ModeUnderscoreToUpperCamelCase:
			content = word.UnderscoreToUpperCamelCase(name)
		case ModeUnderscoreToLowerCamelCase:
			content = word.UnderscoreToLowerCamelCase(name)
		case ModeCamelCaseToUnderscore:
			content = word.CamelCaseToUnderscore(name)
		default:
			log.Fatal("暂不支持该转换模式，请执行 help word 查看帮助文档")
		}

		log.Printf("输出结果：%s \n", content)
	},
}

func init() {
	// 定义标志位
	wordCmd.Flags().StringVarP(&name, "name", "n", "", "请输入单词内容")
	wordCmd.Flags().Int8VarP(&mode, "mode", "m", 0, "请输入单词转换的模式")
}
