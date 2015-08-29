;(function(angular, NProgress) {
  'use strict';

  if (!angular) 
    throw new Error('Angular.js required!');

  var NProgressExist = NProgress && NProgress.start && NProgress.done;

  angular
    .module('duoshuo', [])
    // API set
    .provider('duoshuo', duoshuoProvider)
    // Directives
    .directive('dsThread', createDirective('ds-thread'))
    .directive('dsRecentComments', createDirective('ds-recent-comments'))
    .directive('dsRecentVisitors', createDirective('ds-recent-visitors'))
    .directive('dsThreadCount', createDirective('ds-thread-count'))
    .directive('dsTopThreads', createDirective('ds-top-threads'))
    .directive('dsLogin', createDirective('ds-login'));

  function duoshuoProvider() {
    this.config = config;

    function config(configs) {
      if (!configs)
        return;
      if (!configs.short_name)
        throw new Error('duoshuo.config(); `short_name` is required');

      window.duoshuoQuery = configs;
    }

    this.$get = ['$rootScope',
      function($rootScope) {
        var duoshuo = {};

        // Lowlevel API set
        angular.forEach(['get', 'post', 'ajax'], function(method) {
          duoshuo[method] = function(endpoint, data, callback, errorCallback, skipCheck) {
            if (!window.DUOSHUO)
              throw new Error('duoshuo embed.js is required!');

            var API = window.DUOSHUO.API;
            if (!API)
              throw new Error('duoshuo embed.js must be unstable version!');

            if (NProgressExist) NProgress.start();

            return API[method](endpoint, data, function(result) {
              if (NProgressExist) NProgress.done();

              callback(
                (result.code === 0) ? null : new Error(result.code + ' ' + result.errorMessage),
                result.response,
                result
              );

              if (!skipCheck) 
                $rootScope.$apply();

              return;
            }, function(err) {
              if (NProgressExist) 
                NProgress.done();

              if (errorCallback && typeof(errorCallback) === 'function') {
                return errorCallback(err);
              }
              return;
            });
          }
        });

        // Event wrapper
        duoshuo.on = function(eve, callback, skipCheck) {
          if (['reset', 'ready'].indexOf(eve) === 0)
            return callback(new Error('event not found'));
          var e = eve;
          if (e === 'ready') e = 'reset';
          return window.DUOSHUO.visitor.on(e, function() {
            var self = this;
            var data = this.data;
            callback(null, data, self);
            if (!skipCheck) $rootScope.$apply();
            return;
          });
        };

        // Comments renderer
        duoshuo.render = function(options) {
          if (!window.DUOSHUO || !window.DUOSHUO.initSelector)
            throw new Error('createDirective(); duoshuo embed.js is required!');

          return window.DUOSHUO.initSelector(
            '.ds-thread',
            window.DUOSHUO.selectors['.ds-thread']
          )
        };

        return duoshuo;
      }
    ];
  }

  function createDirective(type) {
    return function directive() {
      return {
        restrict: 'AE',
        replace: true,
        template: '<div class="' + type + '">',
        link: function(scope, element, attrs) {
          angular.element(document).ready(function() {
            if (!window.DUOSHUO || !window.DUOSHUO.initSelector)
              return;

            // Trigger init selector function 
            window.DUOSHUO
              .initSelector('.' + type, window.DUOSHUO.selectors['.' + type])
          });
        }
      };
    }
  }

})(window.angular, window.NProgress);
