package main

import (
	"fmt"
	"io"
	"os"
)

func e[T any](t T, e error) T {
	if e != nil {
		fmt.Fprintln(os.Stderr, e)
		os.Exit(1)
	}
	return t
}

// makeBackup creates a backup file by copying
// src to a new file with a "backup" extension.
func makeBackup(src *os.File) {
	bak := e(os.Create(src.Name() + ".backup"))
	defer bak.Close()

	e(io.Copy(bak, src))
}

// writeAt writes data to f at the given offset.
func writeAt(f *os.File, offset int, data []byte) {
	fmt.Printf("Writing at 0x%04x\n", offset)
	e(f.Seek(int64(offset), 0))
	e(f.Write(data))
}

// clearAt clears sections of f by writing zero bytes.
//
// Each pair of offsets specifies the start offset and
// length of a section to be cleared.
func clearAt(f *os.File, offsets [][2]int) {
	for _, o := range offsets {
		writeAt(f, o[0], make([]byte, o[1]))
	}
}

func main() {
	if len(os.Args) < 2 {
		os.Args = append(os.Args, "user32.dll.mui")
	}

	f := e(os.OpenFile(os.Args[1], os.O_RDWR, 0))
	defer f.Close()
	makeBackup(f)

	writeAt(f, 0x110, []byte{0xba, 0xce})
	clearAt(f, [][2]int{
		{0x154e, 28},
		{0x156c, 24},
		{0x1586, 10},
		{0x1592, 12},
		{0x16d0, 88},
		{0x193e, 32},
		{0x1960, 8},
	})
	fmt.Println("\nDone")
}
