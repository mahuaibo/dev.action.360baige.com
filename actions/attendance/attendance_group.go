package attendance

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"dev.model.360baige.com/models/attendance"
)

type AttendanceGroupAction struct {
}

// 新增
func (*AttendanceGroupAction) Add(args *attendance.AttendanceGroup, reply *attendance.AttendanceGroup) error {
	o := orm.NewOrm()
	o.Using("attendance")
	id, err := o.Insert(args)
	if err == nil {
		reply.Id = id
		reply.CreateTime = args.CreateTime
		reply.UpdateTime = args.UpdateTime
		reply.CompanyId = args.CompanyId
		reply.AttendanceSetupId = args.AttendanceSetupId
		reply.PersonId = args.PersonId
		reply.Name = args.Name
		reply.UserIds = args.UserIds
		reply.Status = args.Status

	}
	return err
}

// 查询 by Id
func (*AttendanceGroupAction) FindById(args *attendance.AttendanceGroup, reply *attendance.AttendanceGroup) error {
	o := orm.NewOrm()
	o.Using("attendance")
	reply.Id = args.Id
	err := o.Read(reply)
	return err
}

// 更新 by Id
func (*AttendanceGroupAction) UpdateById(args *attendance.AttendanceGroup, reply *attendance.AttendanceGroup) error {
	o := orm.NewOrm()
	o.Using("attendance")
	num, err := o.Update(args)
	if err == nil {
		if num > 0 {
			reply.Id = args.Id
		}
	}
	return err
}
