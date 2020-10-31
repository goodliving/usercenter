package model

import "errors"

type Users struct {
	ID       int    	`json:"user_id"`
	Username string 	`json:"username"`
	DisplayName string 	`json:"display_name"`
	DeptId int	`json:"dept_id"`
	DeptName string `json:"dept_name"`
	RoleId int 	`json:"role_id"`
	RoleName string `json:"role_name"`
}

func CheckAuth(username, password string) (*Users, error) {
	var u Users

	err := db.Where("username = ? and password = ?", username, password).Find(&u).Error

	if err != nil {
		return nil, err
	}

	if u.ID > 0 {
		return &u, nil
	}

	return nil, errors.New("未知错误，请检查")
}


func UserInsert(u Users) error {
	err := db.Create(&u).Error

	if err != nil {
		return err
	}

	return nil
}