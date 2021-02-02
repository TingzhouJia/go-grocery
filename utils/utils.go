package utils

import (
	"crypto/md5"
	"fmt"
	"grocery/dto"
	"grocery/enum"
	"io"
)

func Page(Limit,Page int) (limit,offset int){
	if Limit>0 {
		limit=Limit
	}else{
		limit=10
	}
	if Page>0 {
		offset=(Page-1)*limit
	}else {
		offset=-1
	}
	return limit,offset
}

func Sort(Sort string) (sort string){
	if Sort!=""{
		sort=Sort
	}else {
		sort = "created_at desc"
	}
	return sort
}

func Md5(str string) string {
	w := md5.New()
	io.WriteString(w, str)
	md5str := fmt.Sprintf("%x", w.Sum(nil))
	return md5str
}

func BaseReturnBody() dto.Entity {
	return dto.Entity{
		Code: int(enum.ServerError),
		Msg: enum.ServerError.String(),
		Data: nil,
		Total: 0,
		TotalPage: 0,
	}
}

func ReturnBody(Total,Page int,status enum.ResStatus,data interface{},) dto.Entity {
	return dto.Entity{
		Code: int(status),
		Msg: status.String(),
		Total: Total,
		TotalPage: Page,
		Data: data,
	}
}
