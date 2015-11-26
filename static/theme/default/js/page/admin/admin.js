/**
 * Created by rex on 2015/8/12.
 */

var adminService = angular.module("Admin", ['ngRoute','ng.ueditor']);

function adminRouteConfig($routeProvider, $locationProvider){
    $routeProvider.when("/admin", {
        controller: IndexController,
        templateUrl: "/static/theme/default/ng/default.html"
    }).when("/admin/new_article", {
        controller: NewArticleController,
        templateUrl: "/static/theme/default/ng/new_article.html"
    }).when("/admin/edit_article", {
        controller: EditArticleController,
        templateUrl: "/static/theme/default/ng/edit_article.html"
    }).when("/admin/edit_article/:id", {
        controller: EditArticleController,
        templateUrl: "/static/theme/default/ng/edit_article.html"
    }).when("/admin/manage_article", {
        controller: ManageArticleController,
        templateUrl: "/static/theme/default/ng/manage_article.html"
    }).when("/admin/manage_article/:page", {
        controller: ManageArticleController,
        templateUrl: "/static/theme/default/ng/manage_article.html"
    }).when("/admin/manage_project", {
        controller: ManageProjectController,
        templateUrl: "/static/theme/default/ng/manage_project.html"
    }).when("/admin/manage_project/:page", {
        controller: ManageProjectController,
        templateUrl: "/static/theme/default/ng/manage_project.html"
    }).when("/admin/new_project", {
        controller: AddProjectController,
        templateUrl: "/static/theme/default/ng/new_project.html"
    }).when("/admin/edit_project/:id", {
        controller: EditProjectController,
        templateUrl: "/static/theme/default/ng/edit_project.html"
    }).when("/admin/manage_oss", {
        controller: ManageOssController,
        templateUrl: "/static/theme/default/ng/manage_oss.html"
    }).otherwise({
        redirectTo: "/admin"
    });
    $locationProvider.html5Mode(true);
}

adminService.config(adminRouteConfig);
