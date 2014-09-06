			<div class="footer">
				<div class="links">
					<ul class="colume">
						<li class="blogroll">Blogroll</li>
						<li><a href="https://www.byvoid.com" target="_blank">byvoid</a></li>
						<li><a href="http://www.alloyteam.com/" target="_blank">alloy</a></li>
						<li></li>
						<li></li>
						<li></li>
					</ul>

					<ul class="colume">
						<li class="blogroll">Blogroll</li>
						<li><a href="http://www.lyblog.net/">刘洋</a></li>
						<li><a href="http://my.oschina.net/xsilen">xsilen</a></li>
						<li></li>
						<li></li>
						<li></li>
					</ul>

					<ul class="colume">
						<li></li>
						<li></li>
						<li></li>
						<li></li>
						<li></li>
					</ul>

					<ul class="colume">
						<li></li>
						<li></li>
						<li></li>
						<li></li>
						<li></li>
					</ul>
				</div>
				<div class="copyright">©2014 the theme designed by Rex Lee inspired by <a href="https://www.byvoid.com/">byvoid</a>, the program written by Rex Lee with Golang base on <a href="http://beego.me/">Beego</a> framework.</div>

				<script src="/static/syntaxhighlighter/scripts/shCore.js"></script>
				<script src="/static/syntaxhighlighter/scripts/shAutoloader.js"></script>
				<script>
				function path()
				{
				  var args = arguments,
				      result = []
				      ;
				  for(var i = 0; i < args.length; i++)
				      result.push(args[i].replace('@', '/static/syntaxhighlighter/scripts/'));
				  return result
				};
				SyntaxHighlighter.autoloader.apply(null, path(
				  'applescript            @shBrushAppleScript.js',
				  'actionscript3 as3      @shBrushAS3.js',
				  'bash shell             @shBrushBash.js',
				  'coldfusion cf          @shBrushColdFusion.js',
				  'cpp c                  @shBrushCpp.js',
				  'c# c-sharp csharp      @shBrushCSharp.js',
				  'css                    @shBrushCss.js',
				  'delphi pascal          @shBrushDelphi.js',
				  'diff patch pas         @shBrushDiff.js',
				  'erl erlang             @shBrushErlang.js',
				  'groovy                 @shBrushGroovy.js',
				  'java                   @shBrushJava.js',
				  'jfx javafx             @shBrushJavaFX.js',
				  'js jscript javascript  @shBrushJScript.js',
				  'perl pl                @shBrushPerl.js',
				  'php                    @shBrushPhp.js',
				  'text plain             @shBrushPlain.js',
				  'py python              @shBrushPython.js',
				  'ruby rails ror rb      @shBrushRuby.js',
				  'sass scss              @shBrushSass.js',
				  'scala                  @shBrushScala.js',
				  'sql                    @shBrushSql.js',
				  'vb vbnet               @shBrushVb.js',
				  'xml xhtml xslt html    @shBrushXml.js'
				));
				SyntaxHighlighter.defaults.toolbar = false;
				SyntaxHighlighter.defaults.title = '';
				SyntaxHighlighter.all();
				
				/**
				 * back to top config
				 */
				
				var mv_dynamic_to_top = {"text":"To Top","version":"0","min":"300","speed":"300","easing":"easeInOutExpo","margin":"20"};

				/*
				 * Dynamic To Top Plugin
				 * http://www.mattvarone.com
				 *
				 * By Matt Varone
				 * @sksmatt
				 *
				 */
				var mv_dynamic_to_top;(function($,mv_dynamic_to_top){jQuery.fn.DynamicToTop=function(options){var defaults={text:mv_dynamic_to_top.text,min:parseInt(mv_dynamic_to_top.min,10),fade_in:600,fade_out:400,speed:parseInt(mv_dynamic_to_top.speed,10),easing:mv_dynamic_to_top.easing,version:mv_dynamic_to_top.version,id:'dynamic-to-top'},settings=$.extend(defaults,options);if(settings.version===""||settings.version==='0'){settings.text='<span>&nbsp;</span>';}
				if(!$.isFunction(settings.easing)){settings.easing='linear';}
				var $toTop=$('<a href=\"#\" id=\"'+settings.id+'\"></a>').html(settings.text);$toTop.hide().appendTo('body').click(function(){$('html, body').stop().animate({scrollTop:0},settings.speed,settings.easing);return false;});$(window).scroll(function(){var sd=jQuery(window).scrollTop();if(typeof document.body.style.maxHeight==="undefined"){$toTop.css({'position':'absolute','top':sd+$(window).height()-mv_dynamic_to_top.margin});}
				if(sd>settings.min){$toTop.fadeIn(settings.fade_in);}else{$toTop.fadeOut(settings.fade_out);}});};$('body').DynamicToTop();})(jQuery,mv_dynamic_to_top);

				/**
				 * Image load error, load default image
				 */
				$(function (e) {
					$("img").css('background-image','url(/static/img/loadfailed.png)')
				})
				</script>
			</div>