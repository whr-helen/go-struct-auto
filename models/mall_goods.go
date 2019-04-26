package models 

type Mall_goods struct { 
    Id int `xorm:"not null pk autoincr comment('序号') INT(11)" json:"id"` 
    Name string `xorm:"comment('名称')" json:"name"` 
    Title string `xorm:"comment('简介')" json:"title"` 
    Guige string `xorm:"comment('规格')" json:"guige"` 
    Price float64 `xorm:"comment('价格')" json:"price"` 
    Old_price float64 `xorm:"comment('原价')" json:"old_price"` 
    Images string `xorm:"comment('图片')" json:"images"` 
    Datascontent string `xorm:"comment('介绍')" json:"datascontent"` 
    Createtime int `xorm:"comment('创建时间')" json:"createtime"` 
}