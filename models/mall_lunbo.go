package models 

type Mall_lunbo struct { 
    Id int `xorm:"not null pk autoincr comment('序号') INT(11)" json:"id"` 
    Name string `xorm:"comment('名称')" json:"name"` 
    Image string `xorm:"comment('图片')" json:"image"` 
    Goods_id int `xorm:"comment('商品序号')" json:"goods_id"` 
    Createtime int `xorm:"comment('创建时间')" json:"createtime"` 
}