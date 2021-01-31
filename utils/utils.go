package utils

import (
	"crypto/md5"
	"fmt"
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
