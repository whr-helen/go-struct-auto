package models 

type Mall_test struct { 
    Id int `xorm:"not null pk autoincr comment('ID') INT(11)" json:"id"` 
    Admin_id int `xorm:"comment('管理员ID')" json:"admin_id"` 
    Category_id int `xorm:"comment('分类ID(单选)')" json:"category_id"` 
    Category_ids string `xorm:"comment('分类ID(多选)')" json:"category_ids"` 
    Week string `xorm:"comment('星期(单选):monday=星期一,tuesday=星期二,wednesday=星期三')" json:"week"` 
    Flag string `xorm:"comment('标志(多选):hot=热门,index=首页,recommend=推荐')" json:"flag"` 
    Genderdata string `xorm:"comment('性别(单选):male=男,female=女')" json:"genderdata"` 
    Hobbydata string `xorm:"comment('爱好(多选):music=音乐,reading=读书,swimming=游泳')" json:"hobbydata"` 
    Title string `xorm:"comment('标题')" json:"title"` 
    Content string `xorm:"comment('内容')" json:"content"` 
    Image string `xorm:"comment('图片')" json:"image"` 
    Images string `xorm:"comment('图片组')" json:"images"` 
    Attachfile string `xorm:"comment('附件')" json:"attachfile"` 
    Keywords string `xorm:"comment('关键字')" json:"keywords"` 
    Description string `xorm:"comment('描述')" json:"description"` 
    City string `xorm:"comment('省市')" json:"city"` 
    Price string `xorm:"comment('价格')" json:"price"` 
    Views int `xorm:"comment('点击')" json:"views"` 
    Startdate string `xorm:"comment('开始日期')" json:"startdate"` 
    Activitytime string `xorm:"comment('活动时间(datetime)')" json:"activitytime"` 
    Year string `xorm:"comment('年')" json:"year"` 
    Times string `xorm:"comment('时间')" json:"times"` 
    Refreshtime int `xorm:"comment('刷新时间(int)')" json:"refreshtime"` 
    Createtime int `xorm:"comment('创建时间')" json:"createtime"` 
    Updatetime int `xorm:"comment('更新时间')" json:"updatetime"` 
    Weigh int `xorm:"comment('权重')" json:"weigh"` 
    Switch int `xorm:"comment('开关')" json:"switch"` 
    Status string `xorm:"comment('状态')" json:"status"` 
    State string `xorm:"comment('状态值:0=禁用,1=正常,2=推荐')" json:"state"` 
}