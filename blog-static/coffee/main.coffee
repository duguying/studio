$ document
.ready (e) ->
	NProgress.done()
	$ "#about"
	.mouseover (e) -> 
		$(".drop-menu").show()
	.mouseout (e) ->
		$(".drop-menu").hide()
