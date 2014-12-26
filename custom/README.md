## 用户自定义配置

```
# 运行模式
runmode = prod

# mysql
mysqluser = "username"
mysqlpass = "password"
mysqlurls = "127.0.0.1"
mysqlport = 3306
mysqldb   = "blog"

# memcache
memcache_host = "127.0.0.1:11211"

# blog admin
registorable = true
adminemail = "mail@163.com"
adminemailpass = "mail_password"
adminemailhost = "smtp.163.com:25"

# aliyun oss
oss_self_domain = false
oss_get_host = "media.domain.net"
oss_host = "oss-cn-qingdao.aliyuncs.com"
oss_internal = false
oss_host_internal = "oss-cn-qingdao-internal.aliyuncs.com"
oss_id = "xxxxxxxxxxxxxxxx"
oss_key = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
oss_bucket = "bucket_name"

# gravatar
duoshuo_short_name = "duoshuo_name"
logo = "http://gravatar.duoshuo.com/avatar/xxxxxxxxxxxxxxxxxxxxxxxxxxx"

github_token = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
github_user = "duguying"
github_statistics = "./static/upload/data.json"
```

关于memcache_host，其为可选选项，若未开启不影响使用。

关于oss，若是开启oss_self_domain，则可以通过oss_get_host配置oss自定义域名。否则请将oss_self_domain设置为false.

此处的配置可以覆盖`conf/app.conf`的默认配置。
