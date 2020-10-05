package main

import (
    "github.com/360EntSecGroup-Skylar/excelize"
    "fmt"

)

func check(e error){
  if e != nil{
    fmt.Println(e.Error())
    return
  }
}

func main(){
  f, err := excelize.OpenFile("Aprendizaje ganado Bob.xlsx")
  if err != nil {
    fmt.Println(err.Error())
    return
  }
/*Obetener un sólo valor de celda.
  cell, err := f.GetCellValue("aprendizaje ganado","B3")
  if err != nil{
    fmt.Println(err.Error())
    return
  }
*/
  fmt.Println(cell)
  rows, err := f.GetRows("aprendizaje ganado")
  for _, row := range rows {
    for _, colCell := range row{
      print(colCell, "\t")
    }
    fmt.Println()
  }
}
