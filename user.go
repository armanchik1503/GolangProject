package todo

type User struct {
	Id       int    `json:"-"`
	Name     string `json:"name"     binding:"required"`
	Surname  string `json:"surname"  binding:"required"`
	Password string `json:"password" binding:"required"`
}
