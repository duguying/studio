/**
 * SyntaxHighlighter
 * http://alexgorbatchev.com/
 *
 * This brush was originally created by Ildar Shaimordanov
 * homepage:   http://with-love-from-siberia.blogspot.com/
 * brush page: http://with-love-from-siberia.blogspot.com/2009/07/finest-code-syntax-highlighter.html
 */
SyntaxHighlighter.brushes.Cmd = function()
{
    var commands = 'ASSOC AT ATTRIB BREAK CACLS CD CHCP CHDIR CHKDSK CHKNTFS CLS CMD COLOR COMP COMPACT CONVERT COPY DATE '
                 + 'DEL DIR DISKCOMP DISKCOPY DOSKEY ECHO ERASE EXIT FC FIND FINDSTR FORMAT FTYPE GRAFTABL '
                 + 'HELP LABEL MD MKDIR MODE MORE MOVE PATH PAUSE POPD PRINT PROMPT PUSHD RD RECOVER REN RENAME REPLACE '
                 + 'RMDIR SHIFT SORT START SUBST TIME TITLE TREE TYPE VER VERIFY VOL XCOPY';

    var keywords = 'CON DEFINED DO ENABLEDELAYEDEXPANSION ENABLEEXTENSIONS ENDLOCAL FOR GOTO CALL IF IN ELSE NOT NUL REM SET SETLOCAL';

    var variables = 'ALLUSERSPROFILE APPDATA CommonProgramFiles COMPUTERNAME ComSpec DATE FP_NO_HOST_CHECK HOMEDRIVE '
                  + 'HOMEPATH LOGONSERVER NUMBER_OF_PROCESSORS OS Path PATHEXT PROCESSOR_ARCHITECTURE PROCESSOR_IDENTIFIER '
                  + 'PROCESSOR_LEVEL PROCESSOR_REVISION ProgramFiles PROGS PROMPT SANDBOX_DISK SANDBOX_PATH SESSIONNAME '
                  + 'SystemDrive SystemRoot TEMP TIME TMP USERDNSDOMAIN USERDOMAIN USERNAME USERPROFILE windir';

    this.regexList = [
        //
        // REM Comments
        //
        {
            regex: /(^::|rem).*$/gmi,
            css: 'comments'
        },

        //
        // "Strings"
        // 'Strings'
        // `Strings`
        // ECHO String
        //
        {
            regex: SyntaxHighlighter.regexLib.doubleQuotedString,
            css: 'string' 
        },
        {
            regex: SyntaxHighlighter.regexLib.singleQuotedString,
            css: 'string' 
        },
        {
            regex: /`(?:\.|(\\\`)|[^\``\n])*`/g,
            css: 'string' 
        },
        {
            regex: /echo.*$/gmi,
            css: 'string'
        },

        //
        // :Labels
        //
        {
            regex: /^:.+$/gmi,
            css: 'color3' 
        },

        //
        // %Variables%
        // !Variables!
        //
        {
            regex: /(%|!)\w+\1/gmi,
            css: 'variable'
        },

        //
        // %%a variable Refs
        // %1 variable Refs
        //
        {
            regex: /%\*|%%?~?[fdpnxsatz]*[0-9a-z]\b/gmi,
            css: 'variable' 
        },

        //
        // keywords
        //
        {
            regex: new RegExp(this.getKeywords(commands), 'gmi'),
            css: 'keyword'
        },

        //
        // keywords
        //
        {
            regex: new RegExp(this.getKeywords(keywords), 'gmi'),
            css: 'keyword' 
        }
    ];
};

SyntaxHighlighter.brushes.Cmd.prototype = new SyntaxHighlighter.Highlighter();
SyntaxHighlighter.brushes.Cmd.aliases = ['bat', 'cmd', 'batch'];
