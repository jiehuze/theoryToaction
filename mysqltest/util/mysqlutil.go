package util

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" //必须加该驱动，否则无法连接数据库
)

func QueryFromDB(db *sql.DB) {
	rows, err := db.Query("SELECT domainDes FROM userinfo")
	CheckErr(err)
	if err != nil {
		fmt.Println("error:", err)
	} else {
	}
	for rows.Next() {
		var domainDes string

		CheckErr(err)
		err = rows.Scan(&domainDes)
		fmt.Println(domainDes)
	}
}

func UpdateDB(db *sql.DB, uid string) {
	stmt, err := db.Prepare("update userinfo set username=? where uid=?")
	CheckErr(err)
	res, err := stmt.Exec("zhangqi", uid)
	affect, err := res.RowsAffected()
	fmt.Println("更新数据：", affect)
	CheckErr(err)
}

func DeleteFromDB(db *sql.DB, autid int) {
	stmt, err := db.Prepare("delete from userinfo where autid=?")
	CheckErr(err)
	res, err := stmt.Exec(autid)
	affect, err := res.RowsAffected()
	fmt.Println("删除数据：", affect)
}
