package main

import (
	"database/sql"
	"fmt"
	//"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:456789@tcp(127.0.0.1:3306)/test?charset=utf8")
	checkErr(err)
	// fmt.Println(db, err)
	stmt, err := db.Prepare("INSERT INTO userinfo SET username=?,department=?,created=?")
	res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
	fmt.Println(res)
	id, err := res.LastInsertId()
	fmt.Println(id)

	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)

	res, err = stmt.Exec("astaxieupdate", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)

	//查询数据
	rows, err := db.Query("SELECT * FROM userinfo")
	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	//删除数据
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)

	db.Close()

}
func checkErr(err error) {
	if err != nil {
		panic(err)
	}

}
