package models 

type Mall_attachment struct { 
    Id int `xorm:"not null pk autoincr comment('ID') INT(11)" json:"id"` 
    Admin_id int `xorm:"comment('管理员ID')" json:"admin_id"` 
    User_id int `xorm:"comment('会员ID')" json:"user_id"` 
    Url string `xorm:"comment('物理路径')" json:"url"` 
    Imagewidth string `xorm:"comment('宽度')" json:"imagewidth"` 
    Imageheight string `xorm:"comment('高度')" json:"imageheight"` 
    Imagetype string `xorm:"comment('图片类型')" json:"imagetype"` 
    Imageframes int `xorm:"comment('图片帧数')" json:"imageframes"` 
    Filesize int `xorm:"comment('文件大小')" json:"filesize"` 
    Mimetype string `xorm:"comment('mime类型')" json:"mimetype"` 
    Extparam string `xorm:"comment('透传数据')" json:"extparam"` 
    Createtime int `xorm:"comment('创建日期')" json:"createtime"` 
    Updatetime int `xorm:"comment('更新时间')" json:"updatetime"` 
    Uploadtime int `xorm:"comment('上传时间')" json:"uploadtime"` 
    Storage string `xorm:"comment('存储位置')" json:"storage"` 
    Sha1 string `xorm:"comment('文件 sha1编码')" json:"sha1"` 
}