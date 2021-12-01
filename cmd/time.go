/**
 * @Author: cyj19
 * @Date: 2021/12/1 8:55
 */

package cmd

import (
	"github.com/cyj19/crow/internal/timer"
	"github.com/spf13/cobra"
	"log"
	"strconv"
	"strings"
	"time"
)

var calculateTime, duration string

// 时间相关的根命令
var timeCmd = &cobra.Command{
	Use:   "time",
	Short: "时间计算和格式转换",
	Long:  "时间计算和格式转换",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

var nowTimeCmd = &cobra.Command{
	Use:   "now",
	Short: "获取当前时间",
	Long:  "获取当前时间",
	Run: func(cmd *cobra.Command, args []string) {
		layout := "2006-01-02 15:04:05"
		nowTime := timer.GetNowTime()
		log.Printf("输出结果：%s, %d \n", nowTime.Format(layout), nowTime.Unix())
	},
}

var calculateTimeCmd = &cobra.Command{
	Use:   "calc",
	Short: "计算所需时间",
	Long:  "计算所需时间",
	Run: func(cmd *cobra.Command, args []string) {
		var currentTime time.Time
		layout := "2006-01-02 15:04:05"
		location, _ := time.LoadLocation("Asia/Shanghai")
		if calculateTime == "" {
			currentTime = timer.GetNowTime()
		} else {
			var err error
			// 通过空格数判断输入时间的格式
			space := strings.Count(calculateTime, " ")
			if space == 0 {
				layout = "2006-01-02"
			}
			// 使用带时区的转换
			currentTime, err = time.ParseInLocation(layout, calculateTime, location)
			if err != nil {
				// 时间字符串转为时间戳
				t, _ := strconv.Atoi(calculateTime)
				currentTime = time.Unix(int64(t), 0)
			}
		}
		t, err := timer.GetCalculateTime(currentTime, duration)
		if err != nil {
			log.Fatalf("timer.GetCalculateTime err: %v", err)
		}

		log.Printf("输出结果：%s, %d \n", t.Format(layout), t.Unix())
	},
}

func init() {
	timeCmd.AddCommand(nowTimeCmd)
	timeCmd.AddCommand(calculateTimeCmd)
	// 设置命令行标志
	calculateTimeCmd.Flags().StringVarP(&calculateTime, "calculateTime", "c", "",
		`需要计算的时间，有效单位为时间戳或已格式化后的时间`)
	calculateTimeCmd.Flags().StringVarP(&duration, "duration", "d", "",
		`持续时间， 有效单位为"ns", "us", "s", "m", "h"`)
}
