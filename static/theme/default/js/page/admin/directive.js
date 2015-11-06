adminService.directive("deleteArticle", function ($document,$http) {
	return{
		restrict:'A',
		require: 'ngModel',
		link:function(scope, element, attrs,ngModel){
			element.bind("click",function(){
				var id = ngModel.$modelValue.id;

				if (!window.confirm("Sure to Delete ["+ngModel.$modelValue.title+"]?")) {
            		return;
            	}

				$http.post("/api/admin/delete", {
		                params: {"id":parseInt(id)}
		            }).success(function(data){
		                if (data.result) {
		                    console.log("add success.");
		                    for (var i = 0; i < scope.articles.length; i++) {
								if(scope.articles[i].id == parseInt(id)){
									scope.articles.splice(i,1);
								}
							};
							// delete ngModel.$modelValue;
		                } else{
		                    console.log("add failed.", data.msg);
		                };
		            });
			})
		}
	}
});
