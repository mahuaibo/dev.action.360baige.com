package personnel

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"dev.model.360baige.com/models/personnel"
)

type PersonRelationAction struct {
}

// 新增
func (*PersonRelationAction) Add(args *personnel.PersonRelation, reply *personnel.PersonRelation) error {
	o := orm.NewOrm()
	o.Using("personnel")
	id, err := o.Insert(args)
	if err == nil {
		reply.Id = id
		reply.CreateTime = args.CreateTime
		reply.UpdateTime = args.UpdateTime
		reply.CompanyId = args.CompanyId
		reply.AssociationId = args.AssociationId
		reply.AssociatedId = args.AssociatedId
		reply.Type = args.Type
		reply.Status = args.Status

	}
	return err
}

// 查询 by Id
func (*PersonRelationAction) FindById(args *personnel.PersonRelation, reply *personnel.PersonRelation) error {
	o := orm.NewOrm()
	o.Using("personnel")
	reply.Id = args.Id
	err := o.Read(reply)
	return err
}

// 更新 by Id
func (*PersonRelationAction) UpdateById(args *personnel.PersonRelation, reply *personnel.PersonRelation) error {
	o := orm.NewOrm()
	o.Using("personnel")
	num, err := o.Update(args)
	if err == nil {
		if num > 0 {
			reply.Id = args.Id
		}
	}
	return err
}
