fis.config.set('modules.parser.scss', 'sass');
fis.config.set('roadmap.ext.scss', 'css');

fis.config.set('modules.parser.coffee', 'coffee-script');
fis.config.set('roadmap.ext.coffee', 'js');
fis.config.set('settings.optimizer.uglify-js.output.ascii_only', true);

fis.config.set('modules.parser.less', 'less');
fis.config.set('roadmap.ext.less', 'css');

fis.config.set('roadmap.path',[
    {
        reg: /^\/sass\/(.*)/i,
        release: '/static/css/$1'
    },
    {
        reg: /^\/js\/(.*)/i,
        release: '/static/js/$1'
    },
    {
        reg: /^\/syntaxhighlighter\/(.*)/i,
        release: '/static/syntaxhighlighter/$1'
    },
    {
        reg: /^\/ueditor\/(.*)/i,
        release: '/static/ueditor/$1'
    },
    {
        reg: /^\/img\/(.*)/i,
        release: '/static/img/$1'
    }
]);

fis.config.set('project.exclude', [
    /^\/(.*)\.md/i,
    /^\/(.*)\.html/i,
    /^\/sass\/_(.*)\.scss/i,
    /^\/js\/guest\/(.*)/i,
    /^\/js\/admin\/(.*)/i,
]);

fis.config.merge({
    deploy : {
		local : {
            to : '../'
        },

        remote: {
            receiver: 'http://duguying.net/fis?key=xxxxxxxxxxx',
            to: '/root/gopath/src/github.com/duguying/blog'
        }
    },
    
});