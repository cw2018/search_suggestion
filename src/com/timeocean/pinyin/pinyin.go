// pinyin
package pinyin

import(
)

type Pinyin struct {
	pinyinMap map [string] []string;
}

func NewPinyin(pinyinFile string)(pinyin *Pinyin, err error){
	pinyinMap, err := load(pinyinFile);
	if(err != nil){
		return;
	}
	
	pinyin = &Pinyin{pinyinMap};
	return;
}

func (pinyin *Pinyin) Get(word string) (pinyins []string){
	pinyins = pinyin.pinyinMap[word]
	return;
}

func (pinyin *Pinyin) GetFirstByRune(c rune, defaultResult string) string{
	
	result := pinyin.pinyinMap[string(c)];
	if (result == nil || len(result) == 0){
		return defaultResult;
	}
	
	return result[0];
}

func (pinyin *Pinyin) GetFirstRuneByRune(c rune, defaultResult rune) rune{
	s := string(c);
	result := pinyin.pinyinMap[s];
	if (result == nil || len(result) == 0){
		return defaultResult;
	}
	
	for _, fc := range result[0] {
		return fc;
	}
	
	return defaultResult;
}

func (pinyin *Pinyin) Dump(dumpFile string)(err error){
	err = dump(pinyin.pinyinMap, dumpFile);
	return;
}

func (pinyin *Pinyin) WordStr2pinyinStr(wordStr string) string{
	pinyinStr := "";
	
	for _, c := range wordStr {
		defStr := "";
		if (isLetter(c)){
			defStr = string(c)
		}
		pinyinStr += pinyin.GetFirstByRune(c, defStr);
	}
	
	return pinyinStr;
}

func isLetter(c rune) (bool) {
	switch {
	case (c >= 'a') && (c <= 'z'):
		return true;
	case (c >= 'A') && (c <= 'Z'):
		return true;
	case (c >= '0') && (c <= '9'):
		return true;
	}
	
	return false;
}

func (pinyin *Pinyin) WordStr2abbrString(wordStr string) string{
	abbrStr := make([]rune,0);
	
	for _, c := range wordStr {
		defRune := rune(0);
		if (isLetter(c)){
			defRune = c
		}
		
		r := pinyin.GetFirstRuneByRune(c, defRune)
		
		if (r != rune(0)){
			abbrStr = append(abbrStr, r);
		}
	}
	
	return string(abbrStr);
}