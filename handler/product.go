package handler

import (
	"github.com/gin-gonic/gin"
	"grocery/dto"
	"grocery/enum"
	"grocery/model"
	"grocery/query"
	"grocery/service"
	"grocery/utils"
	"net/http"
)

type ProductHandler struct {
	ProductSrv service.ProductSrv
}

func (h *ProductHandler) GetEntity(result model.Product) dto.Product {
	return dto.Product{
		Id:                   result.ProductId,
		Key:                  result.ProductId,
		ProductId:            result.ProductId,
		ProductName:          result.ProductName,
		ProductIntro:         result.ProductIntro,
		CategoryId:           result.CategoryId,
		ProductCoverImg:      result.ProductCoverImg,
		ProductBanner:        result.ProductBanner,
		OriginalPrice:        result.OriginalPrice,
		SellingPrice:         result.SellingPrice,
		StockNum:             result.StockNum,
		Tag:                  result.Tag,
		SellStatus:           result.SellStatus,
		ProductDetailContent: result.ProductDetailContent,
		IsDeleted:            result.IsDeleted,
	}
}

func (h *ProductHandler) ProductInfoHandler(c *gin.Context) {
	entity := utils.BaseReturnBody()
	productId := c.Param("id")
	if productId == "" {
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	u := model.Product{
		ProductId: productId,
	}
	result, err := h.ProductSrv.Get(u)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}

	r := h.GetEntity(*result)

	entity = utils.ReturnBody(0,0,enum.Success,r)
	c.JSON(http.StatusOK, gin.H{"entity": entity})
}

func (h *ProductHandler) ProductListHandler(c *gin.Context) {
	var q query.ListQuery
	entity := utils.BaseReturnBody()
	err := c.ShouldBindQuery(&q)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"entity": entity})
		return
	}
	list, err := h.ProductSrv.List(&q)
	total, err := h.ProductSrv.GetTotal(&q)

	if err != nil {
		panic(err)
	}
	if q.PageSize == 0 {
		q.PageSize = 5
	}
	ret := int(total % q.PageSize)
	ret2 := int(total / q.PageSize)
	totalPage := 0
	if ret == 0 {
		totalPage = ret2
	} else {
		totalPage = ret2 + 1
	}
	var newList []*dto.Product
	for _, item := range list {
		r := h.GetEntity(*item)
		newList = append(newList, &r)
	}

	entity = utils.ReturnBody(total,totalPage,enum.Success,newList)
	c.JSON(http.StatusOK, gin.H{"entity": entity})
}

func (h *ProductHandler) AddProductHandler(c *gin.Context) {
	entity := utils.BaseReturnBody()
	p := model.Product{}
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		return
	}

	r, err := h.ProductSrv.Add(p)
	if err != nil {
		entity.Msg = err.Error()
		return
	}
	if r.ProductId == "" {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		return
	}
	entity.Code = int(enum.Success)
	entity.Msg = enum.Success.String()
	c.JSON(http.StatusOK, gin.H{"entity": entity})

}

func (h *ProductHandler) EditProductHandler(c *gin.Context) {
	p := model.Product{}
	entity := utils.BaseReturnBody()
	err := c.ShouldBindJSON(&p)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		return
	}
	b, err := h.ProductSrv.Edit(p)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		return
	}
	if b {
		entity.Code = int(enum.Accepted)
		entity.Msg = enum.Accepted.String()
		c.JSON(http.StatusOK, gin.H{"entity": entity})
	}

}

func (h *ProductHandler) DeleteProductHandler(c *gin.Context) {
	id := c.Param("id")
	b, err := h.ProductSrv.Delete(id)
	entity :=utils.BaseReturnBody()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"entity": entity})
		return
	}
	if b {
		entity.Code = int(enum.DeleteSuccess)
		entity.Msg = enum.DeleteSuccess.String()
		c.JSON(http.StatusOK, gin.H{"entity": entity})
	}
}
