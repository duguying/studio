module.exports = function(grunt) {
  grunt.initConfig({
    pkg: grunt.file.readJSON('package.json'),

    sass: {
      default_theme: {
        options: {
          sourcemap: 'auto'
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
            'deps/jquery/*.js',
            'deps/jquery/**/*.js',
            'deps/custom/*.js'
        ],
        dest: 'deps/dist/dependence.js'
      },

      blog_js: {
        options: {
          separator: '\n;\n',
          stripBanners: true,
          banner: '/*! hello - v1.2.3 - 2014-2-4 */'
        },
        src: [
          'deps/angular/*.js',
          'deps/angular/**/*.js',
          'deps/duoshuo/*.js',
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
        src: ['deps/dist/dependence.js'],
        dest: 'deps/dist/dependence.min.js'
      },

      blog: {
        src: ['theme/default/dist/blog.js'],
        dest: 'theme/default/dist/blog.min.js'
      },
    },

    clean: [
      "deps/dist",
      "theme/default/dist",
      "theme/default/css/*.css",
      "theme/default/css/*.map"
    ]

  });
  grunt.loadNpmTasks('grunt-contrib-sass');
  grunt.loadNpmTasks('grunt-contrib-uglify');
  grunt.loadNpmTasks('grunt-contrib-concat');
  grunt.loadNpmTasks('grunt-contrib-clean');

  grunt.registerTask('default', ['sass','concat','uglify']);
};
