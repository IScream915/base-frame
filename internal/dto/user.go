package dto

type CreateUserReq struct {
	Account  string `json:"account" binding:"required"`
	NickName string `json:"nickname" binding:"required"`
	Password string `json:"password" binding:"required"`
	Age      uint8  `json:"age" binding:"required"`
	Sex      string `json:"sex" binding:"required"`
}

type UpdateUserReq struct {
	NickName string `json:"nickname"`
	Password string `json:"password"`
	Age      uint8  `json:"age"`
	Sex      string `json:"sex"`
}

type DeleteUserReq struct {
	IDs []uint64 `json:"ids" binding:"required"`
}
