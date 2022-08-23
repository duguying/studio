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
	api.GET("/list_title", ListArticleTitle)                        // 列出文章标题
	api.GET("/search_article", SearchArticle)
	api.GET("/hot_article", HotArticleTitle)     // 热门文章列表
	api.GET("/month_archive", MonthArchive)      // 文章按月归档列表
	api.POST("/user_register", UserRegister)     // 用户注册
	api.GET("/user_simple_info", UserSimpleInfo) // 用户信息
	api.POST("/user_login", UserLogin)           // 用户登陆
	api.GET("/username_check", UsernameCheck)    // 用户名检查
	api.GET("/file/list", PageFile)              // 文件列表
	api.POST("/2fa", TfaAuth)                    // 2FA校验
	api.GET("/sitemap", SiteMap)                 // 列出所有文章URI
}

func SetupAdminAPI(api *gin.RouterGroup) {
	api.GET("/user_info", UserInfo)      // 用户信息
	api.POST("/user_logout", UserLogout) // 用户登出
	api.POST("/put", PutFile)            // 上传文件
	api.POST("/upload", UploadFile)      // 上传归档文件
	api.Any("/xterm", ConnectXTerm)      // 连接xterm

	api.POST("/article", APIWrapper(AddArticle))            // 添加文章
	api.PUT("/article", APIWrapper(UpdateArticle))          // 修改文章
	api.PUT("/article/publish", APIWrapper(PublishArticle)) // 发布草稿
	api.DELETE("/article", APIWrapper(DeleteArticle))       // 删除文章
	api.GET("/article/list_title", ListArticleTitle)        // 列出文章列表
	api.GET("/article", GetArticle)                         // 获取文章

	api.GET("/2faqr", QrGoogleAuth)        // 获取2FA二维码
	api.POST("/upload/image", UploadImage) // 上传图片
}

func SetupAgentAPI(api *gin.RouterGroup) {
	api.GET("/list", middleware.SessionValidate(true), agent.List)             // agent列表
	api.DELETE("/remove", middleware.SessionValidate(true), agent.RemoveAgent) // agent删除（禁用）
	api.Any("/ws", agent.Ws)                                                   // agent连接点
}
