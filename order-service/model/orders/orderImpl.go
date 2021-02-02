package orders

import (
	"context"
	log "github.com/micro/go-micro/v2/logger"
	"grocery/basic/common"
	"grocery/basic/db"
	inventory "grocery/inventory-srv/proto"
	proto "grocery/order-service/proto"
)

func (s service) New(bookId, userId int64) (orderId int64, err error) {
	// 请求销存
	rsp, err := invClient.Sell(context.TODO(), &inventory.Request{
		BookId: bookId, UserId: userId,
	})
	if err != nil {
		log.Infof("[New] Sell 调用库存服务时失败：%s", err.Error())
		return
	}

	// 获取数据库
	o := db.GetDB()
	insertSQL := `INSERT orders (user_id, book_id, inv_his_id, state) VALUE (?, ?, ?, ?)`

	r, err := o.Exec(insertSQL, userId, bookId, rsp.InvH.Id, common.InventoryHistoryStateNotOut)
	if err != nil {
		log.Infof("[New] 新增订单失败，err：%s", err)
		return
	}
	orderId, _ = r.LastInsertId()
	return
}

func (s service) GetOrder(orderId int64) (order *proto.Order, err error) {
	order = &proto.Order{}

	// 获取数据库
	o := db.GetDB()
	// 查询
	err = o.QueryRow("SELECT id, user_id, book_id, inv_his_id, state FROM orders WHERE id = ?", orderId).Scan(
		&order.Id, &order.UserId, &order.BookId, &order.InvHistoryId, &order.State)
	if err != nil {
		log.Infof("[GetOrder] 查询数据失败，err：%s", err)
		return
	}

	return
}

func (s service) UpdateOrderState(orderId int64, state int) (err error) {
	updateSQL := `UPDATE orders SET state = ? WHERE id = ?`

	// 获取数据库
	o := db.GetDB()
	// 更新
	_, err = o.Exec(updateSQL, state, orderId)
	if err != nil {
		log.Infof("[Confirm] 更新失败，err：%s", err)
		return
	}
	return
}
