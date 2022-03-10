# discovery
个人对go框架的一次梳理

由于是思考+踩坑总结,所以每个文件夹内都有一个README.md,包含了对这个模块内的一些考虑

## makefile
使用make 命令进行一些简单的项目管理,比如整理代码,代码检查

思考点:
1. 打包是否使用make
2. 安装依赖是否使用make
3. swagger 是否使用make来启动,顺便生成json文件,还是使用proto文件来生成swagger文档

make中,fmt命令,不要格式化一些自动生成的文件,如kratos生成的
