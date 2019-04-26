package models 

type Mall_smokeview struct { 
    Account_id int `xorm:"comment('用户序号')" json:"account_id"` 
    Number int `xorm:"comment('次数')" json:"number"` 
    Order_id int `xorm:"comment('订单序号')" json:"order_id"` 
    Id int `xorm:"comment('序号')" json:"id"` 
    Name string `xorm:"comment('名称')" json:"name"` 
    Id_0 int `xorm:"comment('序号')" json:"id_0"` 
    Headimgurl string `xorm:"comment('头像')" json:"headimgurl"` 
    Price float64 `xorm:"comment('金额')" json:"price"` 
    Createtime int `xorm:"comment('创建时间')" json:"createtime"` 
    Type string `xorm:"comment('状态:0='未付款',1='已付款',2='已发货',-1='订单超时'')" json:"type"` 
    Id_1 int `xorm:"comment('序号')" json:"id_1"` 
    Account_id_0 int `xorm:"comment('用户序号')" json:"account_id_0"` 
    Code string `xorm:"comment('订单号')" json:"code"` 
}