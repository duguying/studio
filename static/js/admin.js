/**
 * scripts for admin pannel
 * @author rex lee
 * @email duguying2008@gmai;
 */
var ue;
$(document).ready(function (e) {
    
	
	var new_article_menu = $("#new-article-menu");
	var article_manage_menu = $("#article-manage-menu");
	var attach_manage_menu = $("#attach-manage-menu");
	var oss_manage_menu = $("#oss-manage-menu");

	var new_ariticle_box = $("#new-article-box")
	var new_ariticle_clicked = false;
	var article_manage_box = $("article-manage-box");
	var article_manage_clicked = false;
	var attach_manage = $("#attach-manage-box");
	var attach_manage_clicked = false;
	var oss_manage = $("#oss-manage-box");
	var oss_manage_clicked = false;

	$("#new_ariticle").click(function (e) {
		var menu_bar_html = '<label for="article-title" style="margin-left: 10px;color:white;">文章标题</label><input type="text" name="title" id="article-title" style="margin-left: 10px;margin-top: 7px;width: 250px;"><label for="article-tags" style="color: white;margin-left: 10px;margin-right: 10px;">关键词</label><input type="text" name="tags" placeholder="逗号,分隔" id="article-tags" class="article-tags"><button id="submit" style="margin-left: 10px;background-color: white;border-radius: 4px;">发布</button>';
		var editor_html = '<textarea class="m-input d-input" name="content" id="myEditor" style="width:100%;height:430px;"></textarea>';

		if (!new_ariticle_clicked) {
			new_article_menu.html(menu_bar_html);
			new_ariticle_box.html(editor_html);
			// 初始化编辑器
			ue = UE.getEditor('myEditor');
			new_ariticle_clicked = true;
		} else{
			return;
		};

        // 编辑器就绪后
        ue.ready(function() {
            
            $("#submit").click(function (e) {
                var current_content = ue.getContent();;
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

        });
		
		
	});

	$("#article-manage").click(function (e) {
		;
	});

	$("#attach-manage").click(function (e) {
		;
	});

	$("#oss-manage").click(function (e) {
		;
	});
});
