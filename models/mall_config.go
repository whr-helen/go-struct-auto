package models 

type Mall_config struct { 
    Id int `xorm:"not null pk autoincr comment('') INT(11)" json:"id"` 
    Name string `xorm:"comment('变量名')" json:"name"` 
    Group string `xorm:"comment('分组')" json:"group"` 
    Title string `xorm:"comment('变量标题')" json:"title"` 
    Tip string `xorm:"comment('变量描述')" json:"tip"` 
    Type string `xorm:"comment('类型:string,text,int,bool,array,datetime,date,file')" json:"type"` 
    Value string `xorm:"comment('变量值')" json:"value"` 
    Content string `xorm:"comment('变量字典数据')" json:"content"` 
    Rule string `xorm:"comment('验证规则')" json:"rule"` 
    Extend string `xorm:"comment('扩展属性')" json:"extend"` 
}