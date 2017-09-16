package apis

import (
	"github.com/gin-gonic/gin"
	"log"
	"fmt"
	"net/http"
	."../models"
)

func IndexApi(c *gin.Context) {
	c.String(http.StatusOK, "It works")
}

func AddBorrowApi(c *gin.Context) {


	firstName := c.Request.FormValue("borrow_name")


	p := Borrow{BorrowName: firstName}

	ra, err := p.AddBorrow()
	if err != nil {
		log.Fatalln(err)
	}
	msg := fmt.Sprintf("insert successful %d", ra)
	c.JSON(http.StatusOK, gin.H{
		"msg": msg,
	})
}
