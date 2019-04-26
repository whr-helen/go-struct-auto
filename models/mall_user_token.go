package models 

type Mall_user_token struct { 
    Token string `xorm:"comment('Token')" json:"token"` 
    User_id int `xorm:"comment('会员ID')" json:"user_id"` 
    Createtime int `xorm:"comment('创建时间')" json:"createtime"` 
    Expiretime int `xorm:"comment('过期时间')" json:"expiretime"` 
}