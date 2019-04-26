package models 

type Mall_goodsview struct { 
    Id int `xorm:"comment('序号')" json:"id"` 
    Old_price float64 `xorm:"comment('原价')" json:"old_price"` 
    Guige string `xorm:"comment('规格')" json:"guige"` 
    Title string `xorm:"comment('简介')" json:"title"` 
    Name string `xorm:"comment('名称')" json:"name"` 
    Count int `xorm:"comment('商品序号')" json:"count"` 
    Price float64 `xorm:"comment('商品价格')" json:"price"` 
    Images string `xorm:"comment('图片')" json:"images"` 
    Datascontent string `xorm:"comment('介绍')" json:"datascontent"` 
    Createtime int `xorm:"comment('创建时间')" json:"createtime"` 
}