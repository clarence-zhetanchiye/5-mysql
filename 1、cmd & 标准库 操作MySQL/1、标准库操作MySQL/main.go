//TODO:标准库操作MySQL

package main

import (
	"database/sql"
	"fmt"

	//todo:引入下面的包且加一占位符，目的是让该包中的driver.go中83行的init注册函数运行，
	// 从而将“mysql”注册到本标准库中，方便本文件中main函数第一行的 sql.Open 到标准库中能
	// 找到"mysql"，从而使得本标准库中源码可以调用该包中的接口。
	_ "github.com/go-sql-driver/mysql"
)

//TODO:打开数据库的sql.Open()函数的第二个参数的值的最完整形式是
// root:413188ok@tcp(localhost:3306)/shi3yong4gorm?charset=utf8mb4&parseTime=True&loc=Local

func main() {
	mydb, err := sql.Open("mysql", "root:413188ok@tcp(localhost:3306)/test")
	if err != nil {
		fmt.Println(err)
		return
	}
	//Exec方法用于对数据库进行写（增、改、删）操作。
	u := user{username: "bob", password: "123", email: "123@sohu.com"}
	_, err = mydb.Exec("insert into users(username,password,email) values(?,?,?)", u.username, u.password, u.email)
	if err != nil {
		fmt.Println("err2=", err)
		return
	}

	//QueryRow方法用于查询一行。
	myRow := mydb.QueryRow("select id,username,password,email from users where id=?", 2)
	var myres user
	myRow.Scan(&myres.id, &myres.username, &myres.password, &myres.email)
	fmt.Println("myres=", myres)



	//Query方法用于查询若干行数据。
	allRow, err4 := mydb.Query("select id,username,password,email from users")
	if err4 != nil {
		fmt.Println("err4=", err4)
		return
	}
	defer allRow.Close() //这一步最好别缺，当然，根据官方文档，当Next()返回值为false时Close()会自动执行
	var alluser user
	for { //这里可以直接 for allRow.Next {} 从而省掉if语句，更简洁，详细见官方文档示例
		if allRow.Next() {
			allRow.Scan(&alluser.id, &alluser.username, &alluser.password, &alluser.email)
			fmt.Println(alluser)
		} else {
			break
		}
	}


	//Prepare方法用于使用MySQL的预编译功能，即先发送不含参数的sql语句，让MySQL预编译，再发送参数配合已完成的预编译完成sql。
	//MySQL中的预编译功能，可以防止sql注入，也能够提高被反复使用的、只是参数不同的sql的执行效率。
	myStmt, err2 := mydb.Prepare("select id,username,password,email from users where id =?")
	if err2 != nil {
		fmt.Println("err2=", err2)
		return
	}
	meRow := myStmt.QueryRow(4)
	var a user
	err3 := meRow.Scan(&a.id, &a.username, &a.password, &a.email)
	if err3 != nil {
		fmt.Println("err3=", err3)
		return
	}
	fmt.Println("a=", a)

}
type user struct {
	id       int
	username string
	password string
	email    string
}