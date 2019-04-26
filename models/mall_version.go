package models 

type Mall_version struct { 
    Id int `xorm:"not null pk autoincr comment('ID') INT(11)" json:"id"` 
    Oldversion string `xorm:"comment('旧版本号')" json:"oldversion"` 
    Newversion string `xorm:"comment('新版本号')" json:"newversion"` 
    Packagesize string `xorm:"comment('包大小')" json:"packagesize"` 
    Content string `xorm:"comment('升级内容')" json:"content"` 
    Downloadurl string `xorm:"comment('下载地址')" json:"downloadurl"` 
    Enforce int `xorm:"comment('强制更新')" json:"enforce"` 
    Createtime int `xorm:"comment('创建时间')" json:"createtime"` 
    Updatetime int `xorm:"comment('更新时间')" json:"updatetime"` 
    Weigh int `xorm:"comment('权重')" json:"weigh"` 
    Status string `xorm:"comment('状态')" json:"status"` 
}