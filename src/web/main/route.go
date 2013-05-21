package main

import (
	"log"
	"net/http"
	"html/template"
)

type User struct {
	UserName string;
}



func loginHandler(w http.ResponseWriter, r *http.Request){
	log.Println("login...");
	t, err := template.ParseFiles("../../../template/html/login.html")
	if (err != nil){
		log.Println(err);
	}
	
	t.Execute(w, nil);
}


func adminHandler(w http.ResponseWriter, r *http.Request) {
	
	// 获取cookies
	/*
	cookie, err := r.Cookie("admin_name");
	if (err != nil || cookie == nil || cookie.Value == ""){
		http.Redirect(w, r, "/login/", http.StatusFound);
		return;
	}
	*/
	//userName := cookie.Value;
	userName := "hello"
	
	t, err := template.ParseFiles("../../../template/html/admin.html");
	if (err != nil){
		log.Println(err);
	}
	
	t.Execute(w, &User{userName});
}

func notFoundHandler(w http.ResponseWriter, r *http.Request){
	if (r.URL.Path == "/"){
		http.Redirect(w, r, "/login/", http.StatusFound);
		return;
	}
	
	t, err := template.ParseFiles("../../../template/html/404.html");
	
	if (err != nil){
		log.Println(err);
	}
	
	t.Execute(w, nil);
}