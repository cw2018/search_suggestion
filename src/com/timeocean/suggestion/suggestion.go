package suggestion

import (
	"com/timeocean/pinyin"
	"com/timeocean/treemap"
	"log"
	"unicode"
)

type Item struct {
	Keyword string
	Score   float32
	Info    string
}

type Suggestion struct {
	pinyin              *pinyin.Pinyin
	origSuggestionMap   *treemap.RbTree
	pinyinSuggestionMap *treemap.RbTree
	abbrSuggestionMap   *treemap.RbTree
}

func NewSuggestion(pinyinFile string, suggestionItemFile string) (suggestion *Suggestion, err error) {
	suggestion = &Suggestion{}
	suggestion.pinyin, err = pinyin.NewPinyin(pinyinFile)

	if err != nil {
		return
	}

	items, err := load(suggestionItemFile)
	if err != nil {
		return
	}

	suggestion.buildOrigSuggestionMap(items)
	suggestion.buildPinyinSuggestionMap(items);
	suggestion.buildAbbrSuggestionMap(items);

	return
}

func (s *Suggestion) buildOrigSuggestionMap(items []Item) {
	s.origSuggestionMap = treemap.New()

	for _, item := range items {
		s.origSuggestionMap.Replace(treemap.StrKey(item.Keyword), item)
	}
	return
}

func (s *Suggestion) buildPinyinSuggestionMap(items []Item) {
	s.pinyinSuggestionMap = treemap.New();
	
	for _, item := range items {
		pinyinStr := s.pinyin.WordStr2pinyinStr(item.Keyword);
		s.pinyinSuggestionMap.Replace(treemap.StrKey(pinyinStr), item);
	}
	return;
}

func (s *Suggestion) buildAbbrSuggestionMap(items []Item) {
	s.abbrSuggestionMap = treemap.New()

	for _, item := range items {
		abbrStr := s.pinyin.WordStr2abbrString(item.Keyword);
		s.abbrSuggestionMap.Replace(treemap.StrKey(abbrStr), item)
	}
	return
}


func (s *Suggestion) PrintOrigSuggestionMap() {
	log.Println("print origSuggestionMap...")
	s.origSuggestionMap.Do(printItem)
}

func (s *Suggestion) PrintPinyinSuggestionMap() {
	log.Println("print pinyinSuggestionMap...")
	s.pinyinSuggestionMap.Do(printItem)
}

func (s *Suggestion) PrintAbbrSuggestionMap() {
	log.Println("print abbrSuggestionMap...")
	s.abbrSuggestionMap.Do(printItem)
}


func printItem(k treemap.Key, v interface{}) {
	kw := string(k.(treemap.StrKey))
	item := v.(Item)

	log.Println(kw, ":", item.Keyword, item.Score, item.Info)
}

func (s *Suggestion) PrintSubOrig(fromK string, toK string) {
	inOrder := func(k treemap.Key, v interface{}) {
		kw := string(k.(treemap.StrKey))
		item := v.(Item)

		log.Println(kw, ":", item.Score, item.Info)
	}
	
	s.origSuggestionMap.DoSubTree(treemap.StrKey(fromK), treemap.StrKey(toK), inOrder)
}


func (s *Suggestion) Suggest(input string) (items []Item){
	origInput := input;
	pinyinInput := s.pinyin.WordStr2pinyinStr(input);
	
	fetchItem := func (k treemap.Key, v interface{}) {
		items = append(items, v.(Item))
	}
	
	s.origSuggestionMap.DoSubTree(treemap.StrKey(origInput), treemap.StrKey(origInput + string(unicode.MaxRune)), fetchItem)
	s.pinyinSuggestionMap.DoSubTree(treemap.StrKey(pinyinInput), treemap.StrKey(pinyinInput + string(unicode.MaxRune)), fetchItem)
	s.abbrSuggestionMap.DoSubTree(treemap.StrKey(origInput), treemap.StrKey(origInput + string(unicode.MaxRune)), fetchItem)
	
	return;
}



