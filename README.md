# GoRush

![License](https://img.shields.io/badge/license-Apache--2.0-orange)

![Language](https://img.shields.io/badge/language-go-brightgreen)

## 😀1. 介绍

一个为开发者自由、快速、便捷地构建网站服务的网站框架。

## 💘2. 目的

构建这个项目的本意是提供一种方便、快速的自由构建网站方式，让使用者能够根据自己的需求diy网站，并且力图将门槛降到最低。但是，即使对于一个简单的静态网页来说，基本的html和css技能是必要的。

目前个人对该项目的构想经常变化，很难想象怎样的一种方式能够更好地达成这一目的，而最终成果也十分令人期待:star:

## ✨3. 下载

由于配套脚手架功能尚未完善（目前仅能够初始化项目文件夹和配置文件），推荐以包的形式导入：
```shell
import "github.com/levinion/gorush/render"
```

未来将依靠脚手架的方式构建和运行程序（开发中:sweat:)

```shell
go install github.com/levinion/gorush
```

## 🔥4. 使用
### 4.1 项目结构
```shell
.
├── config.toml //配置文件
├── go.mod
├── go.sum
├── main.go
├── resources   //待解析文件存放在此处
└── templates   //模板文件存放在此处
```

### 4.2 配置文件
配置文件默认为`config.toml`，位于项目主目录下，在程序运行前需要先行建立配置文件。下面是配置文件的一个示例：
```toml
[version]
version = "1.0"

[server]
addr = "localhost:9090"
```
### 4.3 项目建立流程
1. 新建文件夹，在该文件夹下进行先期准备：
```shell
git init        #若不使用git请跳过

go mod init <golang项目名>
```
2. 然后分别建立`main.go`文件、`resources`和`templates`目录。

3. 在`resources`文件夹下创建待解析资源（主要是Markdown文件），在`templates`文件夹下添加HTML模板，具体模板和文件写法可参考`doc/example`；模板插入方式详见go语言template包教程。

4. 仿照以下方式建立`main.go`文件

5. 使用`go run .`运行程序

### 4.4. 包使用方式示例:
```go
package main

import (
	"github.com/levinion/gorush/render"
)

func main(){
	r:=render.NewRenderer()		//新建Renderer对象，配置文件将在这一步初始化
	r.Parse("./resources/")		//解析Markdown文件目录，若在上一步传入目录则可略去

	//下面是页面添加流程示例：
    //使用HTML构造页面
	r.RenderHTML("/","./templates/mainPage.html")
    //批量解析并使用MarkDown文件构造页面
	r.GroupRenderMarkdown("/posts/","./templates/posts.html","./templates/default.html")
    //解析并使用单个MarkDown文件构造页面
	// render.RenderMarkdown("/test/","./resources/parts.md","./templates/default.html")

	r.Run()		//调用Run以运行服务，可写入端口，否则将使用配置文件中定义的端口
}

```

### 4.5 命令行（开发中）
```shell
gorush help       # 获取帮助信息
```

## ❤️5. 开发计划
- [x] 提供使用命令行创建项目的脚手架方式。
- [x] 重构项目结构，提供用户友好型文件组织方式。
- [x] 自定义、方便排查的日志系统
- [ ] 脚手架新建项目结构的完善
- [ ] 引入gin框架以优化流程
- [ ] 优化错误处理和日志系统
- [ ] 生成静态网页，方便博客托管
- [ ] 实现文章分类

...