---
title: "Grub引导恢复历险"
date: 2022-09-17T13:26:26+08:00
draft: false
tags:
categories: ["ARCH"]
---
## 事件爆发
由于 Windows10 的一次更新，成功失去了 grub 引导。这是一次历史性的事件——因为我再次回到了原本单 Windows 系统的畅快感之中：我再也不需要进行选择以进入 Windows 了。作为世界的管理员，我无法容忍 Arch 的丧失，于是踏上了寻回 grub 的旅程。

## 旅途经过
### 制作 Arch 启动盘
作为一名管理员，我深知自己需要做什么。

首先我需要钥匙——『启动盘』来开启进入 Arch 的大门。我先在 Windows 上下载了最新的 Arch 镜像，并且用 Rufus 将其烧录入 U 盘。

> Arch 镜像可以很方便地从 [中科大开源镜像站](http://mirrors.ustc.edu.cn/) 中得到; 磁盘制作工具推荐 [Rufus](https://rufus.ie/zh/#)，或者直接安装[图吧工具箱](http://www.tbtool.cn/index.html)，其中带有 Rufus 及其他 Windows 下好用的工具。

### 正式修复开始啦！
插上 U 盘，设置 Bios 从 U 盘启动，我就再次进入到了 Arch 的世界。那虽然是一个什么都没有的、远古的世界，我还是开始了紧张的工作。

#### 查看磁盘分区
`fdisk -l` 命令是一个令人兴奋的工具，键入后我即得到了一张清楚的磁盘地图。

```shell
Device             Start        End   Sectors   Size Type  
/dev/nvme0n1p1      2048     534527    532480   260M EFI System  
/dev/nvme0n1p2    534528     567295     32768    16M Microsoft reserved  
/dev/nvme0n1p3    567296  210282495 209715200   100G Microsoft basic data  
/dev/nvme0n1p4 210282496  482912255 272629760   130G Microsoft basic data  
/dev/nvme0n1p5 998166528 1000214527   2048000  1000M Windows recovery environment  
/dev/nvme0n1p6 482912256  503883775  20971520    10G Linux swap  
/dev/nvme0n1p7 503883776  998166527 494282752 235.7G Linux filesystem
```

一旦拥有了这张地图，我就再也不会迷路了。

> grub 恢复过程中最重要的当属 Arch 的根目录以及 efi 目录。其中根目录所在的分区为 /dev/nvme0n1p7，efi 分区为 /dev/nvme0n1p1。具体请根据 `fdisk -l` 所显示的 Type 选择。

#### 目录的挂载
分区挂载的情况各不相同，具体可以参照自己安装时的挂载方式。由于我采用了 btrfs 文件系统，在进行这项工作的时候绕了弯路。

我记录在羊皮纸上的挂载过程如下：

```shell
# 挂载根目录
mount -t btrfs -o subvol=/@,compress=zstd /dev/nvme0n1p7 /mnt

# 挂载家目录
mount -t btrfs -o subvol=/@home,compress=zstd /dev/nvme0n1p7 /mnt/home

# 挂载 efi 目录
mount /dev/nvme0n1p1 /mnt/boot/efi
```

至此我就已完成了所有需要的目录挂载。

> 正如前面所说，各人的挂载方式应当以安装过程中的挂载方式为准。如果使用的是 btrfs 文件系统，可以采用以上方式进行挂载；如果是传统的 ext4，建议参考以下过程。但因为我没有实践，所以仍有无法解决问题的风险，建议一并参考其他资料。
```shell
# 挂载根目录
mount /dev/nvme0n1p7 /mnt

# 挂载 boot 目录
mount /dev/nvme0n1p1 /mnt/boot
```
#### 重建引导
是时候让秩序恢复正常了！
1. 切换系统环境
   ```shell
   arch-chroot /mnt
   ```
1. （重新）安装相关的包
   ```shell
   pacman -S grub efibootmgr os-prober
   ```
3. 安装 GRUB 到 EFI 分区
   ```shell
   grub-install --target=x86_64-efi --efi-directory=/boot/efi --bootloader-id=ARCH
   ```
4. 生成 grub 配置文件
   ```
   grub-mkconfig -o /boot/grub/grub.cfg
   ```

至此所有工作就结束了，嗯，让我们先来检查一下工作的成果。

```shell
Generating grub configuration file ...  
Found theme: /usr/share/grub/themes/Vimix/theme.txt  
Found linux image: /boot/vmlinuz-linux-zen  
Found initrd image: /boot/amd-ucode.img /boot/initramfs-linux-zen.img  
Found fallback initrd image(s) in /boot:  amd-ucode.img initramfs-linux-zen-fallback.img  
Found linux image: /boot/vmlinuz-linux  
Found initrd image: /boot/amd-ucode.img /boot/initramfs-linux.img  
Found fallback initrd image(s) in /boot:  amd-ucode.img initramfs-linux-fallback.img  
Warning: os-prober will be executed to detect other bootable partitions.  
Its output will be used to detect bootable binaries on them and create new boot entries.  
Found Windows Boot Manager on /dev/nvme0n1p1@/EFI/Microsoft/Boot/bootmgfw.efi  
Adding boot menu entry for UEFI Firmware Settings ...  
done
```

> 正常的输出应该是像上面一样。如果你没有看见 Windows 引导项，不要慌——因为我在重建引导的时候也没有看到（笑）。请接着往下看。

#### 收尾工作
```shell
exit # 退出
umount -R /mnt #卸载分区
reboot #重启
```

然后你应该就能看到熟悉的 grub 界面。

>如果你失去了 Windows 登录选项，那么这可能是来自 Arch 的一个小小的报复——开玩笑的。你只需要进入到 Arch，打开终端，并且重新生成 grub 配置文件即可。

```shell
 sudo grub-mkconfig -o /boot/grub/grub.cfg
```

## 一份很认真的检讨书
这里是某台计算机的管理员，我对这次事故表示十分抱歉。但我仍然认为，惩罚应当落在 Windows 头上。因为我所犯的唯一错误即是升级了 Windows。作为此次事件的教训，我决定禁止 Windows 的一切更新，以不使同样的事故再次发生！