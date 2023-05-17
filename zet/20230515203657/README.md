# C lang

## Elements of C:

### Tokens

* Each basic element is considered as a token:

> _token:_   -> the smallest unit of code, below all listed tokens
		_keyword_  -> e.g. **int**, **main**, **for** etc.
		_identifier_ -> name of function, variable
		_constant_ -> as name describes
		_string-literal_ -> e.g. "some string", includes special sequences
		_operator_ -> logical, arithmetic, assign, compare, etc.
		_punctuator_ -> organizes program text or define tasks which compiler carries out

	* White-space characters - everything like \n \t etc, is considered as white space char, and C compiler ignore those, same as comments.

	* Comment - you can comment multiline via /\* _write something here_ \*/,  or singleline via _//_, **BUT YOU CAN COMMENT NEXT LINE WHEN COMMENTING WITH SINGLELINE BY USING "\" ESCAPE SEQUENCE** they cannot be nested and are ignored by C compiler, you can use comments for test purposes, but you can achieve same with preprocessor directives like _#if_ and _#endif_


	* Token evaluation - when compiler interprets tokens, it tries to consume as many characters as possible this is why **i+++j** evaluates as **i++ + j**, not as **i+ ++j** as many ppl would think.

	* Keywords are reserved, they cant be used as identifiers, with some special circumstances, more [here](https://learn.microsoft.com/en-us/cpp/c-language/c-keywords?view=msvc-170#standard-c-keywords)

	* Identifiers with file-level scope should not be named with _underscore and lowercase letter_, _two underscores_ and _underscore and upper-case letter_. They ANSI C standard allows those identifier names to be **reserved** for compiler use.
	> valid identifiers: 
	j
	count
	temp1
	top\_of\_page
	skip12
	LastNum
	
	* Multibyte and wide characters - **Multibyte character** is a character which is built via sequence of one or more bytes. Each sequence represents each character. Those characters are used in japanese lang for e.g. (Kanji), 
**Wide characters** are multilingual character codes (multilingual character codes are representation of characters via numeric codes in digital form aka character encodings). They are always 16 bits wide, type is *wchar_t*. e.g.

	* Constant - There are 3 types of constants in C - **Floating point**, **Integer**, **Character**

	* Punctuators - **() [] {} * , : = ; ... #**
  

	
## Related to: 

* https://learn.microsoft.com/en-us/cpp/c-language/?view=msvc-170


#c #programming #basics #elements #tokens #comment #constant #identifier #keywords #char #character
