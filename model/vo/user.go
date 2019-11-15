package vo

type LoginReq struct {
	Account  string `form:"account"`
	Password string `form:"password"`
	Type     uint   `form:"type"`
}

type LoginResp struct {
	Token string `json:"token"`
}

func Login() *LoginResp {

	resp := &LoginResp{Token: "hello_login"}
	return resp
}
