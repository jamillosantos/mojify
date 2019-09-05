package main

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"os"
	"os/exec"
	"unicode"

	"github.com/jamillosantos/gitmoji/emoji"
)

type emojiWriter struct {
	r         rune
	returnEOF bool
	writer    *bufio.Writer
}

func newEmojiWriter(returnEOF bool, writer io.Writer) *emojiWriter {
	return &emojiWriter{
		writer:    bufio.NewWriter(writer),
		returnEOF: returnEOF,
	}
}

func (w *emojiWriter) Write(p []byte) (int, error) {
	n := 0
	buf := bytes.NewBuffer(p)
	for {
		r2, _, err := buf.ReadRune()
		if err == io.EOF {
			if w.returnEOF {
				return n, err
			}
			return n, nil
		}
		if err != nil {
			return n, err
		}
		n++
		w.r = r2
		if w.r == ':' {
			emojiBuf := bytes.NewBuffer(nil)
			for {
				w.r, _, err = buf.ReadRune()
				if err == io.EOF {
					if w.returnEOF {
						return n, err
					}
					return n, nil
				}
				if err != nil {
					w.writer.WriteRune(':')
					w.writer.Write(emojiBuf.Bytes())
					return n, err
				}
				n++
				if unicode.IsSpace(w.r) || w.r == '\n' || w.r == '\r' {
					w.writer.WriteRune(':')
					w.writer.Write(emojiBuf.Bytes())
					w.writer.WriteRune(w.r)
					break
				} else if w.r == ':' {
					emojiKey := ":" + emojiBuf.String() + ":"
					emojiCode, ok := emoji.EmojiCodeMap[emojiKey]
					if !ok {
						w.writer.WriteRune(':')
						w.writer.WriteString(emojiKey)
						w.writer.WriteRune(':')
						break
					}
					w.writer.WriteString(emojiCode)
					break
				} else {
					emojiBuf.WriteRune(w.r)
				}
			}
		} else {
			w.writer.WriteRune(w.r)
			if w.r == '\n' {
				w.writer.Flush()
			}
		}
	}
}

func main() {
	// This if for checking if the stdin is being piped.
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 { // Stdin through pipe
		writer := newEmojiWriter(true, os.Stdout)
		for {
			_, err := io.Copy(writer, os.Stdin)
			if err == io.EOF {
				return
			} else if err != nil {
				panic(err)
			}
		}
		return
	}

	// This if for executing like watch.

	if len(os.Args) < 2 {
		panic(errors.New("missing arguments"))
	}
	cmd := exec.Command(os.Args[1], os.Args[2:]...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = newEmojiWriter(false, os.Stdout)
	cmd.Stderr = newEmojiWriter(false, os.Stderr)
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
