---
title: "Wine 与 ATRI 与 .desktop"
date: 2022-09-26T19:29:10+08:00
draft: false
categories: ["ARCH"]
---
## 引言
一定是闲得没事的人才会去把 ATRI 放在 Arch 上跑吧（\^\_\^)。平时在 Arch 上面跑 Osu！，而游玩其他游戏只能切换到 Win——这一整套流程想想就头疼。因此试着用 `Wine` 直接跑 ATRI，不想直接跑通，于是有了这篇文章。

![atri](https://ghproxy.com/github.com/levinion/blog-pic/blob/main/img/ss_47e1f809fa56ab6d48de3c66e3147585b4ee30a2.600x338.jpg)

> ATRI 是一个很好的游戏，不过作为一个不怎么推 Gal 的人，对其也并没有多少发言权。本来是可以在 Steam 上购买的，无奈有些小贵，只得自寻旁门外道（笑），等日后补票。虽然我很喜欢开源社区的理念，认为其贯彻了互联网的共享精神，也并不反感所谓盗版；但我也不反对为开发者和游戏创作者的劳动付费。在付费和开源之间寻求平衡，也是互联网发展所必然面对的课题之一——然而今天要谈的不是这些。

## 喂《萝卜子》喝“红酒”
说到 Linux 上的“红酒”，指的当然是大名鼎鼎的“Wine”。Wine 是类 Unix 操作系统上一个功能强大的软件，其提供了一个兼容层用以转译 Windows 系统命令。Wine 名称来自于“Wine Is Not an Emulator”，表明其不是一个模拟器，而其性能损失也确实微乎其微。对游戏改良的 Wine，也就是 Steam Deck 所使用的 proton 配合 Linux 甚至能够达到较 Windows 更低的延迟和更高的帧率。

我的游戏文件夹（除 osu 外）都位于 Windows 分区下的 D 盘，而 Arch 因为采用了 Btrfs 文件系统，能够直接对其进行读写，因此可以直接在该文件夹内运行。

1. 在 Arch 下挂载 Windows 下的 D 盘
   
   > 一般直接在文件管理器（我用的是 dolphin）中点击对应的磁盘，输入密码认证即可。如果厌倦了每次输入密码认证，我给出了解决办法，请继续往后看。
2. 使用 Wine 运行程序
   ```sh
   # 跳转到 exe 程序所在的文件夹
   cd /run/media/maruka/data/games/ATRI.My.Dear.Moments/ATRI.My.Dear.Moments
   # wine 运行 exe程序
   wine ATRI.My.Dear.Moments/ATRI-MyDearMoments-.exe
   ```
   > 我这里 ATRI 的路径是 `/run/media/maruka/data/games/ATRI.My.Dear.Moments/ATRI.My.Dear.Moments/ATRI-MyDearMoments-.exe` <br> 其中，磁盘一般挂载于 `/run/media/UserName/` 目录下，`data` 是我的 D 盘名称。为求方便，可以在文件管理器中找到游戏的 exe 文件，并将路径复制下来。

## 创建一个快捷方式吧！
1. 反正都在终端输入命令运行了，第一发想一定是创建一个 alias。
   ```sh
   # .config/fish/config.fish 或 .zshrc 或 .bashrc
   alias atri="cd /run/media/maruka/data/games/ATRI.My.Dear.Moments/ATRI.My.Dear.Moments;wine ATRI.My.Dear.Moments/ATRI-MyDearMoments-.exe"
   ```
   > 之所以没有采用 `wine /run/media/maruka/data/games/ATRI.My.Dear.Moments/ATRI.My.Dear.MomentsATRI.My.Dear.Moments/ATRI-MyDearMoments-.exe` 的方式运行，是因为本机报了错，大概是无法直接在另一系统硬盘上运行程序吧。所以先 cd 过去再运行。如果你运行出了问题，请尝试这样做。
2. 创建一个 Desktop Entry
   >.desktop 是实现应用程序菜单的文件，通过创建该文件，可以在应用条目中查看到应用，方便使用 krunner 等程序快速唤起。该配置文件分为系统条目和用户条目。系统条目一般包括存于系统中、可供所有用户使用的应用程序，位于 `/usr/share/applications/` 或 `/usr/local/share/applications/` 文件夹中；用户条目优先级高于系统条目，一般位于 `~/.local/share/applications/` 文件夹中。
   ```sh
   # .desktop 文件示例，机翻自 Arch Wiki
   [Desktop Entry]
   
   # 快捷方式类型
   Type=Application
   
   # 此文件遵循的桌面条目规范的版本
   Version=1.0
   
   # 应用名称
   Name=jMemorize
   
   # 可用做提示的内容
   Comment=Flash card based learning tool
   
   # 运行可执行文件的文件夹
   Path=/opt/jmemorise
   
   # 应用程序的可执行文件，可能带有参数。
   Exec=jmemorize
   
   # 将用于显示此条目的图标名称
   Icon=jmemorize
   
   # 描述此应用程序是否需要在终端中运行
   Terminal=false
   
   # 描述应显示此条目的类别
   Categories=Education;Languages;Java;
   ```
   > 该配置文件除 `Type` 和 `Name` 外都是可选的。下面给出我的配置。
   ```sh
   [Desktop Entry]  
    
   Type= Application  
    
   Name= ATRI-MyDear-Moments  
    
   Path= /run/media/maruka/data/games/ATRI.My.Dear.Moments/ATRI. My.Dear.Moments  
    
   Exec= wine ATRI-MyDearMoments-.exe  
    
   ```
   > Type 选择 Application，Name 自定，Exec 指打开快捷方式时运行的命令，Path 表示命令执行的位置（使用Path 的目的与 alias 中给出的理由相同），没什么好解释的，简单易懂。最后，如果你定义的. desktop 文件属于用户条目，即位于 `~/.local/share/applications/` ，执行以下命令以更新桌面条目；否则就不用管。
   ```sh
   update-desktop-database ~/.local/share/applications
   ```

## 偷懒是最大生产力
邓先生所言：“科技是第一生产力。”我对此不以为然。人类从来都是因偷懒而进步着的。将两人份的劳动缩减为一人份，从每个人中间抠出半个人来，让他们可以懒散度日。就连我们的某某主义也是同样，所谓的“人的自由发展”不正是“不干正事”的别名？

回归正题。要以以上方式运行程序的前提是解锁 Windows 系统硬盘的读写，否则将无法找到应用的路径。每次开机解锁输入密码一整套流程，虽然简单，但我不喜欢。查找了一些资料，对硬盘的认证由 `polkit` 完成。通过修改对应的操作（actions）即可解除用户的认证，省去了输入密码的程序。

```sh
# 修改 .policy 文件
sudo vim /usr/share/polkit-1/actions/org.freedesktop.UDisks2.policy

```

找到 `org.freedesktop.udisk2.filesystem-mount-system` 条目，修改 `allow_active` 字段值 `auth_admin_keep` 为 `yes`

```sh
# 修改前
<defaults>  
   <allow_any>auth_admin</allow_any>  
   <allow_inactive>auth_admin</allow_inactive>  
   <allow_active>auth_admin_keep</allow_active>  
 </defaults>

# 修改后
<defaults>  
   <allow_any>auth_admin</allow_any>  
   <allow_inactive>auth_admin</allow_inactive>  
   <allow_active>yes</allow_active>  
 </defaults>
```

重启验证，认证不再需要输入密码。本来想实现开机自动挂载，但是似乎失败了（）

算了，继续推阿托莉去。对游戏的感想就留待以后再写。
