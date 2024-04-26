package main

import (
	"fmt"
	"github.com/spf13/cast"
	"github.com/xuri/excelize/v2"
)

func main() {
	var r []Result
	r = append(r, Result{CropName: "12", RobotID: "ddd", RobotName: "124", IsVisitor: true, IsIntent: false, IsAnswer: false, TransferCount: 4, ChatCount: 5})
	WriteExcel(r, r)
}

type Result struct {
	CropName      string `json:"crop_name"`
	RobotID       string `json:"robot_id"`
	RobotName     string `json:"robot_name"`
	IsVisitor     bool   `json:"is_visitor"`
	IsIntent      bool   `json:"is_intent"`
	IsAnswer      bool   `json:"is_answer"`
	TransferCount int    `json:"transfer_count"`
	ChatCount     int    `json:"chat_count"`
}

func WriteExcel(testScore, onlineScore []Result) {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	onlineScoreSheetName := "线上环境"
	testScoreSheetName := "测试环境"
	onlineScoreSheet := f.GetSheetName(0)
	f.SetSheetName(onlineScoreSheet, "线上环境")
	f.NewSheet(testScoreSheetName)

	var scores []Result
	scores = append(scores, onlineScore...)

	header := []string{"企业名称", "机器人ID", "机器人名称", "是否开启访客发送关键词时转人工", "是否开启根据用户语意转人工", "是否开启机器人答案评价不满意转人工", "近1周触发转人工次数", "近1周对话量"}

	f.SetSheetRow(onlineScoreSheetName, fmt.Sprintf("A%d", 1), &header)
	for i, row := range onlineScore {
		rowString := []string{row.CropName, row.RobotID, row.RobotName, boolText(row.IsVisitor), boolText(row.IsIntent),
			boolText(row.IsAnswer), cast.ToString(row.TransferCount), cast.ToString(row.ChatCount)}
		err := f.SetSheetRow(onlineScoreSheetName, fmt.Sprintf("A%d", i+2), &rowString)
		if err != nil {
			fmt.Printf("err:%+v\n", err)
		}
	}

	f.SetSheetRow(testScoreSheetName, fmt.Sprintf("A%d", 1), &header)
	for i, row := range testScore {
		rowString := []string{row.CropName, row.RobotID, row.RobotName, boolText(row.IsVisitor), boolText(row.IsIntent),
			boolText(row.IsAnswer), cast.ToString(row.TransferCount), cast.ToString(row.ChatCount)}
		err := f.SetSheetRow(testScoreSheetName, fmt.Sprintf("A%d", i+2), &rowString)
		if err != nil {
			fmt.Printf("err:%+v\n", err)
		}
	}

	if err := f.SaveAs("转人工相关功能客户数据拉取.xlsx"); err != nil {
		fmt.Println(err)
	}
}

func boolText(is bool) string {
	if is {
		return "是"
	}
	return "否"
}
