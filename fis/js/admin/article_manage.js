/**
 * article_manage.js
 * manage articles
 */
$(document).ready(function (e) {

	rex.gen_article_list = function(data){
		header = '<ul><li class="admin-amb-banner"><ul><li class="admin-col-id admin-amb-cell">ID</li><li class="admin-col-title admin-amb-cell">文章标题</li><li class="admin-col-author admin-amb-cell">作者</li><li class="admin-col-time admin-amb-cell">发布时间</li><li class="admin-col-count admin-amb-cell">阅读量</li><li class="admin-col-operation admin-amb-cell">操作</li></ul></li>';
		footer = '<li class="admin-amb-banner"><ul><li class="admin-col-id admin-amb-cell">ID</li><li class="admin-col-title admin-amb-cell">文章标题</li><li class="admin-col-author admin-amb-cell">作者</li><li class="admin-col-time admin-amb-cell">发布时间</li><li class="admin-col-count admin-amb-cell">阅读量</li><li class="admin-col-operation admin-amb-cell">操作</li></ul></li></ul>';
		nav_prev = '<div style="margin-left: 20px;">';

		if (data.page > 1){
			nav_prev += '<a id="prev-page">上一页</a>';
		}

		nav_next = '</a>';
		if (data.nextPage){
			nav_next += '<a id="next-page">下一页</a>'
		}

		if (!data.result){
			return "";
		}else{
			if (data.data){
				list = data.data;

				html = "";
				for (var key in list){
					// console.log key
					html += '<li class="admin-amb-line"><ul><li class="admin-col-id admin-amb-cell" title="'
							+list[key].id+'">'
							+list[key].id+'</li><li class="admin-col-title admin-amb-cell" title="'
							+list[key].title+'"><a target="_blank" href="/article/'
							+list[key].uri+'">'+list[key].title
							+'</a></li><li class="admin-col-author admin-amb-cell" title="'
							+list[key].author+'">'+list[key].author
							+'</li><li class="admin-col-time admin-amb-cell" title="'
							+list[key].time+'">'+list[key].time
							+'</li><li class="admin-col-count admin-amb-cell" title="'
							+list[key].count+'">'+list[key].count
							+'</li><li class="admin-col-operation admin-amb-cell"><a id="edit" data="'
							+list[key].id+'" style="padding-right: 5px;">辑</a><a id="del" data="'
							+list[key].id+'" style="padding-left: 5px;">删</a></li></ul></li>';
				}

				return header+html+footer+nav_prev+nav_next;
			}else{
				return "";
			}
		}
	}

	rex.load_article_list_page = function(page){
		$("#article-manage-box>.loading").show();

		$.ajax({
			type: "get",
			url: "/admin/article/page/"+page,
			dataType: "json",
			success: function(json){
				if (json.result) {
					$("#article-manage-box>.loading").hide();

					var html = rex.gen_article_list(json);
					$("#article-manage-box>.list").html(html);
					$(".admin-amb-banner li").css({"background-color":"#ccc"});
					$(".admin-amb-line:odd li").css({"background-color":"#ddd"});

					$("#article-manage-box>.list").undelegate().delegate("#edit","click",function(e){
						var id = $(this).attr("data");
						var item = null;
						for (var key in json.data){
							var data = json.data;
							if (id == data[key].id){
								item = data[key];
							}
						}

						rex.edit_article(item);
					}).delegate("#del","click",function(e){
						if (!window.confirm("是否确认要删除文档？")){
							return;
						}

						id = $(this).attr("data");
						console.log("delete "+id);
						$.ajax({
							type: "post",
							url: "/delete",
							data: {"id":id},
							dataType: "json",
							success: function(msg){
								alert(msg.msg);
								rex.load_article_list_page(page);
							}
						});
					});

					$("#prev-page").unbind("click").click(function(e){
						rex.load_article_list_page(page-1);
					});
						
					$("#next-page").unbind("click").click(function(e){
						rex.load_article_list_page(page+1);
					});
					
				} else{
					alert("加载失败 " + json.msg);
				};
			},
			error: function(e){
				alert("请求失败，可能是网络的原因");
			}
		});
	}

	$("#article-manage").click(function(e){
		rex.hide_all();
		$("#article-manage-menu").show();
		$("#article-manage-box").show();

		rex.load_article_list_page(1);
	});

	rex.edit_article = function(item){
		rex.hide_all();
		$("#new-article-menu").show();
		$("#editor-box").show();
		rex.clear_editor();
		rex.clear_artcle_form();

		rex.ajaxdata.url = "/update";

		// load editor content
		$("#article-title").val(item.title);
		$("#article-tags").val(item.keywords);
		rex.ue.setContent(item.content);

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
					data:{ id:item.id, title: title, keywords: tags, content: content },
					dataType: "json",
					success: function(json){
						if (json.result) {
							alert("修改成功。");
							rex.clear_editor();
							rex.clear_artcle_form();
						} else{
							alert("修改失败 " + json.msg);
						};
					},
					error: function(e){
						alert("请求失败，可能是网络的原因，请注意保存文稿！");
					}
				});
			};
		});
	}

});