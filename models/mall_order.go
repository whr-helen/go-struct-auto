package models 

type Mall_order struct { 
    Id int `xorm:"not null pk autoincr comment('序号') INT(11)" json:"id"` 
    Code string `xorm:"comment('订单号')" json:"code"` 
    Account_id int `xorm:"comment('用户序号')" json:"account_id"` 
    Rulejson string `xorm:"comment('商品规则')" json:"rulejson"` 
    Datajson string `xorm:"comment('商品信息')" json:"datajson"` 
    Addressjson string `xorm:"comment('地址信息')" json:"addressjson"` 
    Count int `xorm:"comment('商品数量')" json:"count"` 
    Price float64 `xorm:"comment('总价')" json:"price"` 
    Status string `xorm:"comment('状态:0='未付款',1='已付款',2='已发货',3='已完成',-1='订单超时'')" json:"status"` 
    Couriernumber string `xorm:"comment('快递单号')" json:"couriernumber"` 
    Createtime int `xorm:"comment('创建时间')" json:"createtime"` 
}