# rename

rename 是一个简单而强大的命令行工具，用于重命名文件和目录。它提供了多种重命名选项，包括更改文件扩展名、转换文件名格式以及简单的重命名操作。

## 功能

- 更改文件扩展名
- 转换文件名格式（大写、小写、驼峰式、蛇形式、烤串式等）
- 简单重命名

## 安装

确保您已安装Go语言环境，然后运行以下命令：
```bash
go get github.com/S2O3/rename
```


## 使用方法

基本语法：
```bash
rename [from] [to] [options...]
```

### 示例

1. 更改文件名：
rename a.txt b //更改a.txt 为 b.txt
2. 更改文件扩展名：
rename a.txt .mp3 //更改a.txt 为 a.mp3
3. 更改文件名格式
rename a.txt :upper //大写  
             :lower //小写  
             :camel //驼峰  
             :snake //蛇形  
             :kebab //烤串式  
             :plus [append_context]// 增添
## 贡献

欢迎提交问题和拉取请求。对于重大更改，请先开启一个问题讨论您想要更改的内容。









   
