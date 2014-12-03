## 用户定义配置文件

```
# database
mysqluser = "root"
mysqlpass = "lijun"
mysqlurls = "127.0.0.1"
mysqlport = 3306
mysqldb   = "blog"

# memcached
memcache_host = "127.0.0.1:11211"

registorable = true
adminemail = "blog_rex@163.com"
adminemailpass = "xxxxxxxx"
adminemailhost = "smtp.163.com:25"

# oss
oss_self_domain = true
oss_get_host = "media.duguying.net"
oss_host = "oss-cn-qingdao.aliyuncs.com"
oss_internal = false
oss_host_internal = "oss-cn-qingdao-internal.aliyuncs.com"
oss_id = "xxxxxxxxxxxxxxxx"
oss_key = "xxxxxxxxxxxxxxxxxxxxxxxxxxxxxx"
oss_bucket = "duguying"

# duoshuo comment and logo
duoshuo_short_name = "duguying"
logo = "http://gravatar.duoshuo.com/avatar/5fedd018b3227bc4043934309da8c290"

```

关于memcache_host，其为可选选项，若未开启不影响使用。

关于oss，若是开启oss_self_domain，则可以通过oss_get_host配置oss自定义域名。否则请将oss_self_domain设置为false.
