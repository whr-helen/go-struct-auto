package models 

type Mall_article struct { 
    Id int `xorm:"not null pk autoincr comment('序号') INT(11)" json:"id"` 
    Name string `xorm:"comment('标题')" json:"name"` 
    Content string `xorm:"comment('内容')" json:"content"` 
}