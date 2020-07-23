/**
 * @Author: jiehu
 * @Description:
 * @Project: test
 * @File:  xlsxClub
 * @CopyRight: fuxi
 * @Date: 2020/7/16 5:46 下午
 */
package xlsxHelper

import (
	"fmt"
	"github.com/360EntSecGroup-Skylar/excelize/v2"
)

const HeadStyle = `{
    "font":{
        "color":"#09600B"
    },
    "border":[
        {
            "style":1,
            "type":"left",
            "cocor":"#FF0000"
        },
        {
            "style":1,
            "type":"right",
            "cocor":"#FF0000"
        },
        {
            "style":1,
            "type":"top",
            "cocor":"#F0000"
        },
        {
			"style":1,
            "type":"bottom",
            "cocor":"#FF0000"
        }
    ]
	}`

type XlsxClub struct {
	File         *excelize.File
	StreamWriter *excelize.StreamWriter
	StyleHead    interface{}
	StyleBody    interface{}
}

func NewXlsxClub() *XlsxClub {
	return &XlsxClub{
		StyleHead: nil,
		StyleBody: nil,
	}
}

func (this *XlsxClub) Init(sheet string) error {
	var err error
	this.File = excelize.NewFile()
	this.File.NewSheet(sheet)
	//this.File.SetSheetName("sheet1", "test")
	this.File.DeleteSheet("sheet1")
	this.StreamWriter, err = this.File.NewStreamWriter(sheet)
	if err != nil {
		fmt.Println(err)
	}

	return err
}

func (this *XlsxClub) SetTitle(titles []interface{}) {
	styleId, err := this.File.NewStyle(HeadStyle)

	if err != nil {
		fmt.Println(err)
	}
	len := len(titles)
	rows := make([]interface{}, len)

	for k, v := range titles {
		var cell excelize.Cell
		cell.StyleID = styleId
		cell.Value = v
		rows[k] = cell
	}

	cell, _ := excelize.CoordinatesToCellName(1, 1)
	if err := this.StreamWriter.SetRow(cell, rows); err != nil {
		fmt.Println(err)
	}
}

func (this *XlsxClub) WriteRow(axis string, values []interface{}) error {
	return this.StreamWriter.SetRow(axis, values)
}

func (this *XlsxClub) FlushFile(fileName string) error {
	if err := this.StreamWriter.Flush(); err != nil {
		fmt.Println(err)
		return err
	}
	if err := this.File.SaveAs(fileName); err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
