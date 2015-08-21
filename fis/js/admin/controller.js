"use strict";

function NavsController ($scope, $http, $location) {
	$http.get("/api/admin/navlist",null).success(function (data) {
		$scope.navs = data;
		// $scope.isActive = true;
		$scope.currentPath = $location.path().replace("/admin","").replace("/","")
	});
}

function IndexController($scope,$rootScope){
	// $scope.active = true
    console.log("hello index");
    $rootScope.global = {
    	title: "首页",
    	currentPath: ""
    }
    
}

function NewArticleController($scope,$rootScope){
    console.log("hello new article");
    $rootScope.global = {
    	title: "添加文章",
    	currentPath: "new_article"
    }
    $scope.config = {
        toolbars: [
            [
                'source', //源代码
                'forecolor', //字体颜色
                'backcolor', //背景色
                'insertorderedlist', //有序列表
                'insertunorderedlist', //无序列表
                'simpleupload', //单图上传
                // 'insertimage', //多图上传
                '|',
                'justifyleft', //居左对齐
                'justifyright', //居右对齐
                'justifycenter', //居中对齐
                'justifyjustify', //两端对齐
                '|',
                'inserttable', //插入表格
                'fontfamily', //字体
                'fontsize', //字号
                'paragraph', //段落格式
                'insertcode', //代码语言
            ],
            [
                'bold', //加粗
                'italic', //斜体
                'underline', //下划线
                'strikethrough', //删除线
                '|',
                'subscript', //下标
                'fontborder', //字符边框
                'superscript', //上标
                // 'formatmatch', //格式刷
                '|',
                // 'anchor', //锚点
                'undo', //撤销
                'redo', //重做
                
                'indent', //首行缩进
                'snapscreen', //截图

                
                'blockquote', //引用
                'pasteplain', //纯文本粘贴模式
                'selectall', //全选
                // 'print', //打印
                'preview', //预览
                'horizontal', //分隔线
                'removeformat', //清除格式
                'time', //时间
                'date', //日期
                'unlink', //取消链接
                // 'insertrow', //前插入行
                // 'insertcol', //前插入列
                // 'mergeright', //右合并单元格
                // 'mergedown', //下合并单元格
                // 'deleterow', //删除行
                // 'deletecol', //删除列
                // 'splittorows', //拆分成行
                // 'splittocols', //拆分成列
                // 'splittocells', //完全拆分单元格
                // 'deletecaption', //删除表格标题
                // 'inserttitle', //插入标题
                // 'mergecells', //合并多个单元格
                // 'deletetable', //删除表格
                // 'cleardoc', //清空文档
                // 'insertparagraphbeforetable', //"表格前插入行"
                
                // 'edittable', //表格属性
                // 'edittd', //单元格属性
                'link', //超链接
                'emotion', //表情
                'spechars', //特殊字符
                'searchreplace', //查询替换
                'map', //Baidu地图
                // 'gmap', //Google地图
                'insertvideo', //视频
                'help', //帮助
                
                
                'fullscreen', //全屏
                // 'directionalityltr', //从左向右输入
                // 'directionalityrtl', //从右向左输入
                // 'rowspacingtop', //段前距
                // 'rowspacingbottom', //段后距
                // 'pagebreak', //分页
                // 'insertframe', //插入Iframe
                // 'imagenone', //默认
                // 'imageleft', //左浮动
                // 'imageright', //右浮动
                // 'attachment', //附件
                // 'imagecenter', //居中
                // 'wordimage', //图片转存
                // 'lineheight', //行间距
                'edittip ', //编辑提示
                // 'customstyle', //自定义标题
                // 'autotypeset', //自动排版
                // 'webapp', //百度应用
                'touppercase', //字母大写
                'tolowercase', //字母小写
                // 'background', //背景
                // 'template', //模板
                // 'scrawl', //涂鸦
                'music', //音乐
                
                'drafts', // 从草稿箱加载
                'charts', // 图表
            ]
        ],
        autoHeightEnabled: false,
        autoFloatEnabled: false
    };
}

function ManageArticleController($scope,$rootScope){
    console.log("hello manage article");
    $rootScope.global = {
    	title: "管理文章",
    	currentPath: "manage_article"
    }
}

function ManageProjectController($scope,$rootScope){
    console.log("hello manage project");
    $rootScope.global = {
    	title: "管理项目",
    	currentPath: "manage_project"
    }
}

function ManageOssController($scope,$rootScope){
    console.log("hello manage oss");
    $rootScope.global = {
    	title: "OSS管理",
    	currentPath: "manage_oss"
    }
}


