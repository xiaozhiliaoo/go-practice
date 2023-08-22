package main

import (
	"fmt"

	"github.com/xuri/excelize/v2"
)

func main() {
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	// Create a new sheet.
	index, _ := f.NewSheet("Sheet2222")

	sheetName := f.GetSheetName(0)
	sheetName1 := f.GetSheetName(1)
	f.SetSheetName(sheetName, "模型配置")
	f.SetSheetName(sheetName1, "模型配置2")

	// Set value of a cell.
	f.SetCellValue("模型配置", "A2", "Hello world.")
	f.SetCellValue("模型配置2", "B2", 100)
	// Set active sheet of the workbook.
	f.SetActiveSheet(index)
	// Save spreadsheet by the given path.
	if err := f.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}
