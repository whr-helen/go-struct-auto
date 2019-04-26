package models 

type Mall_user_money_log struct { 
    Id int `xorm:"not null pk autoincr comment('') INT(11)" json:"id"` 
    User_id int `xorm:"comment('会员ID')" json:"user_id"` 
    Money float64 `xorm:"comment('变更余额')" json:"money"` 
    Before float64 `xorm:"comment('变更前余额')" json:"before"` 
    After float64 `xorm:"comment('变更后余额')" json:"after"` 
    Memo string `xorm:"comment('备注')" json:"memo"` 
    Createtime int `xorm:"comment('创建时间')" json:"createtime"` 
}