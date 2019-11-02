package vo

type LoginReq struct {
}

type LoginResp struct {
	Token string `json:"token"`
}

func Login() *LoginResp {

	resp := &LoginResp{Token:"hello_login"}
	return resp
}
