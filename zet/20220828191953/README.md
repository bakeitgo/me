# Bash

* How to capitalise first letter in param?

	* Via e.g. funcname "${1^}"

* How to capitalise all string to uppercase?

	* Via e.g. funcname "${1^^}

* Dont use **SINGLE BRACKETS AS IN POSIX** in bash scripting" (param expansion)

	* `shellcheck` doesnt catch it

* IF statement syntax

```bash
#!/bin/bash

greet () {
	local name="there"
	if [[ -n "$1" ]]; then
		name=$1
	else
		name="there"
	fi
	echo "Hello, $name"
}
greet "${1^}" 
```

## Related to:

* https://youtu.be/OX3EvBq3bMs?t=2695
