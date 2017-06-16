package logger

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"dev.model.360baige.com/models/logger"
)

type LoggerAction struct {
}

// 新增
func (*LoggerAction) Add(args *logger.Logger, reply *logger.Logger) error {
	o := orm.NewOrm()
	o.Using("logger")
	id, err := o.Insert(args)
	if err == nil {
		reply.Id = id
		reply.CreateTime = args.CreateTime
		reply.CompanyId = args.CompanyId
		reply.OwnerId = args.OwnerId
		reply.Remark = args.Remark
		reply.Content = args.Content
		reply.Type = args.Type

	}
	return err
}

// 查询 by Id
func (*LoggerAction) FindById(args *logger.Logger, reply *logger.Logger) error {
	o := orm.NewOrm()
	o.Using("logger")
	reply.Id = args.Id
	err := o.Read(reply)
	return err
}

// 更新 by Id
func (*LoggerAction) UpdateById(args *logger.Logger, reply *logger.Logger) error {
	o := orm.NewOrm()
	o.Using("logger")
	num, err := o.Update(args)
	if err == nil {
		if num > 0 {
			reply.Id = args.Id
		}
	}
	return err
}
