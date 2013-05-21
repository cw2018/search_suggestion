package suggestion

import (
	"bufio"
	"encoding/json"
	"io"
	"log"
	"os"
	"strconv"
)

func load(infile string) (items []Item, err error) {
	log.Println("loading file : " + infile)

	file, err := os.Open(infile)
	if err != nil {
		log.Println("Failed to open the input file. error: ", err)
		return
	}
	defer file.Close()

	items = make([]Item, 0, 512)

	br := bufio.NewReader(file)

	for {
		line, isPrefix, err1 := br.ReadLine()

		if err1 != nil {
			if err1 != io.EOF {
				err = err1
			}
			return
		}

		if isPrefix {
			log.Println("line read from file is too long")
			continue
		}

		item := &Item{}

		if nil == json.Unmarshal(line, item) {
			items = append(items, *item)
		}
	}

	return
}

func TestLoad(infile string) {
	items, _ := load(infile)
	for _, item := range items {
		log.Println(item.Keyword)
		log.Println(item.Info)
	}
}

func TestJson(dumpfile string) {
	file, err := os.Create(dumpfile)

	if err != nil {
		log.Println(err)
		return
	}

	defer file.Close()

	bw := bufio.NewWriter(file)

	for i := 0; i < 10; i++ {

		item := Item{
			Keyword: "keyword" + strconv.Itoa(i),
			Score:   float32(i),
			Info:    "{\"hitsCount\":32,\"memo\":\"aaa\"}",
		}

		log.Println(item.Info)

		b, err := json.Marshal(item)
		if err != nil {
			log.Println(err)
			return
		}

		bw.Write(b)
		bw.WriteString("\n")
	}

	bw.Flush()

}
