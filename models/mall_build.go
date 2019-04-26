package models 

type Mall_build struct { 
    Account_id int `xorm:"comment('用户序号')" json:"account_id"` 
    Parent_id int `xorm:"comment('父级序号')" json:"parent_id"` 
    Createtime int `xorm:"comment('创建时间')" json:"createtime"` 
    Id int `xorm:"not null pk autoincr comment('') INT(11)" json:"id"` 
}