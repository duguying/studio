"use strict";

adminService.directive("active", function () {
	return {
		restrict: "A",
		link: function (scope, element) {
			console.log(element);
		}
	};
});
