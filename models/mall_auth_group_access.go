package models 

type Mall_auth_group_access struct { 
    Uid int `xorm:"comment('会员ID')" json:"uid"` 
    Group_id int `xorm:"comment('级别ID')" json:"group_id"` 
}