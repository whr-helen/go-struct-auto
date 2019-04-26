package models 

type Mall_bank struct { 
    Id int `xorm:"not null pk autoincr comment('序号') INT(11)" json:"id"` 
    Bankname string `xorm:"comment('银行')" json:"bankname"` 
    Banksubname string `xorm:"comment('支行')" json:"banksubname"` 
    Code string `xorm:"comment('卡号')" json:"code"` 
    Name string `xorm:"comment('姓名')" json:"name"` 
    Account_id int `xorm:"comment('会员序号')" json:"account_id"` 
    Createtime int `xorm:"comment('创建时间')" json:"createtime"` 
    Type string `xorm:"comment('类型:0='微信',1='支付宝',2='银行'')" json:"type"` 
}