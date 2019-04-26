package models 

type Mall_ems struct { 
    Id int `xorm:"not null pk autoincr comment('ID') INT(11)" json:"id"` 
    Event string `xorm:"comment('事件')" json:"event"` 
    Email string `xorm:"comment('邮箱')" json:"email"` 
    Code string `xorm:"comment('验证码')" json:"code"` 
    Times int `xorm:"comment('验证次数')" json:"times"` 
    Ip string `xorm:"comment('IP')" json:"ip"` 
    Createtime int `xorm:"comment('创建时间')" json:"createtime"` 
}