  			{{{template "inc/header.tpl" .}}}

			<div class="projects">
				<h1><a href="/project">项目</a></h1>
				<div class="intro">
					此处包含了本人接触编程以来所开发或参与开发的项目，这些项目多为开源项目，欢迎参阅。
				</div>
				{{{range $k,$v := .projects_in_page}}}
					<div class="project-box">
						<img src="{{{$v.icon_url}}}" alt="">
						<div class="info">
							<span class="name">{{{$v.name}}}</span><hr>
							<span class="author">作者: {{{$v.author}}}</span>
							<span class="time">时间: {{{date $v.time}}}</span>
							<div>{{{str2html $v.description}}}</div>
						</div>
					</div>
				{{{end}}}
				<div class="nav-project-list">
					{{{if .prev_page_flag}}}
					<a href="{{.prev_page}}" class="page-nav">上一页</a>
					{{{end}}}
					{{{if .next_page_flag}}}
					<a href="{{.next_page}}" class="page-nav">下一页</a>
					{{{end}}}
				</div>
			</div>

			{{{template "inc/footer.tpl" .}}}
