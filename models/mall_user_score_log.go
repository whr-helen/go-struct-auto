package models 

type Mall_user_score_log struct { 
    Id int `xorm:"not null pk autoincr comment('') INT(11)" json:"id"` 
    User_id int `xorm:"comment('会员ID')" json:"user_id"` 
    Score int `xorm:"comment('变更积分')" json:"score"` 
    Before int `xorm:"comment('变更前积分')" json:"before"` 
    After int `xorm:"comment('变更后积分')" json:"after"` 
    Memo string `xorm:"comment('备注')" json:"memo"` 
    Createtime int `xorm:"comment('创建时间')" json:"createtime"` 
}