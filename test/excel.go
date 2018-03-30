package main

import (
	"fmt"

	//"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/tealeg/xlsx"
)

//type ExcelMap struct {
//	Dict map[string]string
//}

//type Record struct {
//	Eng string
//	Zh  string
//}
//func ReadExcel(filePath string) {
//	filePath := "D:/hyf/gopath/src/test/excel.xlsx"
//	xlsx, err := xlsx.OpenFile(filePath)
//	if err != nil {
//		fmt.Println(err)
//	}
//	for _, sheet := range xlsx.Sheets {
//		for _, row := range sheet.Rows {
//			for _, cell := range row.Cells {
//				text := cell.String()
//				fmt.Printf("%s\n", text)
//			}
//		}
//	}
//}

func main() {
	//ExcelMap := make(map[string]string)
	FilePath := "D:/hyf/gopath/src/test/excel.xlsx"
	xlsx, err := xlsx.OpenFile(FilePath)
	if err != nil {
		fmt.Println(err)
	}
	//循环文件中所有工作表
	for _, sheet := range xlsx.Sheets {
		//循环对应工作表中行数
		for _, row := range sheet.Rows {
			row_result := make(map[int]string)
			//循环工作表行数的每一单元格
			for k, cell := range row.Cells {
				fmt.Print(cell)
				row_result[k] = cell.Value
				fmt.Println(row_result)
			}

		}
	}
	//	DaTa := make(map[string]string)

	//	var Recipient string = ""
	//	for {
	//		if Recipient == "E;" {
	//			break
	//		} else {
	//			fmt.Println("请输入需要查询的英文单词")
	//			fmt.Scanln(&Recipient)
	//			var Output = DaTa[Recipient]
	//			Output, ok := DaTa[Recipient]
	//			if ok {
	//				fmt.Println(Output)
	//			} else {
	//				fmt.Println("未找到")
	//			}
	//		}

	//	}
}
