package models 

type Mall_auth_rule struct { 
    Id int `xorm:"not null pk autoincr comment('') INT(11)" json:"id"` 
    Type string `xorm:"comment('menu为菜单,file为权限节点')" json:"type"` 
    Pid int `xorm:"comment('父ID')" json:"pid"` 
    Name string `xorm:"comment('规则名称')" json:"name"` 
    Title string `xorm:"comment('规则名称')" json:"title"` 
    Icon string `xorm:"comment('图标')" json:"icon"` 
    Condition string `xorm:"comment('条件')" json:"condition"` 
    Remark string `xorm:"comment('备注')" json:"remark"` 
    Ismenu int `xorm:"comment('是否为菜单')" json:"ismenu"` 
    Createtime int `xorm:"comment('创建时间')" json:"createtime"` 
    Updatetime int `xorm:"comment('更新时间')" json:"updatetime"` 
    Weigh int `xorm:"comment('权重')" json:"weigh"` 
    Status string `xorm:"comment('状态')" json:"status"` 
}