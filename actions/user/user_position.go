package user

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"dev.model.360baige.com/models/user"
)

type UserPositionAction struct {
}

// 新增
func (*UserPositionAction) Add(args *user.UserPosition, reply *user.UserPosition) error {
	o := orm.NewOrm()
	o.Using("user")
	id, err := o.Insert(args)
	if err == nil {
		reply.Id = id
		reply.CreateTime = args.CreateTime
		reply.UpdateTime = args.UpdateTime
		reply.CompanyId = args.CompanyId
		reply.UserId = args.UserId
		reply.PersonId = args.PersonId
		reply.AccessToken = args.AccessToken
		reply.ExpireIn = args.ExpireIn

	}
	return err
}

// 查询 by Id
func (*UserPositionAction) FindById(args *user.UserPosition, reply *user.UserPosition) error {
	o := orm.NewOrm()
	o.Using("user")
	reply.Id = args.Id
	err := o.Read(reply)
	return err
}

// 更新 by Id
func (*UserPositionAction) UpdateById(args *user.UserPosition, reply *user.UserPosition) error {
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
