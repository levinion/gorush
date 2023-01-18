---
title: "Mysql笔记"
date: 2022-10-10T09:45:10+08:00
draft: true
---
## MySQL 安装

## 使用
### 新建
1. 创建数据库
   `create database [name of database]`
2. 创建数据表
   `create table [name of table]`
### 选择
1. 选择全部
   `select * from [table name]`
2. 选择列
   `select [列1],[列2],[列3]... from [table name]`
3. 去重
   `select distinct [列] from [table name]`
4. 限制数目
   `select ... from ... limit 2`
5. 别名
   `select ... as [别名] from ...`
6. 条件选择
   `select ... from ... where [条件]`
7. between and
   用以选择某列在两值之间的条目
   `select ... from ... where [某项] between ... and ...`
8. 不为空
   `where [条目] is not null`
9. in / not in
   `where [条目] in (具体条目)`
10. like
    模糊选择 (\[\] in  /  \[^ \] not in   /   % 多个     /   _ 一个)
    `where university like '北京%'`
11. max
    `select max(gpa) as gpa from ... where ...`
12. count, avg, rounnd (圆整的数, 小数点后位数)
13. group by 分组查询
14. having 使用别名判断
15. order by 字段 asc/desc
16. join ... on ...
