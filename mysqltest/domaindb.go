package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

const (
	DB_Driver = "user_rw:yF42BO7iMoExfE@tcp(10.124.51.29:3306)/cibn_domain"
)

type domainInfo struct {
	groupId   string
	domainDes string
}

func QueryFromDB(db *sql.DB, domainmap map[string]string) {
	rows, err := db.Query("SELECT `groupId`,`domainDes` FROM groupDomainIndex where `domainDes` like '%account%'")
	//CheckErr(err)
	if err != nil {
		fmt.Println("error:", err)
	} else {
	}
	i := 1
	for rows.Next() {
		var domainDes string
		var groupId string

		//CheckErr(err)
		err = rows.Scan(&groupId, &domainDes)
		fmt.Printf("%d , %s : %s\n", i, groupId, domainDes)
		domainmap[groupId] = domainDes
		i++
	}
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
		fmt.Println("err:", err)
	}
}

func UpdateDB(db *sql.DB, groupId string, domainDes string) {
	stmt, err := db.Prepare("update groupDomain set domain=? where groupId=? and label='adUrl'")
	CheckErr(err)
	res, err := stmt.Exec(domainDes, groupId)
	affect, err := res.RowsAffected()
	fmt.Println("更新数据：", affect)
	CheckErr(err)
}

func OpenDB() (success bool, db *sql.DB) {
	var isOpen bool
	fmt.Println(DB_Driver)
	db, err := sql.Open("mysql", "root:vFAcP9=$xX,Vz0RnjybnX^C@cxdkUmp0@tcp(10.112.40.80:33107)/domain")
	if err != nil {
		isOpen = false
		fmt.Println("open db is false")
	} else {
		isOpen = true
		fmt.Println("open db is true")
	}
	return isOpen, db
}

func main() {
	domainMap := make(map[string]string)
	_, db := OpenDB()

	QueryFromDB(db, domainMap)
	//i := 1
	//for groupid, domainDes := range domainMap {
	//	//if strings.Compare("5b18ad99c96a885f0098bd89", groupid) ==0 {
	//		fmt.Println("===============================")
	//		fmt.Printf("%d , %s : %s\n", i, groupid, domainDes)
	//		fmt.Println("===============================")
	//		domainDes = strings.Replace(domainDes, "69600", "75000", -1)
	//
	//		fmt.Println("+===============================")
	//		fmt.Printf("%d , %s : %s\n", i, groupid, domainDes)
	//		fmt.Println("+===============================")
	//
	//		UpdateDB(db, groupid, domainDes)
	//	//}
	//	i++
	//}
}
