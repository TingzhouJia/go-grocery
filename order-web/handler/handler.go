package handler

import (
	"context"
	"encoding/json"
	hystrix_go "github.com/afex/hystrix-go/hystrix"
	"github.com/micro/go-micro/v2/client"
	log "github.com/micro/go-micro/v2/logger"
	"github.com/micro/go-plugins/wrapper/breaker/hystrix/v2"
	auth "grocery/auth/proto"
	inventory "grocery/inventory-srv/proto"
	order "grocery/order-service/proto"
	"grocery/plugins/session"
	"net/http"
	"strconv"
	"time"
)

var (
	serviceClient order.OrdersService
	authClient    auth.Service
	invClient    inventory.InventorySrvService
)

// Error 错误结构体
type Error struct {
	Code   string `json:"code"`
	Detail string `json:"detail"`
}

func Init() {
	hystrix_go.DefaultVolumeThreshold = 1
	hystrix_go.DefaultErrorPercentThreshold = 1
	cl := hystrix.NewClientWrapper()(client.DefaultClient)
	cl.Init(

		client.Retries(3),

		client.Retry(func(ctx context.Context, req client.Request, retryCount int, err error) (bool, error) {
			//log.Log(req.Method(), retryCount, " client retry")
			return true, nil
		}),
	)
	serviceClient = order.NewOrdersService("mu.micro.book.srv.orders", cl)
	authClient = auth.NewService("mu.micro.book.srv.auth", cl)
}

func New(w http.ResponseWriter, r *http.Request) {
	// 只接受POST请求
	if r.Method != "POST" {
		log.Error("非法请求")
		http.Error(w, "非法请求", 400)
		return
	}

	r.ParseForm()
	bookId, _ := strconv.ParseInt(r.Form.Get("bookId"), 10, 10)

	// 返回结果
	response := map[string]interface{}{}

	// 调用后台服务
	rsp, err := serviceClient.New(context.TODO(), &order.Request{
		BookId: bookId,
		UserId: session.GetSession(w, r).Values["userId"].(int64),
	})

	// 返回结果
	response["ref"] = time.Now().UnixNano()
	if err != nil {
		response["success"] = false
		response["error"] = Error{
			Detail: err.Error(),
		}
	} else {
		response["success"] = true
		response["orderId"] = rsp.Order.Id
	}

	w.Header().Add("Content-Type", "application/json; charset=utf-8")

	// 返回JSON结构
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}