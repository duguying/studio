function NavsController ($scope, $http, $location) {
	// $http.get("/api/admin/navlist",null).success(function (data) {
		// $scope.navs = data;
		$scope.currentPath = $location.path().replace("/admin","").replace("/","");
	// });
}

function IndexController($scope,$rootScope,$http,$sce){
    $rootScope.global = {
    	title: "首页",
    	currentPath: ""
    }

    $http.get("//duoshuo.com/api/posts/list.json?short_name=duguying&order=asc",null).success(function (data) {
        $scope.comments = angular.forEach(angular.fromJson(data.parentPosts), function (comment) {
            comment.message_html = $sce.trustAsHtml(comment.message);
        });
    });
}

function NewArticleController($scope,$rootScope,$http,$location){
    $rootScope.global = {
    	title: "添加文章",
    	currentPath: "new_article"
    }
    // $scope.config = ueditor_option;
    $scope.submit = function () {
        var content = $scope.content;
        var title = $scope.title;
        var keywords = $scope.keywords;
        var abstract = $scope.abstract;
        abstract = (!abstract)?"":abstract;

        $http.post("/api/admin/add", {
                params: {"title":title,"keywords":keywords,"abstract":abstract,"content":content}
            }).success(function(data){
                if (data.result) {
                    alert("add success.");
                    $location.path("/admin/edit_article/"+data["data"]);
                } else{
                    alert("add failed.", data.msg);
                };
            });
    }
}

function EditArticleController($scope,$rootScope,$routeParams,$http) {
    var id = $routeParams.id || 0;
    $scope.id = parseInt(id);

    $http.get("/api/admin/article/"+id,null).success(function (data) {
        $scope.article = data.data;
    });

    $rootScope.global = {
        title: "编辑文章",
        currentPath: "manage_article"
    }
    // $scope.config = ueditor_option;
    $scope.submit = function () {
        var content = $scope.article.Content;
        var title = $scope.article.Title;
        var keywords = $scope.article.Keywords;
        var abstract = $scope.article.Abstract;

        $http.post("/api/admin/update", {
                params: {"id":$scope.id,"title":title,"keywords":keywords,"abstract":abstract,"content":content}
            }).success(function(data){
                if (data.result) {
                    alert("modified success.");
                } else{
                    alert("modified failed.", data.msg);
                };
            });
    }
}

function ManageArticleController($http,$scope,$rootScope,$routeParams){
    $rootScope.global = {
    	title: "管理文章",
    	currentPath: "manage_article"
    }

    var page = $routeParams.page || 1;
    $http.get("/api/admin/article/page/"+page,null).success(function (data) {
        $scope.articles = data.data;
        $scope.has_next = data.nextPage;
        $scope.page = parseInt(page);
    });

}

function ManageProjectController($http,$scope,$rootScope,$routeParams,$sce){
    $rootScope.global = {
    	title: "管理项目",
    	currentPath: "manage_project"
    }

    var page = $routeParams.page || 1;
    $http.get("/api/admin/project/list/"+page,null).success(function (data) {
        // $scope.projects = data.data;
        $scope.total_pages = data.total_pages;
        $scope.has_next = data.has_next;
        $scope.page = parseInt(page);

        $scope.projects = angular.forEach(angular.fromJson(data.data), function (project) {
            project.description_html = $sce.trustAsHtml(project.description);
        });
    });

}

function AddProjectController ($scope,$rootScope) {
    $rootScope.global = {
        title: "项目管理",
        currentPath: "manage_project"
    }
    // $scope.config = simple_ueditor_option;
}

function EditProjectController ($http,$scope,$rootScope,$routeParams) {
    var id = $routeParams.id || 0;
    $scope.id = parseInt(id);

    $rootScope.global = {
        title: "项目管理",
        currentPath: "manage_project"
    }

    $http.get("/api/admin/project/"+id,null).success(function (data) {
        $scope.project = data.data;
    });
}

function ManageOssController($scope,$rootScope){
    $rootScope.global = {
    	title: "OSS管理",
    	currentPath: "manage_oss"
    }
}
