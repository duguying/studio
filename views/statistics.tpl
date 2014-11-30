<!DOCTYPE html>
<html lang="en">
<head>
	<meta charset="UTF-8">
	<title>数据统计</title>
	<meta content="李俊,数据,Statistics" name="keywords" />
	<meta content="李俊的数据" name="description" />
	<script src="//ajax.googleapis.com/ajax/libs/jquery/1.11.1/jquery.min.js"></script>
	<script src="//www.ichartjs.com/ichart.latest.min.js"></script>
	<link href="/static/img/favicon.ico" mce_href="/static/img/favicon.ico" rel="bookmark" type="image/x-icon" /> 
	<link href="/static/img/favicon.ico" mce_href="/static/img/favicon.ico" rel="icon" type="image/x-icon" /> 
	<link href="/static/img/favicon.ico" mce_href="/static/img/favicon.ico" rel="shortcut icon" type="image/x-icon" /> 
</head>
<body>
	<div class="main">
		<div id='ichart-render'></div>
	</div>
</body>
</html>
<script>
$(function(){	
	$.ajax({
	   type: "get",
	   url: "/static/upload/data.json",
	   dataType: "json",
	   success: function(json_data){
	     var data = json_data;

	     var chart = new iChart.Donut2D({
			render:"ichart-render",
            width : 900,
			height : 600,
            background_color:"#fefefe",
            gradient:false,
            color_factor:0.2,
            border:{
                  color:"BCBCBC",
                  width:0
            },
            tip:{
				enable:true,
				showType:'fixed'
			},
            align:"center",
            offsetx:0,
            offsety:0,

            shadow:true,
            shadow_color:"#666666",
            shadow_blur:2,
            showpercent:true,
            decimalsnum:2,
            column_width:"70%",
            bar_height:"70%",
            radius:"90%",
			center:{
				text:'CORE\nLANGUAGE',
				shadow:true,
				shadow_offsetx:0,
				shadow_offsety:2,
				shadow_blur:2,
				shadow_color:'#b7b7b7',
				color:'#6f6f6f'
			},
            title:{
                  text:"使用过的语言统计",
                  color:"#111111",
                  fontsize:20,
                  font:"微软雅黑",
                  textAlign:"center",
                  height:30,
                  offsetx:0,
                  offsety:0
            },
   			//sub_option:{
			// 	label:false,
			// 	color_factor : 0.01
			// },
            sub_option:{
                  border:{
                        color:"#BCBCBC",
                        width:1
                  },
                  label:{
                        fontweight:500,
                        fontsize:11,
                        color:"#4572a7",
                        sign:"square",
                        sign_size:12,
                        border:{
                              color:"#BCBCBC",
                              width:1
                        },
                        background_color:"#fefefe"
                  }
            },
            legend:{
                  enable:true,
                  background_color:"#fefefe",
                  color:"#333333",
                  fontsize:12,
                  border:{
                        color:"#BCBCBC",
                        width:1
                  },
                  column:1,
                  align:"right",
                  valign:"center",
                  offsetx:0,
                  offsety:0
            },
            // type:"pie2d",
            data:data,
		});

	     chart.draw();
	   }
	});
	
});
</script>