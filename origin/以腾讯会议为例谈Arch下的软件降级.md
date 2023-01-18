---
title: "以腾讯会议为例谈Arch下的软件降级"
date: 2022-09-15T13:53:18+08:00
draft: false
tags: []
categories: ["ARCH"]
---

## 引入

首先我想谈一谈为什么会有这篇文章。如果所有的软件都能在最新版本平稳运行，那么软件降级就绝非必要之举了。无论是 Bug 或是其他的一些兼容问题，更多的情况是为解决一些旧的缺陷反而导致了新的更大的问题。在 linux 平台下，Bug 的修复可以说是令人欲哭无泪——尤其是 aur 库中的软件。这些软件通常由个人或团队进行维护，在 Bug 产生时，往往需要数个星期甚至数个月才会进行修复；而这些时间对于个人使用者来说是十分煎熬的。

在 Arch 下的腾讯会议解决方案通常有基于 deepin-wine 的 wemeet，以及腾讯官方的 linux 客户端。参看 aur 库中的数据，前者已经有大半年未更新，更多的使用者转移到了后者上。linux 原生的腾讯会议虽然存在功能缺失（在我的 PC 上无法进行签到），但也能够实现大部分需求。

在我的腾讯会议升级到 3.10 版本后，出现了页面卡死以及无法点击的情况，经常需要切换应用以重新获取焦点。很多人表示在升级后遇到了同样的情况，因此不得不暂时将其降级。

## 具体实现

### 一种常用的降级方式

downgrade 是 Arch 的一个很有名的软件降级应用，虽然无法使用在本次的降级中，但还是稍微谈一下。

运行 `paru -S downgrade` 即可下载相应软件，运行也十分简单——只需 `sudo downgrade 包名` 即可。然后你就可以从菜单中选择想要降级的版本。

### 本次采用的降级方法

本次使用的降级方法，一言以蔽之——手动 makepkg 构建。如果你理解了，那么就不需要往下看。如果你没能理解，那么也许能作为参考。

#### 从 aur 库中查找到旧版本包

腾讯会议的包链接如下：[wemeet-bin](https://aur.archlinux.org/packages/wemeet-bin?O=0)

点进去后，从 view changes 中可以看到包的更改历史。选择相应的版本，并从 download 中的链接下载压缩包。

#### 后续操作

将下载得到的压缩包解压，进入到解压缩后的文件夹，并在其中打开终端，运行

```shell
makepkg
```

如果原软件仍存在，则先将软件卸载

```shell
paru -R wemeet-bin
```

然后直接安装即可

```shell
paru -U *.pkg.tar.zst
```

最后编辑 /etc/pacman.conf 文件，以禁止 wemeet-bin 更新（此处以vim为例）

```shell
sudo vim /etc/pacman.conf
```

部分输出如下

```shell
# Pacman won't upgrade packages listed in IgnorePkg and members of IgnoreGroup
# IgnorePkg   =
# IgnoreGroup =
```

修改 IgnorePkg 配置项并填充包名，输入 `:wq` 以保存退出

```shell
IgnorePkg   = wemeet-bin
```

重启终端即可禁止 wemeet-bin 更新，发现运行 `paru -Syu` 有以下提示：

```shell
warning: wemeet-bin: ignoring package upgrade (3.9.0.1-1 => 3.10.0.400-1)
```

## 补充

1. 本文因为习惯问题使用 paru ，你当然也可以使用 yay 或 pacman ，具体使用方法几乎没有区别。

2. 如果你想要禁止多个软件包更新，只需在 IgnorePkg 配置项中以空格（Space）分隔包名，如：

   ```shell
   IgnorePkg = 包名1 包名2 包名3, ...
   ```