package main

import (
	"homework/02_error/dao"
	"log"
)

/*
问题:1. 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，
应该怎么做请写出代码？

答: 不应该抛给上层。
原因: 因为sql.ErrNoRows 的详情是"sql: no rows in result set",表示所要查询的具体数据不存在,这种可能是一种业务类型的错误，
也可能对于某些需求来说不算一个错误。所以可以直接再dao层处理掉，如果查询遇到这个错误直接返回空数据即可，就不用再进行错误抛给上层
了。上层可以再根据返回的数据可以做一些逻辑。
*/

func main() {
	//方案一
	//当error为sql.ErrNoRows，直接返回空的数据
	user1, err := dao.FindUsernameById1(1)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(user1)
	//方案二
	//避免出现error为sql.ErrNoRows,先查询库中是否有对应的数据
	//如果有再继续查询，没有的话直接返回空数据
	user2, err := dao.FindUsernameById2(2)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(user2)
}
