package user

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"dev.model.360baige.com/models/user"
)

type UserAction struct {
}

// 新增
func (*UserAction) Add(args *user.User, reply *user.User) error {
	o := orm.NewOrm()
	o.Using("user")
	id, err := o.Insert(args)
	if err == nil {
		reply.Id = id
		reply.CreateTime = args.CreateTime
		reply.UpdateTime = args.UpdateTime
		reply.Username = args.Username
		reply.Password = args.Password
		reply.Email = args.Email
		reply.Phone = args.Phone
		reply.Status = args.Status
		reply.Code = args.Code
		reply.CodeTime = args.CodeTime
		reply.AccessToken = args.AccessToken
		reply.ExpireIn = args.ExpireIn

	}
	return err
}

// 查询 by Id
func (*UserAction) FindById(args *user.User, reply *user.User) error {
	o := orm.NewOrm()
	o.Using("user")
	reply.Id = args.Id
	err := o.Read(reply)
	return err
}

// 更新 by Id
func (*UserAction) UpdateById(args *user.User, reply *user.User) error {
	o := orm.NewOrm()
	o.Using("user")
	num, err := o.Update(args)
	if err == nil {
		if num > 0 {
			reply.Id = args.Id
		}
	}
	return err
}
