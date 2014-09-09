/**
 * scripts for admin pannel
 * @author rex lee
 * @email duguying2008@gmai;
 */

$(document).ready(function (e) {
	var ue;
    
	var new_article_menu = $("#new-article-menu");
	var article_manage_menu = $("#article-manage-menu");
	var attach_manage_menu = $("#attach-manage-menu");
	var oss_manage_menu = $("#oss-manage-menu");

	var new_ariticle_box = $("#new-article-box")
	var article_manage_box = $("#article-manage-box");
	var attach_manage_box = $("#attach-manage-box");
	var oss_manage_box = $("#oss-manage-box");

	function show_frame (item) {
		var box1 = new_ariticle_box;
		var banner1 = new_article_menu;
		var box2 = article_manage_box;
		var banner2 = article_manage_menu;
		var box3 = attach_manage_box;
		var banner3 = attach_manage_menu;
		var box4 = oss_manage_box;
		var banner4 = oss_manage_menu;

		if ("box1" == item) {
			box1.show();banner1.show();
			box2.hide();banner2.hide();
			box3.hide();banner3.hide();
			box4.hide();banner4.hide();
		} else if("box2" == item){
			box1.hide();banner1.hide();
			box2.show();banner2.show();
			box3.hide();banner3.hide();
			box4.hide();banner4.hide();
		} else if("box3" == item){
			box1.hide();banner1.hide();
			box2.hide();banner2.hide();
			box3.show();banner3.show();
			box4.hide();banner4.hide();
		}else if("box4" == item){
			box1.hide();banner1.hide();
			box2.hide();banner2.hide();
			box3.hide();banner3.hide();
			box4.show();banner4.show();
		}else{
			;
		};
	};

	(function init () {
		var menu_bar_html = '<label for="article-title" style="margin-left: 10px;color:white;">文章标题</label>\
<input type="text" name="title" id="article-title" style="margin-left: 10px;margin-top: 7px;width: 250px;">\
<label for="article-tags" style="color: white;margin-left: 10px;margin-right: 10px;">关键词</label>\
<input type="text" name="tags" placeholder="逗号,分隔" id="article-tags" class="article-tags">\
<button id="submit" style="margin-left: 10px;background-color: white;border-radius: 4px;">发布</button>';
		var editor_html = '<textarea class="m-input d-input" name="content" id="myEditor" style="width:100%;height:430px;"></textarea>';
		new_article_menu.html(menu_bar_html);
		new_ariticle_box.html(editor_html);
		// 初始化编辑器
		ue = UE.getEditor('myEditor');
		new_article_menu.hide();
		new_ariticle_box.hide();

	})();

	function edit_submit (item) {
		if (!item) {
			// 新建文章
			$("#submit").unbind("click");
			$("#submit").click(function (e) {
	            var current_content = ue.getContent();
	            var current_title = $("#article-title").val();
	            var tags = $("#article-tags").val();

	            $.ajax({
	                type: "post",
	                url: "/add",
	                data: { title: current_title, keywords: tags, content: current_content },
	                dataType: "json",
	                success: function(msg){
	                    console.log(msg);
	                    if (msg.result) {
	                        ue.setContent('');
	                        $("#article-title").val('');
	                        $("#article-tags").val('');
	                        alert("发布成功-"+msg.msg);
	                    } else{
	                        alert("发布失败-"+msg.msg);
	                    };
	                }
	            });
	        });
		} else{
			// 修改文章
			show_frame("box1");
			ue.setContent(item.content);
	        $("#article-title").val(item.title);
	        $("#article-tags").val(item.keywords);

	        $("#submit").unbind("click");
			$("#submit").click(function (e) {
				var current_content = ue.getContent();;
	            var current_title = $("#article-title").val();
	            var tags = $("#article-tags").val();
	            $.ajax({
	                type: "post",
	                url: "/update",
	                data: { id:item.id, title: current_title, keywords: tags, content: current_content },
	                dataType: "json",
	                success: function(msg){
	                	alert(msg.msg);
	                },
	            });
			});
		};

		
	}

	function gen_list (data) {
		var ambl = $(".amb-content");
		var header = '<ul><li class="admin-amb-banner"><ul>\
<li class="admin-col-id admin-amb-cell">ID</li>\
<li class="admin-col-title admin-amb-cell">文章标题</li>\
<li class="admin-col-author admin-amb-cell">作者</li>\
<li class="admin-col-time admin-amb-cell">发布时间</li>\
<li class="admin-col-count admin-amb-cell">阅读量</li>\
<li class="admin-col-operation admin-amb-cell">操作</li>\
</ul></li>';
		var footer = '<li class="admin-amb-banner"><ul>\
<li class="admin-col-id admin-amb-cell">ID</li>\
<li class="admin-col-title admin-amb-cell">文章标题</li>\
<li class="admin-col-author admin-amb-cell">作者</li>\
<li class="admin-col-time admin-amb-cell">发布时间</li>\
<li class="admin-col-count admin-amb-cell">阅读量</li>\
<li class="admin-col-operation admin-amb-cell">操作</li>\
</ul></li></ul>';
		var nav_prev='<div style="margin-left: 20px;">';
		if (data.page>1) {
			nav_prev += '<a id="prev-page">上一页</a>'
		};
		var nav_next='</a>';
		if (data.nextPage) {
			nav_next += '<a id="next-page">下一页</a>'
		};

		if(!data.result){
			return "";
		}else{
			if (data.data) {
				var list = data.data;
				var html = "";
				for(key in list){
					html+='<li class="admin-amb-line"><ul><li class="admin-col-id admin-amb-cell" title="'+list[key].id+'">'+list[key].id+'</li>\
<li class="admin-col-title admin-amb-cell" title="'+list[key].title+'">'+list[key].title+'</li>\
<li class="admin-col-author admin-amb-cell" title="'+list[key].author+'">'+list[key].author+'</li>\
<li class="admin-col-time admin-amb-cell" title="'+list[key].time+'">'+list[key].time+'</li>\
<li class="admin-col-count admin-amb-cell" title="'+list[key].count+'">'+list[key].count+'</li>\
<li class="admin-col-operation admin-amb-cell"><a id="edit" data="'+list[key].id+'" style="padding-right: 5px;">辑</a>\
<a id="del" data="'+list[key].id+'" style="padding-left: 5px;">删</a></li></ul></li>';
				};
				return header+html+footer+nav_prev+nav_next;
			} else{
				return "";
			};
		}
	}

	$("#new-ariticle").click(function (e) {
		show_frame("box1");
        edit_submit();
		
	});

	function get_page (page) {
		ueditor_request_url = "/admin/article/page/"+page;
		$.ajax({
			type: "get",
			url: ueditor_request_url,
			dataType: "json",
			success: function(msg){
				var html = gen_list(msg);
				$("#article-manage-box").html(html);
				$(".admin-amb-banner li").css({"background-color":"#ccc"});
				$(".admin-amb-line:odd li").css({"background-color":"#ddd"});

				$("#article-manage-box").undelegate();
				$("#article-manage-box").delegate("#edit","click",function(){
					var id = $(this).attr("data");
					var item = null;
					for(var key in msg.data){
						var data = msg.data;
						if(id == data[key].id){
							item = data[key];
						};
					}
					// ueditor_request_url = "/update";
					edit_submit(item);
					
				}).delegate("#del","click",function(){
					var id = $(this).attr("data");
					console.log("delete "+id);
					$.ajax({
						type: "post",
						url: "/delete",
						data: {"id":id},
						dataType: "json",
						success: function(msg){
							alert(msg.msg);
							get_page(page);
						},
					});
				});

				$("#prev-page").click(function (e) {
					get_page(page-1);
				});
				$("#next-page").click(function (e) {
					get_page(page+1);
				});
			},
		});
	}

	$("#article-manage").click(function (e) {
		show_frame("box2");
		get_page(1)
	});

	$("#attach-manage").click(function (e) {
		show_frame("box3");
	});

	$("#oss-manage").click(function (e) {
		show_frame("box4");
	});
});
