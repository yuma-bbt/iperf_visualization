package main

import (
        "encoding/csv"
	"strings"
        "fmt"
        "io"
        "os"
)

func main() {

        var fp *os.File
	var time []string
	var bandwidth []string
	var i=0 int
	
        if len(os.Args) < 2 {
                fp = os.Stdin
        } else {
                var err error
                fp, err = os.Open(os.Args[1])
                if err != nil {
                        panic(err)
                }
                defer fp.Close()
        }

        reader := csv.NewReader(fp)
        reader.Comma = '\n'
        reader.LazyQuotes = true // ダブルクオートを厳密にチェックしない！
        for {
                record, err := reader.Read()
                if err == io.EOF {
                        break
                } else if err != nil {
                        panic(err)
                }
                fmt.Println(record)
		hoge[i] = strings.Split(record,"\t")
		
        }
}

