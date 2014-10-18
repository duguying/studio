module.exports = function(grunt) {

  grunt.initConfig({
    pkg: grunt.file.readJSON('package.json'),

    sass: {                              // Task
      dist: {                            // Target
        options: {                       // Target options
          style: 'compressed'
        },
        files: {                         // Dictionary of files
          '../static/css/style.min.css': 'sass/style.scss',       // 'destination': 'source'
          '../static/css/admin.min.css': 'sass/admin.scss'
        }
      }
    },

    coffee: {
      compile: {
        files: {
          'build/admin-footer.js': 'coffee/admin.coffee',
          'build/main-footer.js': 'coffee/main.coffee',
        }
      },
    },

    uglify: {
      options: {
        banner: '/*! <%= pkg.name %> <%= grunt.template.today("dd-mm-yyyy") %> */\n'
      },
      dist: {
        files: {
          '../static/js/main.footer.min.js': ['build/main-footer.js','js/main-config.js'],
          '../static/js/admin.footer.min.js': ['deps/ajaxfileupload.js','build/admin-footer.js'],
        }
      }
    },

  });

  grunt.loadNpmTasks('grunt-contrib-sass');
  grunt.loadNpmTasks('grunt-contrib-coffee');
  grunt.loadNpmTasks('grunt-contrib-uglify');

  grunt.registerTask('default', ['sass','coffee','uglify']);

};