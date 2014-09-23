$ document
.ready (e) ->
	ue = null

	new_article_menu = $("#new-article-menu")
	article_manage_menu = $("#article-manage-menu")
	attach_manage_menu = $("#attach-manage-menu")
	oss_manage_menu = $("#oss-manage-menu")

	new_ariticle_box = $("#new-article-box")
	article_manage_box = $("#article-manage-box");
	attach_manage_box = $("#attach-manage-box");
	oss_manage_box = $("#oss-manage-box");

	show_frame = (item) ->
		box1 = new_ariticle_box;
		banner1 = new_article_menu;
		box2 = article_manage_box;
		banner2 = article_manage_menu;
		box3 = attach_manage_box;
		banner3 = attach_manage_menu;
		box4 = oss_manage_box;
		banner4 = oss_manage_menu;

		if "box1" is item 
			box1.show();banner1.show();
			box2.hide();banner2.hide();
			box3.hide();banner3.hide();
			box4.hide();banner4.hide();
		else if "box2" is item
			box1.hide();banner1.hide();
			box2.show();banner2.show();
			box3.hide();banner3.hide();
			box4.hide();banner4.hide();
		else if "box3" is item
			box1.hide();banner1.hide();
			box2.hide();banner2.hide();
			box3.show();banner3.show();
			box4.hide();banner4.hide();
		else if "box4" is item
			box1.hide();banner1.hide();
			box2.hide();banner2.hide();
			box3.hide();banner3.hide();
			box4.show();banner4.show();
		else
			
	get_local = () ->
		index = window
				.location
				.href
				.replace(/\./g,'_')
				.replace(/:/g,'_')
				.replace(/\//g,'_')
				.replace(/___/g,'_')+'myEditor-drafts-data'
		console.log(index)
		eval('['+localStorage.getItem("ueditor_preference")+']')[0][index]

	edit_submit = (item) ->
		if not item
			# 新建文章
			$("#submit").unbind("click");
			$("#submit").click(e) ->
				current_content = ue.getContent();
				current_title = $("#article-title").val();
				tags = $("#article-tags").val();

				$.ajax
					type: "post"
					url: "/add"
					data: 
						title: current_title
						keywords: tags
						content: current_content
					dataType: "json"
					success: (msg) ->
						console.log(msg);
						if msg.result
							ue.setContent('');
							$("#article-title").val('');
							$("#article-tags").val('');
							alert("发布成功-"+msg.msg);
						else
							alert("发布失败-"+msg.msg);
		else
			# 修改文章
			show_frame("box1");
			ue.setContent(item.content);
			$("#article-title").val(item.title);
			$("#article-tags").val(item.keywords);

			$("#submit").unbind("click");
			$("#submit").click (e) ->
				current_content = ue.getContent();;
				current_title = $("#article-title").val();
				tags = $("#article-tags").val();
				$.ajax
					type: "post",
					url: "/update",
					data: { id:item.id, title: current_title, keywords: tags, content: current_content },
					dataType: "json",
					success: (msg) ->
						alert(msg.msg);

	gen_list = (data) ->
		ambl = $(".amb-content");
		header = '''
<ul><li class="admin-amb-banner"><ul>
<li class="admin-col-id admin-amb-cell">ID</li>
<li class="admin-col-title admin-amb-cell">文章标题</li>
<li class="admin-col-author admin-amb-cell">作者</li>
<li class="admin-col-time admin-amb-cell">发布时间</li>
<li class="admin-col-count admin-amb-cell">阅读量</li>
<li class="admin-col-operation admin-amb-cell">操作</li>
</ul></li>'''
		footer = '''<li class="admin-amb-banner"><ul>
<li class="admin-col-id admin-amb-cell">ID</li>
<li class="admin-col-title admin-amb-cell">文章标题</li>
<li class="admin-col-author admin-amb-cell">作者</li>
<li class="admin-col-time admin-amb-cell">发布时间</li>
<li class="admin-col-count admin-amb-cell">阅读量</li>
<li class="admin-col-operation admin-amb-cell">操作</li>
</ul></li></ul>'''
		nav_prev = '''<div style="margin-left: 20px;">'''

		if data.page > 1
			nav_prev += '''<a id="prev-page">上一页</a>'''
		
		nav_next = '''</a>''';
		if data.nextPage
			nav_next += '''<a id="next-page">下一页</a>'''
		
		if !data.result
			return "";
		else
			if data.data
				list = data.data;

				html = "";
				for key of list
					# console.log key
					html += '<li class="admin-amb-line"><ul><li class="admin-col-id admin-amb-cell" title="'+list[key].id+'">'+list[key].id+'</li>\
<li class="admin-col-title admin-amb-cell" title="'+list[key].title+'">'+list[key].title+'</li>\
<li class="admin-col-author admin-amb-cell" title="'+list[key].author+'">'+list[key].author+'</li>\
<li class="admin-col-time admin-amb-cell" title="'+list[key].time+'">'+list[key].time+'</li>\
<li class="admin-col-count admin-amb-cell" title="'+list[key].count+'">'+list[key].count+'</li>\
<li class="admin-col-operation admin-amb-cell"><a id="edit" data="'+list[key].id+'" style="padding-right: 5px;">辑</a>\
<a id="del" data="'+list[key].id+'" style="padding-left: 5px;">删</a></li></ul></li>';
				
				return header+html+footer+nav_prev+nav_next;
			else
				return "";
	
	get_page = (page) ->
		ueditor_request_url = "/admin/article/page/"+page;
		$.ajax
			type: "get",
			url: ueditor_request_url,
			dataType: "json",
			success: (msg) ->
				html = gen_list(msg);
				$("#article-manage-box").html(html);
				$(".admin-amb-banner li").css({"background-color":"#ccc"});
				$(".admin-amb-line:odd li").css({"background-color":"#ddd"});
				$("#article-manage-box").undelegate();
				$("#article-manage-box")
				.delegate "#edit","click",() ->
					id = $(this).attr("data");
					item = null;
					for key of msg.data
						data = msg.data;
						if id == data[key].id
							item = data[key];
					edit_submit(item);
				.delegate "#del","click",() ->
					if !window.confirm("是否确认要删除文档？")
						return;
					id = $(this).attr("data");
					console.log("delete "+id);
					$.ajax
						type: "post",
						url: "/delete",
						data: {"id":id},
						dataType: "json",
						success: (msg) ->
							alert(msg.msg);
							get_page(page);
				$("#prev-page").click (e) ->
					get_page(page-1);
				$("#next-page").click (e) ->
					get_page(page+1);


	$("#new-ariticle").click (e) ->
		show_frame("box1");
		edit_submit();
		data = get_local();
		if data&&data!=ue.getContent()
			if window.confirm("是否加载上次未保存的内容？")
				ue.setContent(data);
			else
				ue.setContent("");

	$("#article-manage").click (e) ->
		show_frame("box2");
		get_page(1)
	
	$("#attach-manage").click (e) ->
		show_frame("box3");
	
	$("#oss-manage").click (e) ->
		show_frame("box4");
	
	
	menu_bar_html = '<label for="article-title" style="margin-left: 10px;color:white;">文章标题</label>\
<input type="text" name="title" id="article-title" style="margin-left: 10px;margin-top: 7px;width: 250px;">\
<label for="article-tags" style="color: white;margin-left: 10px;margin-right: 10px;">关键词</label>\
<input type="text" name="tags" placeholder="逗号,分隔" id="article-tags" class="article-tags">\
<button id="submit" style="margin-left: 10px;background-color: white;border-radius: 4px;">发布</button>';
	editor_html = '<textarea class="m-input d-input" name="content" id="myEditor" style="width:100%;height:430px;"></textarea>';
	new_article_menu.html(menu_bar_html);
	new_ariticle_box.html(editor_html);
	# 初始化编辑器
	ue = UE.getEditor('myEditor');
	new_article_menu.hide();
	new_ariticle_box.hide();

	