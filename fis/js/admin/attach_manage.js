/**
 * attach_manage.js
 * manage attachment
 */
$(document).ready(function (e) {
	$("#attach-manage").unbind("click").click(function(e){
		rex.hide_all();
		$("#attach-manage-menu").show();
		$("#attach-manage-box").show();
	});
});