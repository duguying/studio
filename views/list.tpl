  		{{{template "inc/header.tpl" .}}}

		<div class="lp-article-list">
			<ul>
				{{{range $k,$v := .list}}}
				<li class="a-l-item">
					<ul>
						<li class="time">{{{date_cn $v.time}}}</li>
						<li class="title"><a href="/article/{{{$v.uri}}}">{{{$v.title}}}</a></li>
					</ul>
				</li>
				{{{end}}}
			</ul>
		</div>
		<div class="lp-nav">
			{{{if .prev_page_flag}}}
			<a href="/list/{{{.prev_page}}}" class="page-nav">上一页</a>
			{{{end}}}
			{{{if .next_page_flag}}}
			<a href="/list/{{{.next_page}}}" class="page-nav">下一页</a>
			{{{end}}}
		</div>

  		{{{template "inc/footer.tpl" .}}}
