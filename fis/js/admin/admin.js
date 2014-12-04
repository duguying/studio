/**
 * initialize admin models
 */
var rex = {};

$(document).ready(function (e) {
	rex.hide_top_menu = function() {
		$(".admin-top-menu>div").hide();
	}

	rex.hide_right_pannel = function () {
		$(".admin-main-pannel>div").hide();
	}

	rex.hide_all = function () {
		rex.hide_top_menu()
		rex.hide_right_pannel()
	}

	rex.clear_editor = function () {
		if (rex.ue) {
			rex.ue.setContent('');
		} else{
			console.log("ueditor have not initialized")
		};
	}

	rex.clear_artcle_form = function () {
		$("#article-title").val("");
		$("#article-tags").val("");
	}

	rex.hide_all();
	rex.ue = UE.getEditor('myEditor');

	$("#default-box").show();

	rex.ajaxdata = {
		url:"",
		method:"get",
		data:"",
	};
});