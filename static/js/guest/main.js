(function() {
  $(document).ready(function(e) {
    return $("#about").mouseover(function(e) {
      return $(".drop-menu").show();
    }).mouseout(function(e) {
      return $(".drop-menu").hide();
    });
  });

}).call(this);
