			<div class="footer">
				<div class="links">
					<ul class="colume">
						<li>link1</li>
						<li>link2</li>
						<li>link3</li>
						<li>link4</li>
						<li>link5</li>
					</ul>

					<ul class="colume">
						<li>link1</li>
						<li>link2</li>
						<li>link3</li>
						<li>link4</li>
						<li>link5</li>
					</ul>

					<ul class="colume">
						<li>link1</li>
						<li>link2</li>
						<li>link3</li>
						<li>link4</li>
						<li>link5</li>
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
				</script>

			</div>