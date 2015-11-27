/*global window,location*/
(function (window) {
  'use strict';

  var Simditor = window.Simditor;
  var directives = angular.module('simditor',[]);

  directives.directive('simditor', function () {
    
    var TOOLBAR_DEFAULT = ['title', 'bold', 'italic', 'underline', 'strikethrough', '|', 'ol', 'ul', 'blockquote', 'code', 'table', '|', 'link', 'image', 'hr', '|', 'indent', 'outdent'];
    
    return {
      require: "?^ngModel",
      link: function (scope, element, attrs, ngModel) {
        element.append("<div style='height:300px;'></div>");

        var toolbar = scope.$eval(attrs.toolbar) || TOOLBAR_DEFAULT;
        scope.simditor = new Simditor({
          textarea: element.children()[0],
          pasteImage: true,
          toolbar: toolbar,
          defaultImage: 'assets/images/image.png',
          upload: location.search != '?noupload' ? {
              url: '/upload'
          } : false
        });

        var $target = element.find('.simditor-body');
        
        function readViewText() {

            ngModel.$setViewValue($target.html());

            if (attrs.ngRequired != undefined && attrs.ngRequired != "false") {
                
                var text = $target.text();
                
                if(text.trim() === "") {
                    ngModel.$setValidity("required", false);
                } else {
                    ngModel.$setValidity("required", true);
                }
            }

        }

        ngModel.$render = function () {
          scope.simditor.focus();
          $target.html(ngModel.$viewValue);
        };

        scope.simditor.on('valuechanged', function () {
          scope.$apply(readViewText);
        });
      }
    };
  });
}(window));
