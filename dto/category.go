package dto

type BaseCategory struct {
	CategoryID string                `json:"categoryID"`
	Name       string                `json:"name"`
	Desc       string                `json:"desc"`
	Order      int                   `json:"order"`
	ParentID   string                `json:"parentId"`
	Children   map[string]*SubCategory `json:"children"`
}

type SubCategory struct {
	CategoryID string                `json:"categoryID"`
	Name       string                `json:"name"`
	Desc       string                `json:"desc"`
	Order      int                   `json:"order"`
	ParentID   string                `json:"parentId"`
	Children   map[string]*ChildCategory `json:"children"`
}

type ChildCategory struct {
	Key        string `json:"key"`
	Id         string `json:"id"`
	CategoryID string `json:"categoryID"`
	Name       string `json:"name"`
	Desc       string `json:"desc"`
	Order      int    `json:"order"`
	ParentID   string `json:"parentId"`
	IsDeleted  bool   `json:"isDeleted"`
}
