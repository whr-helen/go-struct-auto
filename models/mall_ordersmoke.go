package models 

type Mall_ordersmoke struct { 
    Id int `xorm:"not null pk autoincr comment('序号') INT(11)" json:"id"` 
    Order_id int `xorm:"comment('订单序号')" json:"order_id"` 
    Account_id int `xorm:"comment('用户序号')" json:"account_id"` 
    Number int `xorm:"comment('次数')" json:"number"` 
    Price float64 `xorm:"comment('金额')" json:"price"` 
    Type string `xorm:"comment('状态:0='未付款',1='已付款',2='已发货',-1='订单超时'')" json:"type"` 
    Createtime int `xorm:"comment('创建时间')" json:"createtime"` 
}