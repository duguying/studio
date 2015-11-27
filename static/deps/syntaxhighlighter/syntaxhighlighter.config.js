var alloy={};
alloy.path=function(){
  var args = arguments,
      result = []
      ;

  // var json = $.ajax({
  //   url: "/map.json",
  //   async: false
  // }).responseText;

  var data = eval('({})');
  // TODO

  for(var i = 0; i < args.length; i++){
    var item = args[i].replace('@', '!/static/deps/syntaxhighlighter/scripts/')
    item = item.split('!');
    item = item[0] + item[1];
    result.push(item);
  }
      
  return result
};
SyntaxHighlighter.autoloader.apply(null, alloy.path(
  'applescript            @shBrushAppleScript.js',
  'actionscript3 as3      @shBrushAS3.js',
  'bash shell             @shBrushBash.js',
  'bat batch              @shBrushBat.js',
  'coldfusion cf          @shBrushColdFusion.js',
  'cpp c                  @shBrushCpp.js',
  'c# c-sharp csharp      @shBrushCSharp.js',
  'css                    @shBrushCss.js',
  'delphi pascal          @shBrushDelphi.js',
  'diff patch pas         @shBrushDiff.js',
  'erl erlang             @shBrushErlang.js',
  'groovy                 @shBrushGroovy.js',
  'go golang              @shBrushGo.js',
  'java                   @shBrushJava.js',
  'jfx javafx             @shBrushJavaFX.js',
  'js jscript javascript  @shBrushJScript.js',
  'tex latex              @shBrushLatex.js',
  'perl pl                @shBrushPerl.js',
  'php                    @shBrushPhp.js',
  'ps powershell          @shBrushPowerShell.js',
  'text plain             @shBrushPlain.js',
  'py python              @shBrushPython.js',
  'ruby rails ror rb      @shBrushRuby.js',
  'sass scss              @shBrushSass.js',
  'scala                  @shBrushScala.js',
  'sql                    @shBrushSql.js',
  'vb vbnet               @shBrushVb.js',
  'xml xhtml xslt html    @shBrushXml.js'
));
SyntaxHighlighter.defaults.toolbar = false;
SyntaxHighlighter.defaults.title = '';
SyntaxHighlighter.all();

/**
 * back to top config
 */

var mv_dynamic_to_top = {"text":"To Top","version":"0","min":"300","speed":"300","easing":"easeInOutExpo","margin":"20"};

/*
 * Dynamic To Top Plugin
 * http://www.mattvarone.com
 *
 * By Matt Varone
 * @sksmatt
 *
 */
var mv_dynamic_to_top;(function($,mv_dynamic_to_top){jQuery.fn.DynamicToTop=function(options){var defaults={text:mv_dynamic_to_top.text,min:parseInt(mv_dynamic_to_top.min,10),fade_in:600,fade_out:400,speed:parseInt(mv_dynamic_to_top.speed,10),easing:mv_dynamic_to_top.easing,version:mv_dynamic_to_top.version,id:'dynamic-to-top'},settings=$.extend(defaults,options);if(settings.version===""||settings.version==='0'){settings.text='<span>&nbsp;</span>';}
if(!$.isFunction(settings.easing)){settings.easing='linear';}
var $toTop=$('<a href=\"#\" id=\"'+settings.id+'\"></a>').html(settings.text);$toTop.hide().appendTo('body').click(function(){$('html, body').stop().animate({scrollTop:0},settings.speed,settings.easing);return false;});$(window).scroll(function(){var sd=jQuery(window).scrollTop();if(typeof document.body.style.maxHeight==="undefined"){$toTop.css({'position':'absolute','top':sd+$(window).height()-mv_dynamic_to_top.margin});}
if(sd>settings.min){$toTop.fadeIn(settings.fade_in);}else{$toTop.fadeOut(settings.fade_out);}});};$('body').DynamicToTop();})(jQuery,mv_dynamic_to_top);

$("img").parent("a").fancybox();
