SyntaxHighlighter.brushes.Go = function()
{
	// Copyright 2010 Sanchez, Allister Levi

	var declaration_type_words = 'struct func interface map chan package import type const var';

	var control_words = 'goto break continue if else switch default case for range go select return fallthrough defer';

	this.regexList = [
		{ regex: SyntaxHighlighter.regexLib.singleLineCComments,	css: 'comments' },			// one line comments
		{ regex: SyntaxHighlighter.regexLib.multiLineCComments,		css: 'comments' },			// multiline comments
		{ regex: SyntaxHighlighter.regexLib.doubleQuotedString,		css: 'string' },			// strings
		{ regex: SyntaxHighlighter.regexLib.singleQuotedString,		css: 'string' },			// strings
		{ regex: new RegExp(this.getKeywords(declaration_type_words), 'gm'),		css: 'color2 bold' },
		{ regex: new RegExp(this.getKeywords(control_words), 'gm'),		css: 'keyword bold' }
		];
};

SyntaxHighlighter.brushes.Go.prototype	= new SyntaxHighlighter.Highlighter();
SyntaxHighlighter.brushes.Go.aliases	= ['go', 'golang'];
