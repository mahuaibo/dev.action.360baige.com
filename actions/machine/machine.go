package machine

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"dev.model.360baige.com/models/machine"
)

type MachineAction struct {
}

// 新增
func (*MachineAction) Add(args *machine.Machine, reply *machine.Machine) error {
	o := orm.NewOrm()
	o.Using("machine")
	id, err := o.Insert(args)
	if err == nil {
		reply.Id = id
		reply.CreateTime = args.CreateTime
		reply.UpdateTime = args.UpdateTime
		reply.CompanyId = args.CompanyId

	}
	return err
}

// 查询 by Id
func (*MachineAction) FindById(args *machine.Machine, reply *machine.Machine) error {
	o := orm.NewOrm()
	o.Using("machine")
	reply.Id = args.Id
	err := o.Read(reply)
	return err
}

// 更新 by Id
func (*MachineAction) UpdateById(args *machine.Machine, reply *machine.Machine) error {
	o := orm.NewOrm()
	o.Using("machine")
	num, err := o.Update(args)
	if err == nil {
		if num > 0 {
			reply.Id = args.Id
		}
	}
	return err
}
