package search

import (
	"bufio"
	"context"
	"io"
	"log"
	"os"
	"strings"
	"sync"
)

// Result describes one search result
type Result struct {
	// Phrase the phrase you were looking for
	Phrase string
	// Line the line in which the phrase was found
	Line string
	// LineNum line number in which the phrase was found (starting from 1)
	LineNum int64
	//ColNum position number in which the phrase was found (starting from 1)
	ColNum int64
}

func All(ctx context.Context, phrase string, files []string) <-chan []Result {
	ch := make(chan []Result, len(files))
	defer close(ch)

	wg := sync.WaitGroup{}

	for _, file := range files {
		wg.Add(1)
		go func(ch chan []Result, phrase string, filepath string) {
			defer wg.Done()
			resAll := []Result{}
			counter := int64(0)

			file, err := os.Open(filepath)
			if err != nil {
				log.Print(err)
			}

			defer func() {
				err := file.Close()
				if err != nil {
					log.Print(err)
				}
			}()

			reader := bufio.NewReader(file)
			for {
				strFile, err := reader.ReadString('\n')
				counter++

				if strings.Contains(strFile, phrase) {

					resOne := Result{
						Phrase:  phrase,
						Line:    strings.TrimSuffix(strFile, "\n"),
						LineNum: counter,
						ColNum:  int64(strings.Index(strFile, phrase)) + 1,
					}

					resAll = append(resAll, resOne)

				}

				if err == io.EOF {
					if len(resAll) == 0 {
						break
					}
					ch <- resAll
					break
				}
			}

		}(ch, phrase, file)
	}
	wg.Wait()

	return ch
}