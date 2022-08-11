# Readers

* What is a reader?

    * It's an interface which represents *the read end of a stream of data*. It's a part of `io` pckg

    * *Golang* standard library contains many *implementations* of this interface, including: files, network connections, compressors, ciphers and others

    * `io.Reader` interface has a `Read` method: 
    ##
    `func (T) Read(b []byte) (n int, err error)`

    * Read fill given `byte slice` with data and returns number of bytes `filled` and an `error` value. When stream ends it returns an `io.EOF` error

    * The example code creates a strings.Reader and consumes its output 8 bytes at a time. 

```go
package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Reader!")

	b := make([]byte, 8)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}
```


# Related to: 

* https://go.dev/doc/

#golang #readers