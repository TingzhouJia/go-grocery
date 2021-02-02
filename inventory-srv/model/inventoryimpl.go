package model

import (
	"fmt"
	log "github.com/micro/go-micro/v2/logger"
	"grocery/basic/common"
	"grocery/basic/db"
	proto "grocery/inventory-srv/proto"
)

func (s service) Sell(bookId, userId int64) (id int64, err error) {
	tx,err:=db.GetDB().Begin()
	if err !=nil{
		log.Errorf("[Sell] 事务开启失败", err.Error())
		return
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()
	querySQL := `SELECT id, book_id, unit_price, stock, version FROM inventory WHERE book_id = ?`

	inv := &proto.Inv{}

	updateSQL := `UPDATE inventory SET stock = ?, version = ?  WHERE book_id = ? AND version = ?`
	var deductInv func() error
	deductInv= func() (errIn error) {
		errIn = tx.QueryRow(querySQL, bookId).Scan(&inv.Id, &inv.BookId, &inv.UnitPrice, &inv.Stock, &inv.Version)
		if errIn != nil {
			log.Errorf("[Sell] 查询数据失败，err：%s", errIn)
			return errIn
		}

		if inv.Stock < 1 {
			errIn = fmt.Errorf("[Sell] 库存不足")
			log.Error(errIn)
			return errIn
		}

		r, errIn := tx.Exec(updateSQL, inv.Stock-1, inv.Version+1, bookId, inv.Version)
		if errIn != nil {
			log.Errorf("[Sell] 更新库存数据失败，err：%s", errIn)
			return
		}
		if affected, _ := r.RowsAffected(); affected == 0 {
			log.Errorf("[Sell] 更新库存数据失败，版本号%d过期，即将重试", inv.Version)
			// 重试，直到没有库存
			deductInv()
		}
		return
	}
	err = deductInv()
	if err != nil {
		log.Errorf("[Sell] 销存失败，err：%s", err)
		return
	}
	insertSQL := `INSERT inventory_history (book_id, user_id, state) VALUE (?, ?, ?) `
	r, err := tx.Exec(insertSQL, bookId, userId, common.InventoryHistoryStateNotOut)
	if err != nil {
		log.Errorf("[Sell] 新增销存记录失败，err：%s", err)
		return
	}

	// 返回历史记录id，作为流水号使用
	id, _ = r.LastInsertId()

	// 忽略error
	tx.Commit()
	return
}

func (s service) Confirm(id int64, state int) (err error) {
	updateSQL := `UPDATE inventory_history SET state = ? WHERE id = ?;`

	// 获取数据库
	o := db.GetDB()

	// 更新
	_, err = o.Exec(updateSQL, state, id)
	if err != nil {
		log.Errorf("[Confirm] 更新失败，err：%s", err)
		return
	}
	return
}
