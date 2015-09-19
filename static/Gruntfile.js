module.exports = function(grunt) {
  grunt.initConfig({
    pkg: grunt.file.readJSON('package.json'),

    sass: {
      compile: {
        options: {
          sourceMap: true,
          banner: '/*!\n' +
              ' * Package Name: <%= pkg.name %>\n' +
              ' * Author: <%= pkg.author.name %>\n' +
              ' * Version: <%= pkg.version %>\n' +
              ' * Tags:\n' +
              ' */\n'
        },
        files: {
          'css/style.css': ["css/style.scss","!css/_*.scss"],
          'css/admin.css': ["css/admin.scss","!css/_*.scss"]
        }
      }
    },

    uglify: {
      options: {
        banner: '/*! <%= pkg.name %> <%= grunt.template.today("yyyy-mm-dd") %> */\n',
        mangle: {
            except: ['jQuery', 'angular', 'adminService', 'ueditor', 'ngRoute']
        },
      },
      blog: {
        src: [
          'js/global/jquery/*.js',
          'js/global/jquery-plugin/*.js',
          'syntaxhighlighter/scripts/shCore.js',
          'syntaxhighlighter/scripts/shAutoloader.js',
          'syntaxhighlighter/syntaxhighlighter.config.js',
          'js/global/custom/*.js',
          'js/page/blog/*.js'
        ],
        dest: 'build/blog.min.js'
      },
      admin: {
        src: [
          // 'js/global/jquery/*.js',
          'js/global/angular/angular.min.js',
          'js/global/angular-module/angular-route.min.js',
          'js/global/angular-module/angular-ueditor',
          // 'js/global/angular/angular-route.min.js',
          // 'js/global/angular/angular-ueditor.js',
          'js/page/admin/admin.js',
          'js/page/admin/directive.js',
          'js/page/admin/controller.js',
        ],
        dest: "build/admin.min.js"
      }
    }

  });
  grunt.loadNpmTasks('grunt-sass');
  grunt.loadNpmTasks('grunt-contrib-uglify');
  grunt.registerTask('default', ['sass','uglify']);
};
