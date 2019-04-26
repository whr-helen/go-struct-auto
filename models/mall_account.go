package models 

type Mall_account struct { 
    Id int `xorm:"not null pk autoincr comment('序号') INT(11)" json:"id"` 
    Name string `xorm:"comment('名称')" json:"name"` 
    Openid string `xorm:"comment('唯一标识')" json:"openid"` 
    Mobile string `xorm:"comment('手机号')" json:"mobile"` 
    Headimgurl string `xorm:"comment('头像')" json:"headimgurl"` 
    Price float64 `xorm:"comment('价格')" json:"price"` 
    Createtime int `xorm:"comment('创建时间')" json:"createtime"` 
    Code string `xorm:"comment('支付密码')" json:"code"` 
}