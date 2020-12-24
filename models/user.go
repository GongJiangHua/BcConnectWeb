package models

import "BitConnectWeb/db_mysql"

type User struct {
	Id int `form:"id"`
	Phone string `form:"phone"`
	Password string `form:"password"`
}

func (u User)SaveUser() (int64,error) {
	result,err := db_mysql.DB.Exec("insert into user_mes (telephone,password)values (?,?)",u.Phone,u.Password)
	if err != nil {
		return -1,err
	}
	rows,err := result.RowsAffected()
	if err != nil {
		return -1,err
	}
	return rows,nil
}