/*
A common pattern is an io.Reader that wraps another io.Reader, modifying the stream in some way.
For example, the gzip.NewReader function takes an io.Reader (a stream of compressed data) and returns a *gzip.Reader that also implements io.Reader (a stream of the decompressed data).
Implement a rot13Reader that implements io.Reader and reads from an io.Reader, modifying the stream by applying the rot13 substitution cipher to all alphabetical characters.
The rot13Reader type is provided for you. Make it an io.Reader by implementing its Read method. 
*/

package main

import (
	"io"
	"os"
	"strings"
)

func (m rot13Reader) Read(p []byte) (val int, er error) {
	val, er = m.r.Read(p)
	for i := 0; i < len(p); i++ {
		b := p[i]
		if (b >= 'A' && b <= 'M') || (b >= 'a' && b <= 'm') {
			b += 13
		} else if (b >= 'N' && b <= 'Z') || (b >= 'n' && b <= 'z') {
			b -= 13
		}
		p[i] = b
	}

	return val, er
}

type rot13Reader struct {
	r io.Reader
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

/*
You cracked the code!
Program exited.
*/



