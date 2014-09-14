blog
----------
基于beego的博客

# Install #

```shell
go get github.com/beego/bee
go get github.com/yanunon/oss-go-api
git clone https://github.com/duguying/blog
bee run
```

# 功能单元及完成进度 #

## 后端 ##

- [ ] 用户管理
	- [x] 登录
	- [x] 注册
	- [x] 登出
	- [x] 修改用户名
	- [x] 修改邮箱
	- [ ] 邮箱验证
	- [x] 找回密码
		- [x] 发送找回密码邮件
		- [x] 重置密码
	- [x] 修改密码
	- [ ] 销户
- [x] 文章管理
	- [x] 添加文章
	- [x] 修改文章
	- [x] 删除文章
	- [x] 获取文章
	- [x] 文章分页
	- [x] 最热文章列表
	- [x] 所有文章列表页
	- [x] 文章按月份分页
	- [x] 文章按关键词分页
	- [x] 文章阅读统计
	- [x] 管理-文章列表
- [ ] 其他
	- [x] 附件上传
	- [x] 附件数据库记录
	- [ ] 附件列表
	- [ ] 附件删除
	- [x] 阿里云OSS上传
	- [x] 文章草稿保存
	- [x] 文章草稿获取

## 前端 ##

- [x] 整体风格
    - [x] 首页样式
    - [x] 管理后台样式
- [ ] 交互设计
    - [x] 代码高亮
    - [ ] 公式高亮
    - [x] 文章页多说评论框
    - [ ] 管理页编辑器配置
    - [x] 管理页文章添加
    - [x] 管理页文章删除
    - [x] 管理页文章修改

### Attachment ###

配置好的ueditor[下载](http://duguying.oss-cn-qingdao.aliyuncs.com/attach/ueditor.zip "下载")，请下载后解压到`static`目录下。

# License #

MIT License
