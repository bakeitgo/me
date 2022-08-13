# bash / configuring vim

## UNIX Filter

`:.! command` - place stdout of `command` in the current line
`:.,.+2!bash` - run code from current line to current line + 2, and use bash to interpret it.
`nl` - numerise all lines
`!}` - select couple of lines on which we will make operations. e.g. it selects `.,.+9!`  

     1	thING
     2	thING
     3	thING
     4	thING
     5	thING
     6	thING
     7	thING
     8	thING
     9	thING
    10	thING

***
# Related to:

* https://youtu.be/yu2kO2yb7L8?t=5057

#bash #vim #scripting #scripts
