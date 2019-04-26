自动构建工具使用

① 推荐使用方法(支持 linux or mac)
生成数据库所有表的数据结构：
./bin/automatic -d dbname -path ./models  //-d dbname改为自己数据库的名称  -path ./models 是默认值可以改动

② 修改生成工具代码(支持 linux or mac or windows)
如果生成出来的结构不是我们所需要的可以修改 automatic.go 文件
go run automatic.go -d dbname -path ./models

