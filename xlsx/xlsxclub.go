///**
// * @Author: jiehu
// * @Description:
// * @Project: test
// * @File:  xlsxclub
// * @CopyRight: fuxi
// * @Date: 2020/7/16 3:44 下午
// */
package main
//
//import (
//"log"
//"os"
//
////"github.com/extrame/xls"
//"github.com/tealeg/xlsx"
//)
//
//var xlsxTitle = []string{"字段1", "字段2", "字段3", "字段4"}
//var cell *xlsx.Cell
//
//func main() {
//
//	pwd, _ := os.Getwd()
//	targetPath := pwd + `\result.xlsx`
//	xlsxFile := getXlsxFile(targetPath)
//	xlsxSheet := xlsxFile.Sheets[0]
//
//	dataToInsert := []string{"切片内数据1", "切片内数据2", "切片内数据3", "切片内数据4"}
//	insertRow(xlsxSheet, &dataToInsert)
//
//	xlsxFile.Save(targetPath)
//
//}
//
//// 将一个切片指针对应的数据插入到xlsx.sheet中
//func insertRow(sheet *xlsx.Sheet, rowDataPtr *[]string) {
//	row := sheet.AddRow()
//	rowData := *rowDataPtr
//	for _, v := range rowData {
//		cell := row.AddCell()
//		cell.Value = v
//	}
//
//}
//
//// 获取xlsx.File对象的指针，如果文件路径不存在则新建一个文件，并返回其指针
//
//func getXlsxFile(filePath string) *xlsx.File {
//	var file *xlsx.File
//	if _, err := os.Stat(filePath); err == nil {
//		file, err = xlsx.OpenFile(filePath)
//		if err != nil {
//			log.Fatal(err)
//		}
//
//	} else {
//		file = xlsx.NewFile()
//		sheet, err := file.AddSheet("sheet1")
//		if err != nil {
//			log.Fatal(err)
//		}
//		insertRow(sheet, &xlsxTitle)
//	}
//
//	return file
//
//}
