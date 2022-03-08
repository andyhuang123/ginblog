package forms

type PasswordLoginForm struct {
	// 密码  binding:"required"为必填字段,长度大于3小于20
	PassWord  string `form:"password" json:"password" binding:"required,min=3,max=20"`
	//用户名
	Username string `form:"name" json:"name" binding:"required"`

	Captcha   string `form:"captcha" json:"captcha" binding:"required,min=5,max=5"` // 验证码

	CaptchaId string `form:"captcha_id" json:"captcha_id" binding:"required"`       // 验证码id

}

type UserListForm struct {

	Page  int `form:"page" json:"page" binding:"required"`

	PageSize int `form:"pagesize" json:"pagesize" binding:"required"`
}
