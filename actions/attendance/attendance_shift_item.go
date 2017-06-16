package attendance

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"dev.model.360baige.com/models/attendance"
)

type AttendanceShiftItemAction struct {
}

// 新增
func (*AttendanceShiftItemAction) Add(args *attendance.AttendanceShiftItem, reply *attendance.AttendanceShiftItem) error {
	o := orm.NewOrm()
	o.Using("attendance")
	id, err := o.Insert(args)
	if err == nil {
		reply.Id = id
		reply.CreateTime = args.CreateTime
		reply.UpdateTime = args.UpdateTime
		reply.CompanyId = args.CompanyId
		reply.AttendanceShiftId = args.AttendanceShiftId
		reply.Pos = args.Pos
		reply.SignIn = args.SignIn
		reply.SignInStart = args.SignInStart
		reply.SignInEnd = args.SignInEnd
		reply.SignBack = args.SignBack
		reply.SignBackStart = args.SignBackStart
		reply.SignBackEnd = args.SignBackEnd
		reply.Type = args.Type
		reply.SignInErrata = args.SignInErrata
		reply.SignBackErrata = args.SignBackErrata

	}
	return err
}

// 查询 by Id
func (*AttendanceShiftItemAction) FindById(args *attendance.AttendanceShiftItem, reply *attendance.AttendanceShiftItem) error {
	o := orm.NewOrm()
	o.Using("attendance")
	reply.Id = args.Id
	err := o.Read(reply)
	return err
}

// 更新 by Id
func (*AttendanceShiftItemAction) UpdateById(args *attendance.AttendanceShiftItem, reply *attendance.AttendanceShiftItem) error {
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
