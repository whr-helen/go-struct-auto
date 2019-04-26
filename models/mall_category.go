package models 

type Mall_category struct { 
    Id int `xorm:"not null pk autoincr comment('') INT(11)" json:"id"` 
    Pid int `xorm:"comment('父ID')" json:"pid"` 
    Type string `xorm:"comment('栏目类型')" json:"type"` 
    Name string `xorm:"comment('')" json:"name"` 
    Nickname string `xorm:"comment('')" json:"nickname"` 
    Flag string `xorm:"comment('')" json:"flag"` 
    Image string `xorm:"comment('图片')" json:"image"` 
    Keywords string `xorm:"comment('关键字')" json:"keywords"` 
    Description string `xorm:"comment('描述')" json:"description"` 
    Diyname string `xorm:"comment('自定义名称')" json:"diyname"` 
    Createtime int `xorm:"comment('创建时间')" json:"createtime"` 
    Updatetime int `xorm:"comment('更新时间')" json:"updatetime"` 
    Weigh int `xorm:"comment('权重')" json:"weigh"` 
    Status string `xorm:"comment('状态')" json:"status"` 
}