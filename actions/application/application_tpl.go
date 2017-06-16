package application

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"dev.model.360baige.com/models/application"
)

type ApplicationTplAction struct {
}

// 新增
func (*ApplicationTplAction) Add(args *application.ApplicationTpl, reply *application.ApplicationTpl) error {
	o := orm.NewOrm()
	o.Using("application")
	id, err := o.Insert(args)
	if err == nil {
		reply.Id = id
		reply.CreateTime = args.CreateTime
		reply.UpdateTime = args.UpdateTime
		reply.CompanyId = args.CompanyId
		reply.UserId = args.UserId
		reply.Name = args.Name
		reply.Image = args.Image
		reply.Site = args.Site
		reply.Type = args.Type
		reply.Desc = args.Desc
		reply.Status = args.Status
		reply.Price = args.Price

	}
	return err
}

// 查询 by Id
func (*ApplicationTplAction) FindById(args *application.ApplicationTpl, reply *application.ApplicationTpl) error {
	o := orm.NewOrm()
	o.Using("application")
	reply.Id = args.Id
	err := o.Read(reply)
	return err
}

// 更新 by Id
func (*ApplicationTplAction) UpdateById(args *application.ApplicationTpl, reply *application.ApplicationTpl) error {
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
