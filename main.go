package main

import (
	//	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
	"./model"
)

var sqlDB *sql.DB

func main() {

	sqlDB = model.Open();
	//model.Open();

	//db,err := sql.Open("mysql", "root:@tcp(10.211.55.9:3306)/dqcenter?charset=utf8")
	//prerr(err)
	id := test(sqlDB)
	//db := opendb("root:@tcp(10.20.70.215:3306)/dqcenter?charset=utf8")
	//
	//ids := insert(sqlDB)
	ids := insert(sqlDB)
	//query(db)
	//update(db,id)
	fmt.Println(id)
	fmt.Println(ids)
}
func test(db *sql.DB) int64 {
	fmt.Println(1)

	return 1
}

//打开数据库连接
func opendb(dbstr string) (*sql.DB) {
	//dsn: [username[:password]@][protocol[(address)]]/dbname[?param1=value1&paramN=valueN]

	db, err := sql.Open("mysql", dbstr)
	prerr(err)
	return db
}

//插入数据
func insert(db *sql.DB) int64 {

	stmt, err := db.Prepare("INSERT INTO dq_borrow SET borrow_name=?")
	prerr(err)

	res, err := stmt.Exec("go-2017-9-16")
	prerr(err)

	id, err := res.LastInsertId()
	prerr(err)

	fmt.Println(id)
	return id

}

//更新数据
func update(db *sql.DB, id int64) {
	stmt, err := db.Prepare("update dq_borrow set name=? where borrow_id=?")
	prerr(err)

	res, err := stmt.Exec("abloz2", id)
	prerr(err)

	affect, err := res.RowsAffected()
	prerr(err)

	fmt.Println(affect)
}

//查询数据
func query(db *sql.DB) {

	rows, err := db.Query("SELECT * FROM dq_borrow")
	prerr(err)

	for rows.Next() {
		var id int
		var name string
		var department string
		var created string
		err = rows.Scan(&id, &name, &department, &created)
		prerr(err)
		fmt.Println(id)
		fmt.Println(name)
		fmt.Println(department)
		fmt.Println(created)
	}
}

//删除数据
func del(db *sql.DB, id int64) {
	stmt, err := db.Prepare("delete from dq_borrow where borrow_id=?")
	prerr(err)

	res, err := stmt.Exec(id)
	prerr(err)

	affect, err := res.RowsAffected()
	prerr(err)

	fmt.Println(affect)
}
func prerr(err error) {
	if err != nil {
		panic(err)
	}
}
