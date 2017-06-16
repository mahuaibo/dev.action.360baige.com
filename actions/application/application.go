package application

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"dev.model.360baige.com/models/application"
)

type ApplicationAction struct {
}

// 新增
func (*ApplicationAction) Add(args *application.Application, reply *application.Application) error {
	o := orm.NewOrm()
	o.Using("application")
	id, err := o.Insert(args)
	if err == nil {
		reply.Id = id
		reply.CreateTime = args.CreateTime
		reply.UpdateTime = args.UpdateTime
		reply.CompanyId = args.CompanyId
		reply.UserId = args.UserId
		reply.ApplicationTplId = args.ApplicationTplId
		reply.Name = args.Name
		reply.Type = args.Type
		reply.Status = args.Status

	}
	return err
}

// 查询 by Id
func (*ApplicationAction) FindById(args *application.Application, reply *application.Application) error {
	o := orm.NewOrm()
	o.Using("application")
	reply.Id = args.Id
	err := o.Read(reply)
	return err
}

// 更新 by Id
func (*ApplicationAction) UpdateById(args *application.Application, reply *application.Application) error {
	o := orm.NewOrm()
	o.Using("application")
	num, err := o.Update(args)
	if err == nil {
		if num > 0 {
			reply.Id = args.Id
		}
	}
	return err
}
