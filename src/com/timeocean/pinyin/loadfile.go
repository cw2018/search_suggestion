// loadfile
package pinyin

import (
	"log"
	"os"
	"bufio"
	"io"
	"strings"
)

func load(infile string)(pinyinsMap map[string] []string, err error) {
	log.Println("loading file : " + infile);
	
	file, err := os.Open(infile);
	if (err != nil) {
		log.Println("Failed to open the input file. error: ", err);
		return;
	}
	defer file.Close();
	
	pinyinsMap = make(map[string] []string);
	
	br := bufio.NewReader(file);
	
	for {
		line, isPrefix, err1 := br.ReadLine();
		
		if (err1 != nil) {
			if (err1 != io.EOF){
				err = err1;
			}
			return;
		}

		
		if (isPrefix){
			log.Println("line read from file is too long : ", line);
			continue;
		}
						
		str := string(line);

		strs := strings.Split(str, " ");

		if (len(strs) <= 0){
			continue;
		}
		
		word := strs[0];
		
		word = strings.TrimSpace(word);
		
		if (len(word) <= 0){
			continue;
		}
		
		pinyins := make([]string, 0);
		
		for i := 1; i < len(strs); i++{
			pinyin := strs[i];
			pinyin = strings.TrimSpace(pinyin);
			if (len(pinyin) > 0){
				pinyins = append(pinyins, pinyin);
			}
		}
		
		if (len(pinyins) > 0){
			pinyinsMap[word] = pinyins;
		}
	}

	return;
}
