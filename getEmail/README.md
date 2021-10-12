# 邮箱采集脚本

邮箱采集是提取这个网站的`https://app.snov.io/`的搜索结果。

## 各个目录简介

lib是基础的工具目录

里面有read 和 report两个文件。

- read里面Readua函数主要是随机获取user-Agent，也可以根据自身需求写其他函数，修改相应的head头。
- report里面封装的是get和post请求。

snov目录下的getValue就是获取email的文件

里面有三个函数

GetToken 获取token，为之后的请求做准备

GetEmail 获取Email，获取到带email的数据

JsonToSlice JsonToMap 把json数据解析然后打印出来，还有必要的lastId，都会打印出来。

## 使用指南

jsonValue函数主要传入两个参数，第一个是需要搜索的域名，第二个lastid，主要是用做进行下一次搜索的起点，第一次默认是0.

第二个可能需要修改的点是snov文件夹下的第一个函数getValue函数的

~~~
	data["client_id"] = []string{"d47803b976c4a49c450ef673f6594a17"}
	data["client_secret"] = []string{"90dbd0018dd18eae5a3e1c3e9c3c8c64"}
~~~

如果程序出错，很有可能是额度不够（送的），我们可以在用新邮箱注册，然后找到这个更换，继续爬取，这两个东西存在的url`https://app.snov.io/account#/api`





最后附带一个这个网站的使用文档

~~~
https://snovio.cn/api?lang=zh#Authentification
~~~

我在这里用的接口主要是  电子邮件查找器---->域名搜索 V.2

