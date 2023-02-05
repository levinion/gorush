# GoRush

![License](https://img.shields.io/badge/license-Apache--2.0-orange)

![Language](https://img.shields.io/badge/language-go-brightgreen)

## 😀1. 介绍

一个为开发者自由、快速、便捷地构建网站服务的网站框架。

## 💘2. 目的

构建这个项目的本意是提供一种方便、快速的自由构建网站方式，让使用者能够根据自己的需求diy网站，并且力图将门槛降到最低。但是，即使对于一个简单的静态网页来说，基本的html和css技能是必要的。

目前个人对该项目的构想经常变化，很难想象怎样的一种方式能够更好地达成这一目的，而最终成果也十分令人期待:star:

## ✨3. 下载

安装最新版本的脚手架

```shell
go install github.com/levinion/gorush@latest
```

## 🔥4. 使用

### 4.1 命令行

```sh
gorush help       # 获取帮助信息

gorush get        # 初始化资源——仅需在项目初始化时调用

gorush build      # 从模板构建 docs 公开文件夹

gorush serve      # 构建和运行服务——该命令已包括 build 流程

gorush clean      # 快速清理项目文件

gorush new        # 新建文章或页面
```
具体使用方法请见 `gorush help`

### 4.2 项目结构

使用 `gorush get` 初始化得到项目结构如下（仅显示目录）：

```shell
.
├── assets                 # 静态资源文件夹
│   └── css
├── config.toml            # 配置文件
├── content                # 用户工作区
│   ├── pages              
│   └── posts
│       └── drafts         # 草稿，构建时将默认忽略此处的内容
├── License                # 默认使用的License
├── README.md              # 空README文件
├── static                 # static文件夹内资源将以原样生成在docs文件夹根目录，暂不能存在目录
└── templates              # 模板文件夹，用户将在此自定义页面样式
    ├── common             # 通用的一些模板，包括导航区、头文件等
    ├── pages              # 全部页面模板
    └── posts              # 文章的一些模板，默认为Basic
        └── Basic

```

### 4.3 配置文件

配置文件默认为`config.toml`，位于项目主目录下，在程序运行前需要先行建立配置文件。若已使用了`gorush get`将生成下列的配置文件：

```toml
[pages]
contentOnlyPages=["about"]  //基础页面列表

[posts]
excludeDirs=["drafts"]      //文章构建所忽略的目录

[server]
addr="localhost:9090"       //监听地址
```

### 4.4 项目建立流程

1. 新建文件夹，在该文件夹下进行先期准备：

```shell
git init        # 若不使用git请跳过

gorush get      # 初始化项目
```

2. 在 `content/posts` 文件夹中创建一些有价值的内容，建议只使用 `gorush new` 命令创建。该命令将自动生成创建时间等元数据。具体使用方法如下：

```sh
# 创建一个新文章
gorush new 分类/文章名称  
```
请一定注意在新建文章时分配其所在的分类，否则将直接在posts文件夹下生成，构建时将忽略对其的处理；你当然也可以在构建前将其移动到合适的文件夹中。使用 `-p` 或 `--page` 指定创建一个新页面。
```sh
# 创建一个新页面
gorush new -p 页面名称  
```
3. 使用 `gorush serve` 在浏览器对应的端口预览，默认为"localhost:9090"

4. 若您想要发布到平台上（如 Github Pages），在确认无误的情况下也可以直接使用 `gorush build` 构建。

5. 警告：请不要在任何情形下操作 docs 文件夹，因为该文件夹将在运行 `gorush serve` 或 `gorush build` 时自动删除和重构。对想要更改网页内容的用户来说，应修改 `templates` 下的模板文件。docs文件夹应只作为发布用途使用。

### 4.5 工作流程
```sh
mkdir xxx & cd xxx                       //新建工作文件夹
gorush get                               //初始化项目
gorush new {{category}}/{{post}}         //新建文章
writing ...                              //具体文章的写作过程
gorush serve                             //启动服务并在浏览器中预览
```

## ❤️5. 开发计划

- [x] 提供使用命令行创建项目的脚手架方式。
- [x] 重构项目结构，提供用户友好型文件组织方式。
- [x] 自定义、方便排查的日志系统
- [x] 脚手架新建项目结构的完善
- [x] ~~引入gin框架以优化流程~~使用 `http/template` 进行解析
- [x] 优化错误处理和日志系统
- [x] 生成静态网页，方便博客托管
- [x] 实现文章分类

...

## 6. 设计理念和实现细节

### 6.1 GoRush的设计理念

由模板生成页面，模板为社区上传，项目中只保留基础模板（二进制嵌入）。

文章作为内容的载体，其目的应当是清晰的；因此使用单级标签——也就是唯一分类保证其内容的纯洁性。在我使用Hugo时（Hugo是一个非常优秀的框架），厌倦了标签和分类的分配工作。当你习惯了单分类，你会发现这样的方式是简洁高效的。

用户应当是内容的提供者，具有部分的自由性，即：
    - 用户所能定义的数据由开发者决定
    - 用户通过开发者提供的构建文章命令开始写作
    - 开发者应当给出标准的接口供模板编写者使用
因此，开发者提供的只能是最基础的接口，但要能够将全部写作过程概括其中。

## 6.2 具体实现方法

1. 使用 `gin` 框架，利用用户所提供内容填充模板(后选择改用 `html/template` 替代 `gin` )
    模板放在同一文件夹内，具体结构如下
    ```sh
    -project
        -templates
            -template1
                -index.html
            -template2
                -index.html
            ...
    ```
2. 使用 Cobra 构建命令行程序
   - 命令行能够新建文章，使用户统一操作
   - 命令行能够导入基础模板，使未引入社区模板（未来目标）的用户能够实现基本写作需求

3. 选用 GoldMark 作为 Markdown 解析工具

4. 使用 embed 进行静态资源嵌入

5. 主要的工作流均由Builder完成，它负责文件的解析读取、构建（静态化）、路由的相关工作，其中文件解析和存取采取并发机制，拥有较好的性能。