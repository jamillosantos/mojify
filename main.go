package main

import (
	"bufio"
	"bytes"
	"io"
	"os"
	"unicode"

	"github.com/jamillosantos/gitmoji/emoji"
)

func main() {
	bufIn := bufio.NewReader(os.Stdin)
	bufOut := bufio.NewWriter(os.Stdout)
	defer bufOut.Flush()
	for {
		r, _, err := bufIn.ReadRune()
		if err == io.EOF {
			return
		}
		if err != nil {
			panic(err)
		}
		if r == ':' {
			emojiBuf := bytes.NewBuffer(nil)
			for {
				r, _, err = bufIn.ReadRune()
				if err == io.EOF {
					bufOut.WriteRune(':')
					bufOut.Write(emojiBuf.Bytes())
					return
				}
				if err != nil {
					panic(err)
				}
				if unicode.IsSpace(r) || r == '\n' || r == '\r' {
					bufOut.WriteRune(':')
					bufOut.Write(emojiBuf.Bytes())
					bufOut.WriteRune(r)
					break
				} else if r == ':' {
					emojiKey := ":" + emojiBuf.String() + ":"
					emojiCode, ok := emoji.EmojiCodeMap[emojiKey]
					if !ok {
						bufOut.WriteRune(':')
						bufOut.WriteString(emojiKey)
						bufOut.WriteRune(':')
						break
					}
					bufOut.WriteString(emojiCode)
					break
				} else {
					emojiBuf.WriteRune(r)
				}
			}
		} else {
			bufOut.WriteRune(r)
			if r == '\n' {
				bufOut.Flush()
			}
		}
	}
}
