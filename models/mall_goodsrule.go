package models 

type Mall_goodsrule struct { 
    Id int `xorm:"not null pk autoincr comment('序号') INT(11)" json:"id"` 
    Goods_id int `xorm:"comment('商品序号')" json:"goods_id"` 
    Count int `xorm:"comment('商品序号')" json:"count"` 
    Price float64 `xorm:"comment('商品价格')" json:"price"` 
    Postage float64 `xorm:"comment('邮费')" json:"postage"` 
    Createtime int `xorm:"comment('创建时间')" json:"createtime"` 
}