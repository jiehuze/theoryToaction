/**
 * @Author: jiehu
 * @Description:
 * @Project: test
 * @File:  xxx
 * @CopyRight: fuxi
 * @Date: 2020/7/16 4:07 下午
 */
package main

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
	"math/rand"

	"theoryToaction/xlsx/xlsxHelper"
)

func test() {
	file := excelize.NewFile()
	streamWriter, err := file.NewStreamWriter("Sheet1")
	if err != nil {
		fmt.Println(err)
	}
	styleID, err := file.NewStyle(`{"font":{"color":"#777777"}}`)
	if err != nil {
		fmt.Println(err)
	}
	if err := streamWriter.SetRow("A1", []interface{}{
		excelize.Cell{StyleID: styleID, Value: "Data"}}); err != nil {
		fmt.Println(err)
	}
	for rowID := 2; rowID <= 102400; rowID++ {
		row := make([]interface{}, 50)
		for colID := 0; colID < 50; colID++ {
			row[colID] = rand.Intn(640000)
		}
		cell, _ := excelize.CoordinatesToCellName(1, rowID)
		if err := streamWriter.SetRow(cell, row); err != nil {
			fmt.Println(err)
		}
	}
	if err := streamWriter.Flush(); err != nil {
		fmt.Println(err)
	}
	if err := file.SaveAs("Book1.xlsx"); err != nil {
		fmt.Println(err)
	}
}

func main() {
	//xlsx := excelize.NewFile()
	//// Create a new sheet.
	//index := xlsx.NewSheet("测试")
	//// Set value of a cell.
	//xlsx.SetCellValue("Sheet2", "A2", "Hello world.")
	//xlsx.SetCellValue("Sheet1", "B2", 100)
	//// Set active sheet of the workbook.
	//xlsx.SetActiveSheet(index)
	//// Save xlsx file by the given path.
	//err := xlsx.SaveAs("./Book1.xlsx")
	//if err != nil {
	//	fmt.Println(err)
	//}

	//test()

	club := xlsxHelper.NewXlsxClub()

	club.Init("测试")

	row := make([]interface{}, 10)
	for colID := 0; colID < 10; colID++ {
		row[colID] = "你好"
	}

	club.SetTitle(row)

	club.FlushFile("book.xlsx")
}
