package util

import (
	"fmt"
)

const clearLine = "\r\x1b[K"

type ProgressWriter struct {
	FileName    string
	TotalBytes  uint64
	WritedBytes uint64
}

func (pw *ProgressWriter) Write(b []byte) (int, error) {
	length := len(b)
	pw.WritedBytes += uint64(length)

	updateProgress(pw)

	return length, nil
}

// \r\x1b[K
// https://en.wikipedia.org/wiki/ANSI_escape_code#Control_Sequence_Introducer_commands
// https://ru.wikipedia.org/wiki/%D0%A3%D0%BF%D1%80%D0%B0%D0%B2%D0%BB%D1%8F%D1%8E%D1%89%D0%B8%D0%B5_%D0%BF%D0%BE%D1%81%D0%BB%D0%B5%D0%B4%D0%BE%D0%B2%D0%B0%D1%82%D0%B5%D0%BB%D1%8C%D0%BD%D0%BE%D1%81%D1%82%D0%B8_ANSI#CSI-%D0%BA%D0%BE%D0%B4%D1%8B

func updateProgress(pw *ProgressWriter) {
	fmt.Printf("%vDownloading '%v': %v / %v",
		clearLine,
		pw.FileName,
		FormatBytes(pw.WritedBytes),
		FormatBytes(pw.TotalBytes),
	)
}
