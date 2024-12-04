package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r *rot13Reader) Read(p []byte) (int, error) {
	n, err := r.r.Read(p)
	if err != nil {
		return n, err
	}
	
	// https://qiita.com/ykmkn/items/8a66e02f3b57b04e24a9
	for i := 0; i < len(p); i++ {
	    if p[i] >= 'A' && p[i] <= 'Z' {
		    p[i] = (p[i]-'A'+13)%26 + 'A'
	    } else if p[i] >= 'a' && p[i] <='z' {
		    p[i] = (p[i]-'a'+13)%26 + 'a'
	    }
	}
	return n, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}

