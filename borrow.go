package main

import (
	"reflect"
	"./model"
)

/*entity*/
type Borrow struct {
	BorrowName             string    `column:"borrow_name"`
}


/*dao*/
type IBorrowDao interface {
	model.IBaseDao
}

type borrowDao struct {
	model.BaseDao
}
var borrowDaoImpl IBorrowDao

func BorrowDao() IBorrowDao {
	if borrowDaoImpl == nil {
		borrowDaoImpl = &borrowDao{model.BaseDao{EntityType: reflect.TypeOf(new(Borrow)).Elem()}}
		borrowDaoImpl.Init()
	}

	return borrowDaoImpl
}

func main(){
	borrow := new(Borrow)
	borrow.BorrowName = "sdsd"
	BorrowDao().Save(borrow)

}