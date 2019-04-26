package models 

type Mall_address struct { 
    Id int `xorm:"not null pk autoincr comment('序号') INT(11)" json:"id"` 
    Account_id int `xorm:"comment('用户序号')" json:"account_id"` 
    Name string `xorm:"comment('名称')" json:"name"` 
    Mobile string `xorm:"comment('手机号')" json:"mobile"` 
    Address string `xorm:"comment('省市区')" json:"address"` 
    Details string `xorm:"comment('详细地址')" json:"details"` 
    Createtime int `xorm:"comment('创建时间')" json:"createtime"` 
}