/**
 * scripts for admin pannel
 * @author rex lee
 * @email duguying2008@gmai;
 */

//按钮的操作
function insertHtml() {
    var value = prompt('插入html代码', '');
    um.execCommand('insertHtml', value)
}

function isFocus(){
    alert(um.isFocus())
}

function doBlur(){
    um.blur()
}

function createEditor() {
    enableBtn();
    um = UM.getEditor('myEditor');
}

function getAllHtml() {
    alert(UM.getEditor('myEditor').getAllHtml())
}

function getContent() {
    var arr = [];
    arr.push("使用editor.getContent()方法可以获得编辑器的内容");
    arr.push("内容为：");
    arr.push(UM.getEditor('myEditor').getContent());
    alert(arr.join("\n"));
}

function getPlainTxt() {
    var arr = [];
    arr.push("使用editor.getPlainTxt()方法可以获得编辑器的带格式的纯文本内容");
    arr.push("内容为：");
    arr.push(UM.getEditor('myEditor').getPlainTxt());
    alert(arr.join('\n'))
}

function setContent(isAppendTo) {
    var arr = [];
    arr.push("使用editor.setContent('欢迎使用umeditor')方法可以设置编辑器的内容");
    UM.getEditor('myEditor').setContent('欢迎使用umeditor', isAppendTo);
    alert(arr.join("\n"));
}

function setDisabled() {
    UM.getEditor('myEditor').setDisabled('fullscreen');
    disableBtn("enable");
}

function setEnabled() {
    UM.getEditor('myEditor').setEnabled();
    enableBtn();
}

function getText() {
    //当你点击按钮时编辑区域已经失去了焦点，如果直接用getText将不会得到内容，所以要在选回来，然后取得内容
    var range = UM.getEditor('myEditor').selection.getRange();
    range.select();
    var txt = UM.getEditor('myEditor').selection.getText();
    alert(txt)
}

function getContentTxt() {
    var arr = [];
    arr.push("使用editor.getContentTxt()方法可以获得编辑器的纯文本内容");
    arr.push("编辑器的纯文本内容为：");
    arr.push(UM.getEditor('myEditor').getContentTxt());
    alert(arr.join("\n"));
}

function hasContent() {
    var arr = [];
    arr.push("使用editor.hasContents()方法判断编辑器里是否有内容");
    arr.push("判断结果为：");
    arr.push(UM.getEditor('myEditor').hasContents());
    alert(arr.join("\n"));
}

function setFocus() {
    UM.getEditor('myEditor').focus();
}

function deleteEditor() {
    disableBtn();
    UM.getEditor('myEditor').destroy();
}

function disableBtn(str) {
    var div = document.getElementById('btns');
    var btns = domUtils.getElementsByTagName(div, "button");
    for (var i = 0, btn; btn = btns[i++];) {
        if (btn.id == str) {
            domUtils.removeAttributes(btn, ["disabled"]);
        } else {
            btn.setAttribute("disabled", "true");
        }
    }
}

function enableBtn() {
    var div = document.getElementById('btns');
    var btns = domUtils.getElementsByTagName(div, "button");
    for (var i = 0, btn; btn = btns[i++];) {
        domUtils.removeAttributes(btn, ["disabled"]);
    }
}

// 实例化编辑器
function initUeditor() {
    var um = UM.getEditor('myEditor');
    um.addListener('blur',function(){
        $('#focush2').html('编辑器失去焦点了')
    });
    um.addListener('focus',function(){
        $('#focush2').html('')
    });
    
}

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
		var menu_bar_html = '<label for="article-title" style="margin-left: 10px;color:white;">文章标题</label><input type="text" name="title" id="article-title" style="margin-left: 10px;margin-top: 7px;width: 250px;"><label for="article-tags" style="color: white;margin-left: 10px;margin-right: 10px;">关键词</label><input type="text" name="tags" placeholder="逗号,分隔" id="article-tags" class="article-tags"><button style="margin-left: 10px;">发布</button>';
		var editor_html = '<script type="text/plain" id="myEditor" style="width:100%;height:464px;"></script>';

		if (!new_ariticle_clicked) {
			new_article_menu.html(menu_bar_html);
			new_ariticle_box.html(editor_html);
			// 初始化编辑器
			initUeditor();
			new_ariticle_clicked = true;
		} else{
			new_ariticle_box.show();
			new_ariticle_box.show();
		};
		
		
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
