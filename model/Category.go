package model

type Category struct {
	CategoryId string `json:"categoryId" gorm:"column:category_id"`
	Name string `json:"name" gorm:"column:name"`
	Desc string `json:"desc" gorm:"column:desc"`
	Order string `json:"order" gorm:"column:order"`
	ParentId string `json:"parentId" gorm:"column:parent_id"`
	IsDeleted bool `json:"isDeleted" gorm:"column:is_deleted"`

}
