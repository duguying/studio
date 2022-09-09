package action

import (
	"duguying/studio/service/action/agent"
	"duguying/studio/service/middleware"

	"github.com/gin-gonic/gin"
)

func SetupFeAPI(api *gin.RouterGroup) {
	api.GET("/get_article", GetArticleShow)                         // 获取文章详情
	api.GET("/article/view_count", ArticleViewCount)                // 文章浏览统计
	api.GET("/list", ListArticleWithContent)                        // 列出文章
	api.GET("/list_tag", ListArticleWithContentByTag)               // 按Tag列出文章
	api.GET("/list_archive_monthly", ListArticleWithContentMonthly) // 按月归档文章内容列表
	api.GET("/list_title", APIWrapper(ListArticleTitle))            // 列出文章标题
	api.GET("/search_article", SearchArticle)                       // 搜索文章
	api.GET("/hot_article", HotArticleTitle)                        // 热门文章列表
	api.GET("/month_archive", MonthArchive)                         // 文章按月归档列表
	api.POST("/user_register", UserRegister)                        // 用户注册
	api.GET("/user_simple_info", APIWrapper(UserSimpleInfo))        // 用户信息
	api.POST("/user_login", APIWrapper(UserLogin))                  // 用户登陆
	api.GET("/username_check", UsernameCheck)                       // 用户名检查
	api.POST("/2fa", TfaAuth)                                       // 2FA校验
	api.GET("/sitemap", SiteMap)                                    // 列出所有文章URI
	api.POST("/save_error_logger", APIWrapper(SaveErrorLogger))     // 前端错误日志
}

func SetupAdminAPI(api *gin.RouterGroup) {
	api.GET("/user_info", APIWrapper(UserInfo))                 // 用户信息
	api.POST("/user_logout", APIWrapper(UserLogout))            // 用户登出
	api.PUT("/change_password", APIWrapper(ChangePassword))     // 修改密码
	api.GET("/login_history", APIWrapper(ListUserLoginHistory)) // 列举用户登录历史
	api.GET("/message/count", APIWrapper(UserMessageCount))     // 用户消息计数
	api.POST("/put", APIWrapper(PutFile))                       // 上传文件
	api.POST("/upload", APIWrapper(UploadFile))                 // 上传归档文件
	api.Any("/xterm", ConnectXTerm)                             // 连接xterm

	api.POST("/article", APIWrapper(AddArticle))                      // 添加文章
	api.PUT("/article", APIWrapper(UpdateArticle))                    // 修改文章
	api.PUT("/article/publish", APIWrapper(PublishArticle))           // 发布草稿
	api.DELETE("/article", APIWrapper(DeleteArticle))                 // 删除文章
	api.GET("/article/list_title", APIWrapper(ListAdminArticleTitle)) // 列出文章列表
	api.GET("/article", APIWrapper(AdminGetArticle))                  // 获取文章
	api.GET("/article/current_md5", APIWrapper(GetArticleCurrentMD5)) // 获取文章当前内容MD5

	api.GET("/2faqr", QrGoogleAuth)                      // 获取2FA二维码
	api.POST("/upload/image", APIWrapper(UploadImage))   // form 表单上传图片
	api.DELETE("/file/delete", APIWrapper(DeleteFile))   // 删除文件
	api.GET("/file/list", APIWrapper(PageFile))          // 文件列表
	api.PUT("/file/sync_cos", APIWrapper(FileSyncToCos)) // 文件同步到 cos

	api.GET("/album/list", APIWrapper(ListAlbumFiles))
}

func SetupAgentAPI(api *gin.RouterGroup) {
	api.GET("/list", middleware.SessionValidate(true), agent.List)             // agent列表
	api.DELETE("/remove", middleware.SessionValidate(true), agent.RemoveAgent) // agent删除（禁用）
	api.Any("/ws", agent.Ws)                                                   // agent连接点
}
