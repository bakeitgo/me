# REGEX



* What is regular expression?
	* Schema which impose how string should look like

* Why is it better to allow rather than deny?
	* When you deny, things gonna be complicated, it is better to specify which string you want to allow.

* REGEX
	* `a` - matches *a* character in whole string
	* `^a` - matches only string which start with *a*
	* `a$` - matches only string which ends with *a*
	* `^exact$` - matches only *exact* word
	* `cat|dog` - matches *cat* or *dog* word in whole string
	* `(cat|dog)foot - matches *catfoot* or *dogfoot*
	* ` . ` - any single character
	* ` .... ` | `.{4}` - everything which have >=4 characters
	* `^....$` | `^.{4}$` - everything which have =4 characters
	* `^.{1,5}$ - everything which have between =1 & =5 chars
	* `^.{6,}$ - everything which have >=6 characters
	* `^.{1,}$ | `^.+$` - everything which have >=1 chars
	* `^.{0,}$ | ` ^.*$ ` - everything*
	* `^\.*$ - every *.* character
	* `[adls]` - every string which contains *a,d,l,s* chars
	* `^[^a]` - everything which not begins with *a* char
	* `[^t]$` - everything which not ends with *t* char
	* 

## Related:

* https://youtu.be/z7nluhUWYRs?t=11790

### Theory: 

* https://github.com/manish-old/ebooks-2/blob/master/O'Reilly%20-%20Mastering%20Regular%20Expressions%202nd%20Edition.pdf

* "Missing Semester" on YouTube (Regular Expressions)

### Practise:

* https://alf.nu/RegexGolf?world=regex&level=r00

## Commands: 


#REGEX #regularexpression
