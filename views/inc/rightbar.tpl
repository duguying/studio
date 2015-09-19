			<div class="box-right">

				<div class="box">
					<a class="box-title">最热文章</a>
					<ul>
						{{{range $k,$v := .hottest}}}
							<li>
								<a href="/article/{{$v.title}}">
									<span>{{{$v.title}}}</span>
								</a>
								<span class="view-count">{{{$v.count}}}</span>
							</li>
						{{{end}}}
					</ul>
				</div>
				<div class="box">
					<a class="box-title">按月归档</a>
					<ul>
						{{{range $k,$v := .count_by_month}}}
							<li>
								<a href="/archive/{{{$v.year}}}/{{{$v.month}}}/1">
									<span>{{{$v.date}}} [{{{$v.number}}}]</span>
								</a>
							</li>
						{{{end}}}
					</ul>
				</div>

			</div>
