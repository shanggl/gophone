package phone_dao

import (
	_db "database/sql" 
	"log"
)

//查询所有部门列表

func GetAllDepartment(db * _db.DB) ([] Department,error ){
	rows,err:=db.Query(`select d_id,d_name from department order by 1`)
	if nil!=err{
		log.Fatal(err)
	}

	var ret [] Department
	for rows.Next() {
		var d * Department
		d=new(Department)
		rows.Scan(&d.DeptId, &d.DeptName)
		ret=append(ret,*d)
	}
	rows.Close()
	return ret,err

}


//根据部门查询人员ID/姓名


func GetPersonListByDeptId(db * _db.DB,id int32,)([] Person,error){
	var ret [] Person

	 
	rows, err := db.Query(`select p.p_id,p.p_name,p.p_phone,p.p_tel,
	p.p_mail ,p.p_position,p.t_id,t.t_name,
	p.d_id,d.d_name
	from person p,department d,team t
	where p.t_id=t.t_id
	and p.d_id=d.d_id
	and p.d_id=?`,id)

	if err != nil {
		log.Fatal(err)
	} 
 
	for rows.Next() {
		var p * Person
		p=new(Person)
		rows.Scan(&p.PId,&p.Name,&p.Phone,&p.Tel,&p.Mail,&p.Position,&p.TeamId,&p.TeamName,&p.DeptId,&p.DeptName)
		ret=append(ret,*p)
	}
	rows.Close()
	return ret,err
}


func GetAllPerson(db * _db.DB) ([] Person,error){
	var ret [] Person

	rows,err:=db.Query(`select p.p_id,p.p_name,p.p_phone,p.p_tel,
	p.p_mail ,p.p_position,p.t_id,t.t_name,
	p.d_id,d.d_name
	from person p,department d,team t
	where p.t_id=t.t_id
	and p.d_id=d.d_id`)

	if nil!=err{
		log.Fatal(err)
	}


	for rows.Next() {
		var p * Person
		p=new(Person)
		rows.Scan(&p.PId,&p.Name,&p.Phone,&p.Tel,&p.Mail,&p.Position,&p.TeamId,&p.TeamName,&p.DeptId,&p.DeptName)
		ret=append(ret,*p)
	}
	rows.Close()
	return ret,err
}

//查询某一人
func GetPersonById(db * _db.DB,id) (Person,error){
	var ret Person

	rows,err:=db.Query(`select p.p_id,p.p_name,p.p_phone,p.p_tel,
	p.p_mail ,p.p_position,p.t_id,t.t_name,
	p.d_id,d.d_name
	from person p,department d,team t
	where p.t_id=t.t_id
	and p.d_id=d.d_id
	and p.p_id=?`,id)

	if nil!=err{
		log.Fatal(err)
	} 
	for rows.Next() {
		var p * Person
		p=new(Person)
		rows.Scan(&p.PId,&p.Name,&p.Phone,&p.Tel,&p.Mail,&p.Position,&p.TeamId,&p.TeamName,&p.DeptId,&p.DeptName)
		ret=*p
	}
	rows.Close()
	return ret,err
}


func GetPersonByQuery(db * _db.DB,query string)([] Person,error){
var p [] Person
	var err error
return p,err
}
