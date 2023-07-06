package main

import (
	"fmt"
	"time"
)

func main() {
	// 当前时间
	now := time.Now()

	// 上周的起始时间和结束时间
	lastWeekStart := now.AddDate(0, 0, int(-7-now.Weekday()+1))
	lastWeekEnd := lastWeekStart.AddDate(0, 0, 6)

	// 上个月的起始时间和结束时间
	lastMonthStart := time.Date(now.Year(), now.Month()-1, 1, 0, 0, 0, 0, time.Local)
	lastMonthEnd := lastMonthStart.AddDate(0, 1, -1)

	// 下周的起始时间和结束时间
	nextWeekStart := now.AddDate(0, 0, int(7-now.Weekday()+1))
	nextWeekEnd := nextWeekStart.AddDate(0, 0, 6)

	// 下个月的起始时间和结束时间
	nextMonthStart := time.Date(now.Year(), now.Month()+1, 1, 0, 0, 0, 0, time.Local)
	nextMonthEnd := nextMonthStart.AddDate(0, 1, -1)

	// 去年的起始时间和结束时间
	lastYearStart := time.Date(now.Year()-1, 1, 1, 0, 0, 0, 0, time.Local)
	lastYearEnd := lastYearStart.AddDate(1, 0, -1)

	// 明年的起始时间和结束时间
	nextYearStart := time.Date(now.Year()+1, 1, 1, 0, 0, 0, 0, time.Local)
	nextYearEnd := nextYearStart.AddDate(1, 0, -1)

	// 输出结果
	fmt.Println("Last week:", lastWeekStart, "-", lastWeekEnd)
	fmt.Println("Last month:", lastMonthStart, "-", lastMonthEnd)
	fmt.Println("Next week:", nextWeekStart, "-", nextWeekEnd)
	fmt.Println("Next month:", nextMonthStart, "-", nextMonthEnd)
	fmt.Println("Last year:", lastYearStart, "-", lastYearEnd)
	fmt.Println("Next year:", nextYearStart, "-", nextYearEnd)
}
