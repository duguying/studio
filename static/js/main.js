/**
 * main.js
 * @author rex
 */
;$(document).ready(function (e) {
	NProgress.done();

	$("#about").mouseover(function (e) {
		$(".drop-menu").show();
	}).mouseout(function (e) {
		$(".drop-menu").hide();
	});
});