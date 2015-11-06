module.exports = function(grunt) {
  grunt.initConfig({
    pkg: grunt.file.readJSON('package.json'),

    sass: {
      default_theme: {
        options: {
          sourceMap: true,
          banner: '/*!\n' +
              ' * Package Name: <%= pkg.name %>\n' +
              ' * Author: <%= pkg.author.name %>\n' +
              ' * Theme: Default\n' +
              ' * Version: <%= pkg.version %>\n' +
              ' * Tags:\n' +
              ' */\n'
        },
        files: {
          'theme/default/css/common.css': ["theme/default/css/common.scss","!theme/default/css/_*.scss"],
          'theme/default/css/style.css': ["theme/default/css/style.scss","!theme/default/css/_*.scss"],
          'theme/default/css/admin.css': ["theme/default/css/admin.scss","!theme/default/css/_*.scss"]
        }
      }
    },

    concat: {
      css: {
        options: {
          separator: '\n\n',
          stripBanners: true,
          banner: '/*! hello - v1.2.3 - 2014-2-4 */'
        },
        src: [
          'theme/default/css/common.css',
          'theme/default/css/style.css',
          'theme/default/css/admin.css'
        ],
        dest: 'theme/default/dist/blog.css'
      },

      js: {
        options: {
          separator: '\n;\n',
          stripBanners: true,
          banner: '/*! hello - v1.2.3 - 2014-2-4 */'
        },
        src: [
            'dependence/jquery/*.js',
            'dependence/jquery/**/*.js',
            'dependence/custom/*.js'
        ],
        dest: 'dependence/dist/dependence.js'
      },

      blog_js: {
        options: {
          separator: '\n;\n',
          stripBanners: true,
          banner: '/*! hello - v1.2.3 - 2014-2-4 */'
        },
        src: [
          'dependence/angular/*.js',
          'dependence/angular/**/*.js',
          'dependence/duoshuo/*.js',
          'theme/default/js/page/**/*.js'
        ],
        dest: "theme/default/dist/blog.js"
      }
    },

    uglify: {
      options: {
        banner: '/*! <%= pkg.name %> <%= grunt.template.today("yyyy-mm-dd") %> */\n',
        mangle: {
            except: ['jQuery', 'angular', 'ueditor', 'ngRoute']
        },
      },

      deps: {
        src: ['dependence/dist/dependence.js'],
        dest: 'dependence/dist/dependence.min.js'
      },

      blog: {
        src: ['theme/default/dist/blog.js'],
        dest: 'theme/default/dist/blog.min.js'
      },
    }

  });
  grunt.loadNpmTasks('grunt-sass');
  grunt.loadNpmTasks('grunt-contrib-uglify');
  grunt.loadNpmTasks('grunt-contrib-concat');
  grunt.registerTask('default', ['sass','concat','uglify']);
};
