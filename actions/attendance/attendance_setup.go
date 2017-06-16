package attendance

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"dev.model.360baige.com/models/attendance"
)

type AttendanceSetupAction struct {
}

// 新增
func (*AttendanceSetupAction) Add(args *attendance.AttendanceSetup, reply *attendance.AttendanceSetup) error {
	o := orm.NewOrm()
	o.Using("attendance")
	id, err := o.Insert(args)
	if err == nil {
		reply.Id = id
		reply.CreateTime = args.CreateTime
		reply.UpdateTime = args.UpdateTime
		reply.CompanyId = args.CompanyId
		reply.SigninStime = args.SigninStime
		reply.SigninEtime = args.SigninEtime
		reply.SignoutStime = args.SignoutStime
		reply.SignoutEtime = args.SignoutEtime
		reply.Type = args.Type
		reply.Datatype = args.Datatype
		reply.Name = args.Name
		reply.ShortName = args.ShortName
		reply.Color = args.Color
		reply.Timelength = args.Timelength

	}
	return err
}

// 查询 by Id
func (*AttendanceSetupAction) FindById(args *attendance.AttendanceSetup, reply *attendance.AttendanceSetup) error {
	o := orm.NewOrm()
	o.Using("attendance")
	reply.Id = args.Id
	err := o.Read(reply)
	return err
}

// 更新 by Id
func (*AttendanceSetupAction) UpdateById(args *attendance.AttendanceSetup, reply *attendance.AttendanceSetup) error {
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
