package phone_dao

type Person struct{
	PId  		int32
	Name 		string
	Phone 		string
	Tel   		string
	Mail  		string
	Position 	string
	TeamId   	int32
	TeamName 	string
	DeptId		int32
	DeptName	string
}

/*
select p.p_id,p.p_name,p.p_phone,p.p_tel,
p.p_mail ,p.p_position,p.t_id,t.t_name,
p.d_id,d.d_name
from person p,department d,team t
where p.t_id=t.t_id
and p.d_id=d.d_id

*/
