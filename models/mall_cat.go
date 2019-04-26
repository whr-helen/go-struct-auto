package models 

type Mall_cat struct { 
    Id int `xorm:"not null pk autoincr comment('序号') INT(11)" json:"id"` 
    Goods_id int `xorm:"comment('商品序号')" json:"goods_id"` 
    Count int `xorm:"comment('商品数量')" json:"count"` 
    Account_id int `xorm:"comment('用户序号')" json:"account_id"` 
    Createtime int `xorm:"comment('创建时间')" json:"createtime"` 
}