---
title: "Linux的51单片机开发环境搭建"
date: 2022-11-27T23:21:39+08:00
draft: false
tags: ["Linux","单片机"]
---

因为最近新开了一门 51 单片机课程，所以搞了台单片机开发板来玩。传统的单片机开发环境基本上都是 Windows+Keil，而 Keil 的编辑器又显得落后。如果你也想要脱离 Windows 和 Keil 那种传统的开发生态，拥抱 Linux 的开发环境的话，不妨尝试 PlatformIO。

## 1 PlatformIO
PlatformIO 是一个物联网开发的开源生态系统，支持跨平台的开发和调试。我平时使用 VSCode 作为主力编辑器；PlatformIO 插件对 VSCode 的支持很好。当然，如果你使用其他编辑器或 ide，那也当然没问题；PlatfromIO 在 Clion、QT、Vim 等主流软件中均由相关插件支持。这里以 VSCode 为例：

1. 首先安装 PlatformIO IDE 插件

![](https://ghproxy.com/github.com/levinion/blog-pic/blob/main/img/Pasted%20image%2020221127234938.png)

2. 打开 PlatfromIO IDE 首页

![](https://ghproxy.com/github.com/levinion/blog-pic/blob/main/img/Pasted%20image%2020221128000736.png)

3. 从 Platfroms 中搜索 Intel MCS-51 (8051) 进行下载

   ![image-20221128004948062](https://ghproxy.com/github.com/levinion/blog-pic/blob/main//img/image-20221128004948062.png)

4. 点击新建项目

   ![image-20221128005430885](https://ghproxy.com/github.com/levinion/blog-pic/blob/main//img/image-20221128005430885.png)

5. 填写项目名称；选择 STC89C52RC；进入下一步

![image-20221128005630955](https://ghproxy.com/github.com/levinion/blog-pic/blob/main//img/image-20221128005630955.png)

6. include 文件夹下存放项目需要的头文件；主程序放在 src 目录下；platformio.ini 是 PlatformIO的配置文件；

![image-20221128010233061](https://ghproxy.com/github.com/levinion/blog-pic/blob/main/img/image-20221128010233061.png)

## 2. hello world
正如每种语言的入门程序一定是 hello world 那样，学习 51 单片机的初次实践必定是点亮 led 灯。让我们向 src 中添加以下代码：
```c
//main.c
#include "8051.h"
#define LED P1_0

void main(){
    LED = 0;
}
```

你可能会注意到这个程序和你在 Keil 中见到的并不完全一样（也就是说和大多数的教程并不相同），但没有任何关系；PlatformIO 使用的编译器是 sdcc，因此和 Keil 具有不同的指令，比如定义端口时使用 “ \_ ” ，而不是 “ \^ ” ；又比如使用宏定义而不是“ sbit ”...


## 3. 确认串口驱动
一般情况下，Linux 自带串口驱动。但为了确保驱动切实存在，运行以下命令（当然是在已连接单片机的情况下）

```bash
sudo dmesg | grep tty
```

![image-20221128012501552](https://ghproxy.com/github.com/levinion/blog-pic/blob/main/img/image-20221128012501552.png)

当看到方框中信息表明驱动已安装。万一（<s>真的会有这个万一吗？</s>）没能看到此条消息，请先安装驱动，不要往下走。

## 4. 构建和烧录

![image-20221128014316253](https://ghproxy.com/github.com/levinion/blog-pic/blob/main/img/image-20221128014316253.png)

Platform IO 中已经事先预留好了构建和烧录选项，一般来说我们只需要鼠标点点就行，但还是稍微了解一下其中原理。

Platform IO 构建过程使用的编译器，正如我们之前所说，是 sdcc。其全称是 Small Device C Compiler，这是一个使用 GPL 授权的开源工具集；这也意味着我们能够毫无顾虑地将其投入到生产中（点名批评某 K 字开头的软件）。其指令与 Keil 稍有不同。

烧录使用的是 stcgal。stcgal 是 STC Windows 软件的全功能开源替代品；它支持广泛的 MCU，非常便携，适合自动化。使用命令行运行的指令如下：

```bash
stcgal -P stc89 -p /dev/ttyUSB0 .pio/build/STC89C52RC/firmware.hex
```

其中 `-P` 指定单片机的型号，`-p` 指定串口，此处填写上面查找到的串口名称。（有时能够自动找到串口和单片机型号，因此不是必填项）

我所使用的是某中科技的 STC89C52 开发板，已知在使用 Platform IO 自带的 stcgal 进行烧录会卡在 `Cycling power: done` ，无法继续进行烧录。据判断，是未指定开发板型号。因此使用指令进行烧录；出于偷懒的原因，在 `~/.zshrc` 中添加一条别名：

```bash
alias stc89="stcgal -P stc89 -t 24000 .pio/build/STC89C52RC/firmware.hex"
```

输入指令后，按提示手动断电上电

![image-20221128021600659](https://ghproxy.com/github.com/levinion/blog-pic/blob/main/img/image-20221128021600659.png)

当看到“ Disconnected! ”时表明烧录完毕

![image-20221128021716933](https://ghproxy.com/github.com/levinion/blog-pic/blob/main/img/image-20221128021716933.png)

> 若提示串口 Permission denied ，尝试运行 `sudo chmod 777 /dev/ttyUSB0` 增加权限。

## 参考
1. [stcgal - STC MCU ISP 闪存工具](https://gitee.com/mirrors/stcgal)
2. [Linux操作系统搭建51单片机开发环境（国产桌面操作系统deepin）](https://blog.csdn.net/RYMCU/article/details/111350775)
3. [关于 stcgal 烧写 STC89C52 的问题](https://blog.csdn.net/Narukara/article/details/120623921#:~:text=%E8%A7%A3%E5%86%B3%E6%96%B9%E6%B3%95)
