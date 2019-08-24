package mysql

import (
	"database/sql"
	"fmt"
	//_"github.com/go-sql-driver/mysql"
	"os"
)

var db *sql.DB

func init(){
	db,_=sql.Open("mysql","root:191513@tcp(127.0.0.1:3306)/slave?charset=utf8")
	db.SetMaxOpenConns(1000)
	err:=db.Ping()
	if err!=nil{
		fmt.Println("fail to connect to db")
		os.Exit(1)
	}
}

func DBConn()*sql.DB{
	return  db
}
//type Record interface{}
//type DataIterator struct {
//	cache Record
//}
//func (tf *DataIterator) parseRows(rows *sql.Rows, cSize int) (error) {
//	columns, err := rows.Columns()
//	if err != nil {
//		return err
//	}
//	size := len(columns)
//	pts := make([]interface{}, size)
//	container := make([]interface{}, size)
//	tf.cache = make([]Record, cSize)
//	cursor := 0
//	for i := range pts {
//		pts[i] = &container[i]
//	}
//	for rows.Next() {
//		err = rows.Scan(pts...)
//		if err != nil {
//			return err
//		}
//		var r Record = make(map[string]interface{}, size)
//		for i, column := range columns {
//			r[column] = container[i]
//		}
//		tf.cache[cursor] = r
//		cursor++
//	}
//	return nil
//}

