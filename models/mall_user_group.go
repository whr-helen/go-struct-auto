package models 

type Mall_user_group struct { 
    Id int `xorm:"not null pk autoincr comment('') INT(11)" json:"id"` 
    Name string `xorm:"comment('组名')" json:"name"` 
    Rules string `xorm:"comment('权限节点')" json:"rules"` 
    Createtime int `xorm:"comment('添加时间')" json:"createtime"` 
    Updatetime int `xorm:"comment('更新时间')" json:"updatetime"` 
    Status string `xorm:"comment('状态')" json:"status"` 
}