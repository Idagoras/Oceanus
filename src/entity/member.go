package entity

type Member struct {
	Name string `form:"name" json:"name" binding:"required,NameValid"`
	Age  int    `form:"age" json:"age" binding:"required,gt=10,lt=100"`
}
