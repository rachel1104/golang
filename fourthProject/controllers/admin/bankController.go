package admin

import (
	"fourthProject/models"
	"github.com/gin-gonic/gin"
)

type BankController struct {
	BaseController
}

func (con BankController) Index(c *gin.Context) {
	//開始事務
	tx := models.DB.Begin()

	//拋出異常
	defer func() {
		if r := recover(); r != nil {
			//遇到錯誤時回滾事務
			tx.Rollback()
			con.error(c)
		}
	}()
	//張三給李四轉賬
	u1 := models.Bank{Id: 1}
	tx.Find(&u1)
	u1.Balance = u1.Balance - 100
	if err := tx.Save(&u1).Error; err != nil {
		tx.Rollback()
		con.error(c)
		return
	} else if u1.Balance < 0 {
		tx.Rollback()
		con.error(c)
		return
	}
	//在李四的賬戶裡面增加100元
	u2 := models.Bank{Id: 2}
	tx.Find(&u2)
	u2.Balance = u2.Balance + 100
	if err := tx.Save(&u2).Error; err != nil {
		tx.Rollback()
		con.error(c)
		return
	}
	tx.Commit()
	//轉賬成功
	con.success(c)
}
