package main

import (
	"fmt"
	material "mysqlang/material"
)

/*					operator表
+----+-------+------+-------+
| id | money | rate | note  |
+----+-------+------+-------+
|  1 |   100 |  0.5 | hi    |
|  2 |   200 |    1 | hello |
|  3 |   300 |  1.5 | hey   |
+----+-------+------+-------+
*/

func main() {
	var res []Res
	material.MyDb.Table("operator").
		Select("id, money+rate AS plus, money*rate AS multiple").Where("id=?", 1).Find(&res)
	//SELECT id, money+rate AS plus, money*rate AS multiple FROM `operator` WHERE id=1
	fmt.Println("结果=", res) //结果= [{1 100.5 50}]

	var result []Result
	material.MyDb.Table("operator").
		Select("id, money * 9 AS multiple, note + '-end' AS str, CONCAT(note, '-last') AS plus_str").
		Where("id=?", 1).Find(&result)
	//SELECT id, money * 9 AS multiple, note + '-end' AS str, CONCAT(note, '-last') AS plus_str FROM `operator` WHERE id=1
	fmt.Println("id=", result[0].Id) //id=1
	fmt.Println("积=", result[0].Multiple) //积=900
	fmt.Println("串相加=", result[0].Str) //串相加=0  //todo:因此可见 + 号在MySQL中不能用于连接字符串！只能用CONCAT函数。
	fmt.Println("串相连=", result[0].PlusStr) //串相连= hi-last
}
type Res struct {
	Id int
	Plus float64
	Multiple int
}

type Result struct {
	Id int
	Multiple int
	Str string
	PlusStr string
}

var tabSql =`CREATE TABLE guan1lian2gorm.operator (
  id int NOT NULL AUTO_INCREMENT,
  money int NOT NULL,
  rate float NOT NULL,
  note varchar(30),
  PRIMARY KEY (id)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;`

var dataSql = "INSERT INTO operator VALUES (0, 100, 0.5, 'hi'), (0, 200, 1, 'hello'), (0, 300, 1.5, 'hey');"

func init() {
	material.GetDB()
	material.MyDb.Exec("DROP TABLE IF EXISTS operator;")
	material.MyDb.Exec(tabSql)
	material.MyDb.Exec(dataSql)
}