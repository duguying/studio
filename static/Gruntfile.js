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
      js: {
        src: 'js/admin/*.js',
        dest: 'build/admin.min.js'
      }
    }
    
  });
  grunt.loadNpmTasks('grunt-sass');
  grunt.loadNpmTasks('grunt-contrib-uglify');
  grunt.registerTask('default', ['sass','uglify']);
};