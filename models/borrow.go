package models

import (
	"../model"
)


type Borrow struct {
	BorrowName             string    `json:"borrow_name" form:"borrow_name"`
}

func (p *Borrow) AddBorrow() (id int64, err error) {

	rs, err := model.Open().Exec("INSERT INTO dq_borrow(borrow_name ) VALUES (?)", p.BorrowName)
	if err != nil {
		return
	}
	id, err = rs.LastInsertId()
	return
}