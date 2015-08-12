module Admin {
    export interface Scope {
        navs: string;
    }

    export interface Http {
        get(url:string,params:any);
    }

    export class TestController {
        constructor ($scope: Scope, $http: Http) {
            $http.get("/admin/api/navlist",null).success(function (data) {
                $scope.navs = data;
            });

        }
    }
}