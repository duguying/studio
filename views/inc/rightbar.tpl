			<div class="box-right">

				<div class="box">
					<a class="box-title">最热文章</a>
					<ul>
						{{range $k,$v := .hottest}}
							<li><a href="/article/{{$v.title}}"><span>{{$v.title}} - {{$v.count}}</span></a></li>
						{{end}}
					</ul>
				</div>
				<div class="box">
					<a class="box-title">按月归档</a>
					<ul>
						{{range $k,$v := .count_by_month}}
							<li><a href=""><span>{{$v.month}} [{{$v.number}}]</span></a></li>
						{{end}}
					</ul>
				</div>
				
			</div>
