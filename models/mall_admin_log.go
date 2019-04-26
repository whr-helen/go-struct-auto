package models 

type Mall_admin_log struct { 
    Id int `xorm:"not null pk autoincr comment('ID') INT(11)" json:"id"` 
    Admin_id int `xorm:"comment('管理员ID')" json:"admin_id"` 
    Username string `xorm:"comment('管理员名字')" json:"username"` 
    Url string `xorm:"comment('操作页面')" json:"url"` 
    Title string `xorm:"comment('日志标题')" json:"title"` 
    Content string `xorm:"comment('内容')" json:"content"` 
    Ip string `xorm:"comment('IP')" json:"ip"` 
    Useragent string `xorm:"comment('User-Agent')" json:"useragent"` 
    Createtime int `xorm:"comment('操作时间')" json:"createtime"` 
}