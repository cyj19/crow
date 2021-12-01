/**
 * @Author: cyj19
 * @Date: 2021/12/1 8:52
 */

// Package timer 时间计算和格式转换处理
package timer

import "time"

// 设置本地时区
var location, _ = time.LoadLocation("Asia/Shanghai")

func GetNowTime() time.Time {
	return time.Now().In(location)
}

func GetCalculateTime(currentTime time.Time, d string) (time.Time, error) {
	duration, err := time.ParseDuration(d)
	if err != nil {
		return time.Time{}, err
	}
	return currentTime.Add(duration).In(location), nil
}
