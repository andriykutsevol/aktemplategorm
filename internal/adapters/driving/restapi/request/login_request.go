package request

type Login struct {
	UserName string `json:"UserName" binding:"required"`
	Password string `json:"Password" binding:"required"`
}
