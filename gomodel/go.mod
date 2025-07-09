//模块名
module gomodel
//sdk
go 1.24
//当前模块依赖的包
require (
	//dependency latest
)
exclude (
	//dependency latest
)
//修改依赖包的路径或者版本
//当依赖包发生过迁移
//原始包无法访问
//使用本地包替换原始包
replace (
	//source latest => target latest
)
//撤回
//当前项目作为其他项目的依赖，如果这个版本出现了问题则撤回该版本5 
retract (
    v1.0.0
)