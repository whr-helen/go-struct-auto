package models 

type Mall_user struct { 
    Id int `xorm:"not null pk autoincr comment('ID') INT(11)" json:"id"` 
    Group_id int `xorm:"comment('组别ID')" json:"group_id"` 
    Username string `xorm:"comment('用户名')" json:"username"` 
    Nickname string `xorm:"comment('昵称')" json:"nickname"` 
    Password string `xorm:"comment('密码')" json:"password"` 
    Salt string `xorm:"comment('密码盐')" json:"salt"` 
    Email string `xorm:"comment('电子邮箱')" json:"email"` 
    Mobile string `xorm:"comment('手机号')" json:"mobile"` 
    Avatar string `xorm:"comment('头像')" json:"avatar"` 
    Level int `xorm:"comment('等级')" json:"level"` 
    Gender int `xorm:"comment('性别')" json:"gender"` 
    Birthday string `xorm:"comment('生日')" json:"birthday"` 
    Bio string `xorm:"comment('格言')" json:"bio"` 
    Money float64 `xorm:"comment('余额')" json:"money"` 
    Score int `xorm:"comment('积分')" json:"score"` 
    Successions int `xorm:"comment('连续登录天数')" json:"successions"` 
    Maxsuccessions int `xorm:"comment('最大连续登录天数')" json:"maxsuccessions"` 
    Prevtime int `xorm:"comment('上次登录时间')" json:"prevtime"` 
    Logintime int `xorm:"comment('登录时间')" json:"logintime"` 
    Loginip string `xorm:"comment('登录IP')" json:"loginip"` 
    Loginfailure int `xorm:"comment('失败次数')" json:"loginfailure"` 
    Joinip string `xorm:"comment('加入IP')" json:"joinip"` 
    Jointime int `xorm:"comment('加入时间')" json:"jointime"` 
    Createtime int `xorm:"comment('创建时间')" json:"createtime"` 
    Updatetime int `xorm:"comment('更新时间')" json:"updatetime"` 
    Token string `xorm:"comment('Token')" json:"token"` 
    Status string `xorm:"comment('状态')" json:"status"` 
    Verification string `xorm:"comment('验证')" json:"verification"` 
}