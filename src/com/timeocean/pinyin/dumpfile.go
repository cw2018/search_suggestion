package pinyin

import(
	"os"
	"log"
	"bufio"
)

func dump(pinyins map[string] []string, dumpfile string)(err error) {
	file, err := os.Create(dumpfile);
	
	if (err != nil){
		log.Panicln(err);
		return;
	}
	
	defer file.Close();
	
	bw := bufio.NewWriter(file);
	
	for key, value := range pinyins {
		bw.WriteString(key);
		bw.WriteString("@");
		for _, pinyin := range value {
			bw.WriteString(pinyin);
			bw.WriteString("#");
		}
		bw.WriteString("\n");
	}
	
	bw.Flush();
	
	return;
}