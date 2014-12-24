<link rel="stylesheet" href="//cdnjs.cloudflare.com/ajax/libs/fancybox/2.1.5/jquery.fancybox.min.css">
<link href="/static/img/favicon.ico" mce_href="/static/img/favicon.ico" rel="bookmark" type="image/x-icon" /> 
<link href="/static/img/favicon.ico" mce_href="/static/img/favicon.ico" rel="icon" type="image/x-icon" /> 
<link href="/static/img/favicon.ico" mce_href="/static/img/favicon.ico" rel="shortcut icon" type="image/x-icon" /> 
        <div class="header">
  				<div class="icons">
  					<a href="http://my.oschina.net/duguying" target="_black"><span title="follow me on oschina" class="icon-osc icon"></span></a>
            <a href="https://github.com/duguying" target="_black"><span title="follow me on Github" class="icon-github icon"></span></a>
  					<a href="http://weibo.com/duguying2008" target="_black"><span title="find me on Weibo" class="icon-weibo icon"></span></a>
  					<a href="http://gplus.to/duguying" target="_black"><span title="find me on g+" class="icon-gplus icon"></span></a>
  					<a href="https://twitter.com/duguying" target="_black"><span title="find me on Twitter" class="icon-twitter icon"></span></a>
  				</div>
  				<ul class="menu">
  					<li id="about">
              <a>关于</a>
              <div class="drop-menu">
                <span class="droplist-array-down">◆</span>
                <ul class="drop-list">
                  <li>Blog</li>
                  <li><a href="/resume/statistics">Statistics</a></li>
                </ul>
              </div>
            </li>
  					<li>
              <a href="/project">项目</a>
            </li>
  					<li><a href="/list">列表</a></li>
  					<li><a href="/">博文</a></li>
  				</ul>
  				<div class="banner">
  					<a href="/"><span class="title">独孤影</span></a>
  				</div>
  				<div class="gap">
            {{if eq .is_admin "admin"}}
  					<a href="/admin" title="管理页面">
              <img class="gravatar" src="/logo" />
            </a>
            {{else}}
            <img class="gravatar" src="/logo" />
            {{end}}
  				</div>
  			</div>