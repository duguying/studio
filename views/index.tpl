{{{template "inc/header.tpl" .}}}

<div class="article-list">
	{{{range $k,$v := .articles_in_page}}}
		<div class="article" itemscope itemtype="http://schema.org/Article">
			<a class="article-title" title="{{{$v.title}}}" href="/article/{{{$v.uri}}}" itemprop="name">{{{$v.title}}}</a>
			<div class="article-ps">
				Tag {{{$v.keywords|tags|str2html}}} on <span datetime="{{{$v.time}}}" itemprop="datePublished" class="post-time">{{{$v.time}}}</span> by <span title="作者: {{{$v.author}}}" itemprop="author" class="author-name">{{{$v.author}}}</span> view <span title="{{{$v.count}}}次阅读" class="view-count">{{{$v.count}}}</span>
			</div>
			<div class="article-content" itemprop="articleBody">
				{{{str2html $v.content}}}
			</div>
		</div>
		<hr>
	{{{end}}}
	{{{if .prev_page_flag}}}
	<a href="{{{.prev_page}}}" class="page-nav">上一页</a>
	{{{end}}}
	{{{if .next_page_flag}}}
	<a href="{{{.next_page}}}" class="page-nav">下一页</a>
	{{{end}}}
</div>

{{{template "inc/rightbar.tpl" .}}}

{{{template "inc/footer.tpl" .}}}
