package dao

import (
	"database/sql"
	"github.com/pkg/errors"
	"homework/02_error/mysql"
)

type UserTest struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

var USER_TEST_TABLE = "user_test"

//解决方案1
func FindUsernameById1(id int) (*UserTest, error) {
	user := &UserTest{}
	query := "SELECT id,username FROM " + USER_TEST_TABLE + " where id=?"
	err := mysql.TestMysql.QueryRow(query, id).Scan(&user.Id, &user.Username)
	if err != nil {
		if err == sql.ErrNoRows {
			return user, nil
		}
		return nil, errors.Wrap(err, "dao:FindUsernameById1 failed")
	}
	return user, nil
}

//解决方案2
func FindUsernameById2(id int) (*UserTest, error) {
	user := &UserTest{}
	var count int
	query := "SELECT COUNT(1) as count FROM " + USER_TEST_TABLE + " where id=?"
	err := mysql.TestMysql.QueryRow(query, id).Scan(&count)
	if err != nil {
		return nil, errors.Wrap(err, "dao:FindUsernameById2 failed")
	}
	//数据不存在,直接返回空数据
	if count == 0 {
		return user, nil
	}
	//存在进行数据查询
	query = "SELECT id,username FROM " + USER_TEST_TABLE + " where id=?"
	err = mysql.TestMysql.QueryRow(query, id).Scan(&user.Id, &user.Username)
	if err != nil {
		return nil, errors.Wrap(err, "dao:FindUsernameById2 failed")
	}
	return user, nil
}
