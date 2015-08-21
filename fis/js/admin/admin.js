/**
 * Created by rex on 2015/8/12.
 */

"use strict";
var adminService = angular.module("Admin", ['ngRoute','ng.ueditor']);

function adminRouteConfig($routeProvider, $locationProvider){
    $routeProvider.when("/admin", {
        controller: IndexController,
        templateUrl: "/static/ng/default.html"
    }).when("/admin/new_article", {
        controller: NewArticleController,
        templateUrl: "/static/ng/new_article.html"
    }).when("/admin/manage_article", {
        controller: ManageArticleController,
        templateUrl: "/static/ng/manage_article.html"
    }).when("/admin/manage_project", {
        controller: ManageProjectController,
        templateUrl: "/static/ng/manage_project.html"
    }).when("/admin/manage_oss", {
        controller: ManageOssController,
        templateUrl: "/static/ng/manage_oss.html"
    }).otherwise({
        redirectTo: "/admin"
    });
    $locationProvider.html5Mode(true);
}

adminService.config(adminRouteConfig);

