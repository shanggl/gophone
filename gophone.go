package main

import(
	"net/http"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
//	"os"
	"log"
	phone_ui "phone_ui"
	phone_dao "phone_dao"
	"strconv"
	"strings"

)

var  db * sql.DB

func init(){
// init db
	var err error
	db, err = sql.Open("sqlite3", "./contact.db")
	if err != nil {
		log.Fatal(err)
	}
	

}

// index 
func getIndex(w http.ResponseWriter,r * http.Request){

//获取所有部门
	dList,err:=phone_dao.GetAllDepartment(db)
	if nil!=err{
		log.Fatal("QueryList Error:",err)
	}
//根据部门查询所有人员ID/名称
        	
	for i,v:=range(dList){
		log.Println(v.DeptName,v.DeptId)
		pList,err:=phone_dao.GetPersonListByDeptId(db,	v.DeptId)
		if nil!=err{
			log.Println("error query Person list  by id",err)
		} 
		dList[i].PList=pList

	}

	var m=make(map[string] interface{})

	m["DeptList"]=dList
	
	err=phone_ui.IndexTpl.Execute(w,m)
	if nil!=err{
		log.Println("err Serv Index",err)
	} 

	

}
//query for preson by id

func getPerson(w http.ResponseWriter, r* http.Request){
	r.ParseForm()
	pid,ok:=r.Form["pid"]
	if  ok {
		log.Println("query pid  ",pid)

		id,_:=strconv.ParseInt(pid[0],10,64)

		p,err:=phone_dao.GetPersonById(db,id)
		if nil!=err{
			log.Println("error query Person list  by id",err)
		} 
		var m=make(map[string] interface{})
		m["person"]=p
		
		err=phone_ui.PersonTpl.Execute(w,m)
		if nil!=err{
			log.Println("err query Person",err)
		} 
	
	}else{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("pid not found"))
	}
	

}

//query request

func getQuery(w http.ResponseWriter, r* http.Request){
	r.ParseForm()
	q,ok:=r.Form["q"]
	if  ok {
		log.Println("query value  ",q)
		v:=strings.TrimSpace(q[0])

		pList,err:=phone_dao.GetPersonByQuery(db,v)
		if nil!=err{
			log.Println("error query value",err)
		} 
		var m=make(map[string] interface{})
		m["pList"]=pList
		
		err=phone_ui.QueryResultTpl.Execute(w,m)
		if nil!=err{
			log.Println("err query Person",err)
		} 
	
	}else{
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("value not found"))
	}

}



func main(){
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Println("Begin Phone Book Server")

	http.HandleFunc("/",getIndex)
	http.HandleFunc("/p",getPerson)
	http.HandleFunc("/q",getQuery)

	err:=http.ListenAndServe(":9090",nil)
	if nil!=err{
		log.Fatal("Listen bind Error:",err)
	}

	defer db.Close()

}
