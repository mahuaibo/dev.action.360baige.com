package user

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"dev.model.360baige.com/models/user"
	"dev.model.360baige.com/models/batch"
	"strings"
	"time"
)

// 查询 by username
func (*UserAction) FindByUsername(args *user.User, reply *user.User) error {
	o := orm.NewOrm()
	o.Using("user")
	reply.Username = args.Username
	reply.Status = 0
	err := o.Read(reply, "username","status")
	return err
}

// 查询 by email
func (*UserAction) FindByEmail(args *user.User, reply *user.User) error {
	o := orm.NewOrm()
	o.Using("user")
	reply.Email = args.Email
	reply.Status = 0
	err := o.Read(reply, "email","status")
	return err
}

// 查询 by phone
func (*UserAction) FindByPhone(args *user.User, reply *user.User) error {
	o := orm.NewOrm()
	o.Using("user")
	reply.Phone = args.Phone
	reply.Status = 0
	err := o.Read(reply, "phone","status")
	return err
}
// 查询 by AccessTicket
func (*UserAction) FindByAccessTicket(args *user.User, reply *user.User) error {
	o := orm.NewOrm()
	o.Using("user")
	reply.AccessTicket = args.AccessTicket
	err := o.Read(reply, "AccessTicket")
	return err
}

// 2.UpdateByIds 修改多个,默认更改状态为-1，只适合id,更改status,update_time
func (*UserAction) UpdateByIds(args *batch.BatchModify, reply *batch.BackNumm) error {
	o := orm.NewOrm()
	o.Using("user")            //查询数据库
	qs := o.QueryTable("user") //查询表名
	if (args.UpdateTime == 0) {
		args.UpdateTime = time.Now().UnixNano() / 1e6
	}
	if (args.Status == 0) {
		args.Status = -1
	}
	idsArg := strings.Split(args.Ids, ",")
	qs = qs.Filter("id__in", idsArg)
	num, err := qs.Update(orm.Params{
		"status":      args.Status,
		"update_time": args.UpdateTime,
	})
	reply.Num = num
	return err
}
