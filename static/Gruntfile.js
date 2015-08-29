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
          'css/style.css': ['!css/_common.scss','css/style.scss'],
          'css/admin.css': ['!css/_common.scss','css/admin.scss']
        }
      }
    },

    uglify: {
      options: {
        banner: '/*! <%= pkg.name %> <%= grunt.template.today("yyyy-mm-dd") %> */\n'
      },
      jslib: {
        src: ['js/global/jquery.min.js','js/guest/jquery.fancybox.min.js'],
        dest: 'js/lib.min.js'
      },
      jsarticle: {
        src: [
          'syntaxhighlighter/scripts/shCore.js',
          'syntaxhighlighter/scripts/shAutoloader.js',
          'js/guest/syntaxhighlighter.config.js',
          'js/guest/main.js'
        ],
        dest: 'js/article.min.js'
      },
      jsmain: {
        src: 'js/guest/main.js',
        dest: 'js/main.min.js'
      },
      js: {
        src: 'js/admin/*.js',
        dest: 'js/admin.min.js'
      }
    }
    
  });
  grunt.loadNpmTasks('grunt-sass');
  grunt.loadNpmTasks('grunt-contrib-uglify');
  grunt.registerTask('default', ['sass','uglify']);
};