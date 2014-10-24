$ document
.ready (e) ->
	$ "#about"
	.mouseover (e) -> 
		$(".drop-menu").show()
	.mouseout (e) ->
		$(".drop-menu").hide()
