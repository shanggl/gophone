package main

import(
	"net/http"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
//	"os"
	"log"
	phone_ui "phone_ui"
	phone_dao "phone_dao"

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
//取所有信息，生成详情页
	allPList,err:=phone_dao.GetAllPerson(db)
	if nil!=err{
		log.Println("error query all Person list ",err)
	} 

	var m=make(map[string] interface{})

	m["allPList"]=allPList
	m["DeptList"]=dList
	
	err=phone_ui.IndexTpl.Execute(w,m)
	if nil!=err{
		log.Println("err Serv Index",err)
	} 

	

}
//query for preson by id

func getQuery(w http.ResponseWriter, r* http.Request){

}



func main(){
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Println("Begin Phone Book Server")

	http.HandleFunc("/p",getIndex)
	http.HandleFunc("/q",getQuery)

	err:=http.ListenAndServe(":9090",nil)
	if nil!=err{
		log.Fatal("Listen bind Error:",err)
	}

	defer db.Close()

}
