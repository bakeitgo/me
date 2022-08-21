#bash 

* Data types in bash: number/integer, string, boolean

* bash support default values for parameters, see `man bash` /Parameter Expansion

* We can use `test` *[ expr ]* or *[[ expr ]]* (last one is modern version) to evaluate status of files / cheeck values ->  `man test` *Note: **>**, **<** must be escaped or wrapped with quotes, otherwise shell will interpret those as redirection operators*

* `(( expr )) - designed for integers -> used to perform arithmetic truth tests, instead of `[[ "$INT" -lt 0 ]]` we can use `(( INT < 0 ))`

* Logical operators

| Operator | test [[ ]] | test (( )) |
|----------|------------|------------|
| AND | -a | && |
| OR | -o | || |
| NOT | ! | ! |

* first && second - second runs only if first evalutes to *true*

* first || second - second runs only if first evaluates to *false*



# Commands

* `local varname` - declare local scope

* `declare varname` - declare variable with specific type using options

* `unset varname` - unset a variable 

* `$?` - return / exit status of a last used code (0 - true, 1~ false)

* `$1`, `$2`, etc.

* `$#` - displays count of args passed to script

* `read` - read values from stdin ( we cannot use it with *pipeline*, use redirecting input instead)




#bash #parameters #local #declare #unset #exit #return
