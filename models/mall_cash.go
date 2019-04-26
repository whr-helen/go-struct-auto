package models 

type Mall_cash struct { 
    Id int `xorm:"not null pk autoincr comment('序号') INT(11)" json:"id"` 
    Code string `xorm:"comment('编号')" json:"code"` 
    Bank_id int `xorm:"comment('银行序号')" json:"bank_id"` 
    Account_id int `xorm:"comment('会员序号')" json:"account_id"` 
    Price float64 `xorm:"comment('金额')" json:"price"` 
    Status string `xorm:"comment('状态')" json:"status"` 
    Content string `xorm:"comment('回复')" json:"content"` 
    Createtime int `xorm:"comment('创建时间')" json:"createtime"` 
    Bank_info string `xorm:"comment('卡信息')" json:"bank_info"` 
}