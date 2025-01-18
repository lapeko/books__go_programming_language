package main

import (
	"bytes"
	"fmt"
	"github.com/lapeko/books__go_programming_language/chapter-7-interfaces/num"
	"io"
	"log"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
	logFolderName     = "logs"
	logPrefix         = "log"
	logSuffix         = ".txt"
	logFileSize   int = 40
)

var (
	once sync.Once
	mut  sync.Mutex
)

type fileLogWriter struct {
	writer            io.Writer
	currentLogFileNum int
}

func (flw *fileLogWriter) Write(p []byte) (n int, err error) {
	timeLogPrefix := time.Now().Format("[02.01.06 15:04:05] ")
	buf := bytes.NewBufferString(timeLogPrefix)
	buf.Write(p)
	n, err = flw.writer.Write(buf.Bytes())

	mut.Lock()
	defer mut.Unlock()

	writtenBytes, bytesToWrite := 0, len(p)
	var file *os.File
	for writtenBytes < len(p) {
		var e error
		file, e = openForAppend(flw.currentLogFileNum)
		if e != nil {
			log.Printf("open file failure%v\n", err)
			return
		}
		stat, e := file.Stat()
		if e != nil {
			log.Printf("file stat read failure%v\n", err)
			file.Close()
			return
		}

		freeFileSpace := logFileSize - int(stat.Size())
		nextChunkSize := num.Min(freeFileSpace, bytesToWrite)

		if freeFileSpace == nextChunkSize {
			flw.currentLogFileNum++
		}

		_, e = file.Write(p[writtenBytes : writtenBytes+nextChunkSize])
		file.Close()

		writtenBytes += nextChunkSize
		bytesToWrite -= nextChunkSize
	}

	return
}

func NewFileLogWriter(writer io.Writer) io.Writer {
	flw := &fileLogWriter{
		writer:            writer,
		currentLogFileNum: 1,
	}

	once.Do(func() {
		files, err := os.ReadDir(filepath.Join(logFolderName))

		if err != nil {
			if err := os.Mkdir(filepath.Join(logFolderName), 0755); err != nil {
				panic(fmt.Sprintf("Failed to create logs folder: %v", err))
			}
		}

		var nameNums []int
		for _, f := range files {
			if f.IsDir() {
				continue
			}
			name := f.Name()
			if !strings.HasPrefix(name, logPrefix) || !strings.HasSuffix(name, logSuffix) {
				continue
			}
			strNum := strings.TrimSuffix(strings.TrimPrefix(name, logPrefix), logSuffix)
			num, err := strconv.Atoi(strNum)
			if err != nil {
				log.Printf("log file parse error: %v\n", err)
				continue
			}
			nameNums = append(nameNums, num)
		}

		if len(nameNums) == 0 {
			return
		}

		slices.Sort(nameNums)
		flw.currentLogFileNum = nameNums[len(nameNums)-1]
	})

	return flw
}

func getLogFilePath(fileNum int) string {
	fileName := fmt.Sprintf("%s%d%s", logPrefix, fileNum, logSuffix)
	return filepath.Join(logFolderName, fileName)
}

func openForAppend(logFileNum int) (*os.File, error) {
	file, err := os.OpenFile(getLogFilePath(logFileNum), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0755)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func main() {
	withFileLogger := NewFileLogWriter(os.Stdout)
	_, _ = withFileLogger.Write([]byte("Here are we go. This is a long log and should be written into several files and also logs into Stdout\n"))
}
