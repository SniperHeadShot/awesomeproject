package util

import (
	"github.com/tealeg/xlsx"
)

// 构建Excel头样式
func BuildExcelHeadStyle() *xlsx.Style {
	style := xlsx.NewStyle()
	return style
}

// 构建Excel正文样式
func BuildExcelBodyStyle() *xlsx.Style {
	style := xlsx.NewStyle()
	return style
}

// 生成表头
func BuildSheetHead(sheet *xlsx.Sheet, excelHead *[]string) {
	row := sheet.AddRow()
	channelSize := len(*excelHead)
	for i := 0; i < channelSize; i++ {
		cell := row.AddCell()
		cell.SetStyle(BuildExcelHeadStyle())
		cell.Value = (*excelHead)[i]
	}
}
