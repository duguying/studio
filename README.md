# blog [![Build Status](https://travis-ci.org/duguying/blog.svg)](https://travis-ci.org/duguying/blog)

基于beego的博客

## Build ##

go version 1.3+

```shell
go get
go build
```

## Attachment ##

配置好的ueditor[下载](http://duguying.oss-cn-qingdao.aliyuncs.com/attach/ueditor.zip "下载")，请下载后解压到`static`目录下。

## Install ##

你可以直接导入`blog.sql`文件到你的mysql数据库，也可以访问`/install`页面自动导入数据库。
之后按照`Attachment`部分安装ueditor。
最后，将`conf/app.conf`配置为可注册(`registorable = true`)，然后访问`/registor`注册你的第一个帐号，记得注册完成后再配置文件中关闭可注册状态(`registorable = false`)。

## Config ##

请看custom目录下的[README.md](https://github.com/duguying/blog/tree/master/custom "config")。

## License ##

MIT License