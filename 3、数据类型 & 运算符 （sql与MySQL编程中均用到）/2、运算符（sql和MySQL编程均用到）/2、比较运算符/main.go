package main

import (
	"fmt"
	"mysqlang/material"
)

/*
+----+-----+------+-------+
| id | usr | age  | note  |
+----+-----+------+-------+
|  1 | a   |   10 | hi    |
|  2 | b   | NULL | hello |
|  3 | c   |   12 | NULL  |
+----+-----+------+-------+
*/

func main() {
	//----------------------------------------------  =  ---------------------------------------------------
	var res1 Cp
	material.MyDb.Table("compare").
		Select("id, age=10 AS one_zero").Where("id=?", 1).Find(&res1)
	//SELECT id, age=10 AS one_zero FROM `compare` WHERE id=1
	fmt.Println("res1=", res1) //res1= {1 1}

	//----------------------------------------------  !=  ---------------------------------------------------
	var res2 Cp
	material.MyDb.Table("compare").
		Select("id, usr != 'x' AS one_zero").Where("id=?", 1).Find(&res2)
	//SELECT id, usr != 'x' AS one_zero FROM `compare` WHERE id=1
	fmt.Println("res2=", res2)//res2= {1 1}

	//----------------------------------------------  <  ---------------------------------------------------
	var res3 Cp
	material.MyDb.Table("compare").
		Select("id, id<2 AS one_zero").Where("id=?", 2).Find(&res3)
	//SELECT id, id<2 AS one_zero FROM `compare` WHERE id=2
	fmt.Println("res3=", res3)//res3= {2 0}

	//----------------------------------------------  BETWEEN AND  -------------------------------------------
	var res4 Cp
	material.MyDb.Table("compare").
		Select("id, age BETWEEN 10 AND 11 AS one_zero").Where("id=?", 3).Find(&res4)
	//SELECT id, age BETWEEN 10 AND 11 AS one_zero FROM `compare` WHERE id=3
	fmt.Println("res4=", res4)//res4= {3 0}

	//----------------------------------------------  IN  ---------------------------------------------------
	var res5 Cp
	material.MyDb.Table("compare").
		Select("id, usr IN ('a', 'b', 'c') AS one_zero").Where("id=?", 1).Find(&res5)
	//SELECT id, usr IN ('a', 'b', 'c') AS one_zero FROM `compare` WHERE id=1
	fmt.Println("res5=", res5)//res5= {1 1}

	//----------------------------------------------  IS NULL  ----------------------------------------------
	var res6 Cp
	material.MyDb.Table("compare").
		Select("id, age IS NULL AS one_zero").Where("id=?", 2).Find(&res6)
	//SELECT id, age IS NULL AS one_zero FROM `compare` WHERE id=2
	fmt.Println("res6=", res6)//res6= {2 1}

	//----------------------------------------------  LIKE  ---------------------------------------------------
	var res7 Cp
	material.MyDb.Table("compare").
		Select("id, note LIKE 'h%' AS one_zero").Where("id=?", 2).Find(&res7)
	//SELECT id, note LIKE 'h%' AS one_zero FROM `compare` WHERE id=2
	fmt.Println("res7=", res7)//res7= {2 1}

	//----------------------------------------------  REGEXP  ---------------------------------------------------
	var res8 Cp
	material.MyDb.Table("compare").
		Select("id, note REGEXP '^h' AS one_zero").Where("id=?", 1).Find(&res8)
	//SELECT id, note REGEXP '^h' AS one_zero FROM `compare` WHERE id=1
	fmt.Println("res8=", res8)//res8= {1 1}

}

type Cp struct {
	Id      int
	OneZero int
}

var tabSql = `CREATE TABLE guan1lian2gorm.compare (
  id int NOT NULL AUTO_INCREMENT,
  usr varchar(30) NOT NULL,
  age int,
  note varchar(20),
  PRIMARY KEY (id)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=latin1;`

//todo:可以看看下面一行中的null
var dataSql = "INSERT INTO compare VALUES (0, 'a', 10, 'hi'), (0, 'b', null, 'hello'), (0, 'c', 12, null);"

func init() {
	material.GetDB()
	material.MyDb.Exec("DROP TABLE IF EXISTS compare;")
	material.MyDb.Exec(tabSql)
	material.MyDb.Exec(dataSql)
}
