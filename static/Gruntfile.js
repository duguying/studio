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
          'build/css/style.css': ['!sass/_common.scss','sass/style.scss'],
          'build/css/admin.css': ['!sass/_common.scss','sass/admin.scss']
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