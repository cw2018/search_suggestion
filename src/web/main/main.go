// main project main.go
package main

import (
	"log"
	"net/http"
	"com/timeocean/pinyin"
	"com/timeocean/treemap"
	"com/timeocean/suggestion"
	"math/rand"
	"time"
)

func main() {
	log.Println("hello");
	testSuggestion();
	//testPinyin();
	//testRbtree();
}

func testSuggestion(){
	//suggestion.TestJson("../../../resources/items.txt")
	//suggestion.TestLoad("../../../resources/items.txt")
	suggestion, _ := suggestion.NewSuggestion("../resources/wordpinyins.new.txt","../resources/items.txt");
	
	//suggestion.PrintSubOrig("keyword3", "keyword8");
	//suggestion.PrintAbbrSuggestionMap();
	//suggestion.PrintOrigSuggestionMap();
	//suggestion.PrintPinyinSuggestionMap();
	
	items := suggestion.Suggest("zly");
	for _, item := range items{
		log.Println("suggestion: ", item.Keyword, item.Score, item.Info);
	}
}

func testPinyin(){
 	pinyin, err := pinyin.NewPinyin("../../../resources/wordpinyins.new.txt");

	if (err != nil){
		log.Println(err);
		return;
	}
	
	log.Println(pinyin.Get("和"));
	log.Println(pinyin.Get("了"));
	log.Println(pinyin.Get("我"));
	//err = pinyin.Dump("../../../resources/dump.txt");
	if (err != nil){
		log.Println(err);
	}
	return;
}


func webstart(){
	http.Handle("/css/", http.FileServer(http.Dir("../../../template")));
	http.Handle("/js/", http.FileServer(http.Dir("../../../template")));
	http.HandleFunc("/admin/", adminHandler);
	http.HandleFunc("/login/", loginHandler);
	http.HandleFunc("/", notFoundHandler);
	http.ListenAndServe(":8888", nil);
}

func testRbtree(){
	t := time.Now();
	nanosec := t.Nanosecond();
	seed := int64(nanosec);
	rand.Seed(seed);
	key := rand.Int();
	log.Println("kkk: ", key);
	
	tree := treemap.New();
	
	tree.Add(treemap.IntKey(key), "haha");
	tree.Add(treemap.IntKey(rand.Int()), "hello");
	
	log.Println(tree.Find(treemap.IntKey(key)).Value);
	
}
