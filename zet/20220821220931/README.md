# bash

### Loops

* `While`

	* `while commands; do commands; done`

	* break breaks loop, continue skipps current iteration

* `Until` - same as *while* loop, but does the opposite (use opposite expression to do same as *while* loop

* Loops process stdin via redirection e.g. `<`

```bash
#!/bin/bash

# while-read: read lines from a file

while read distro version release; do
	printf "Distro: %s\tVersion: %s\tReleased: %s\n" \
		"$distro" \
		"$version" \
		"$release"
done < distros.txt
```
