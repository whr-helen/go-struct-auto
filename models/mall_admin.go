package models 

type Mall_admin struct { 
    Id int `xorm:"not null pk autoincr comment('ID') INT(11)" json:"id"` 
    Username string `xorm:"comment('用户名')" json:"username"` 
    Nickname string `xorm:"comment('昵称')" json:"nickname"` 
    Password string `xorm:"comment('密码')" json:"password"` 
    Salt string `xorm:"comment('密码盐')" json:"salt"` 
    Avatar string `xorm:"comment('头像')" json:"avatar"` 
    Email string `xorm:"comment('电子邮箱')" json:"email"` 
    Loginfailure int `xorm:"comment('失败次数')" json:"loginfailure"` 
    Logintime int `xorm:"comment('登录时间')" json:"logintime"` 
    Createtime int `xorm:"comment('创建时间')" json:"createtime"` 
    Updatetime int `xorm:"comment('更新时间')" json:"updatetime"` 
    Token string `xorm:"comment('Session标识')" json:"token"` 
    Status string `xorm:"comment('状态')" json:"status"` 
}