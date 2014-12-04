/**
 * new_article.js
 * add article
 */
$(document).ready(function (e) {
	
	$("#new-article").unbind("click").click(function(e){
		rex.hide_all();
		$("#new-article-menu").show();
		$("#editor-box").show();
		rex.clear_editor();
		rex.clear_artcle_form();

		rex.ajaxdata.url = "/add";

		$("#submit").unbind("click").click(function (e) {
			var title = $("#article-title").val();
			var tags = $("#article-tags").val();
			tags = tags.replace(/，/, ',')
			var content = rex.ue.getContent();

			if (title.length <= 0) {
				alert("标题不能为空！");
				return;
			} else{
				$.ajax({
					type: "post",
					url: rex.ajaxdata.url,
					data:{
						"title": title,
						"keywords": tags,
						"content": content
					},
					dataType: "json",
					success: function(json){
						if (json.result) {
							alert("添加成功。");
							rex.clear_editor();
							rex.clear_artcle_form();
						} else{
							alert("添加失败 " + json.msg);
						};
					},
					error: function(e){
						alert("请求失败，可能是网络的原因，请注意保存文稿！");
					}
				});
			};
		});

	});
	
});