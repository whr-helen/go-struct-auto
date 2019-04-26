package models 

type Mall_auth_group struct { 
    Id int `xorm:"not null pk autoincr comment('') INT(11)" json:"id"` 
    Pid int `xorm:"comment('父组别')" json:"pid"` 
    Name string `xorm:"comment('组名')" json:"name"` 
    Rules string `xorm:"comment('规则ID')" json:"rules"` 
    Createtime int `xorm:"comment('创建时间')" json:"createtime"` 
    Updatetime int `xorm:"comment('更新时间')" json:"updatetime"` 
    Status string `xorm:"comment('状态')" json:"status"` 
}