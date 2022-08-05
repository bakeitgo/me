# Exported names

* How to export names in GO?
	* Name is exported when it begins with uppercase letter e.g *Pizza*.
	* Names which are written in lowercase, aren't exported
	* When importing packets, you can refer to names which has been exported. Names which are not exported, aren't available apart from *package* where has been declared.
	* In order to consume exported name, you have to use *packagename.Method* (Note: , method needs to begin with uppercase letter)
	* When name is not found, *undefined* err is printed to stderr










# Related to: 

* https://go.dev/doc/



#golang #export #exportednames #exported-names
