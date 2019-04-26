package models 

type Mall_user_rule struct { 
    Id int `xorm:"not null pk autoincr comment('') INT(11)" json:"id"` 
    Pid int `xorm:"comment('父ID')" json:"pid"` 
    Name string `xorm:"comment('名称')" json:"name"` 
    Title string `xorm:"comment('标题')" json:"title"` 
    Remark string `xorm:"comment('备注')" json:"remark"` 
    Ismenu int `xorm:"comment('是否菜单')" json:"ismenu"` 
    Createtime int `xorm:"comment('创建时间')" json:"createtime"` 
    Updatetime int `xorm:"comment('更新时间')" json:"updatetime"` 
    Weigh int `xorm:"comment('权重')" json:"weigh"` 
    Status string `xorm:"comment('状态')" json:"status"` 
}