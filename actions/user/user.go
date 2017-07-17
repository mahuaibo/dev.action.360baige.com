package user
//
//import (
//	_ "github.com/go-sql-driver/mysql"
//	"github.com/astaxie/beego/orm"
//	"dev.model.360baige.com/models/user"
//	"dev.model.360baige.com/models/paginator"
//	"dev.model.360baige.com/models/batch"
//	"strings"
//	"encoding/json"
//	"time"
//)
//
//type UserAction struct {
//}
//
//// 新增
//func (*UserAction) Add(args *user.User, reply *user.User) error {
//	o := orm.NewOrm()
//	o.Using("user")
//	id, err := o.Insert(args)
//	if err == nil {
//		reply.Id = id
//		reply.CreateTime = args.CreateTime
//		reply.UpdateTime = args.UpdateTime
//		reply.Username = args.Username
//		reply.Password = args.Password
//		reply.Email = args.Email
//		reply.Phone = args.Phone
//		reply.Status = args.Status
//		reply.Code = args.Code
//		reply.CodeTime = args.CodeTime
//		reply.AccessTicket = args.AccessTicket
//		reply.ExpireIn = args.ExpireIn
//
//	}
//	return err
//}
//
//// 查询 by Id
//func (*UserAction) FindById(args *user.User, reply *user.User) error {
//	o := orm.NewOrm()
//	o.Using("user")
//	reply.Id = args.Id
//	reply.Status = 0
//	err := o.Read(reply,"id","status")
//	return err
//}
//
//// 查询 by username
//func (*UserAction) FindByUsername(args *user.User, reply *user.User) error {
//	o := orm.NewOrm()
//	o.Using("user")
//	reply.Username = args.Username
//	reply.Status = 0
//	err := o.Read(reply, "username","status")
//	return err
//}
//
//// 查询 by email
//func (*UserAction) FindByEmail(args *user.User, reply *user.User) error {
//	o := orm.NewOrm()
//	o.Using("user")
//	reply.Email = args.Email
//	reply.Status = 0
//	err := o.Read(reply, "email","status")
//	return err
//}
//
//// 查询 by phone
//func (*UserAction) FindByPhone(args *user.User, reply *user.User) error {
//	o := orm.NewOrm()
//	o.Using("user")
//	reply.Phone = args.Phone
//	reply.Status = 0
//	err := o.Read(reply, "phone","status")
//	return err
//}
//// 查询 by AccessTicket
//func (*UserAction) FindByAccessTicket(args *user.User, reply *user.User) error {
//	o := orm.NewOrm()
//	o.Using("user")
//	reply.AccessTicket = args.AccessTicket
//	err := o.Read(reply, "AccessTicket")
//	return err
//}
//
//// 更新 by Id
//func (*UserAction) UpdateById(args *user.User, reply *user.User) error {
//	o := orm.NewOrm()
//	o.Using("user")
//	num, err := o.Update(args)
//	if err == nil {
//		if num > 0 {
//			reply.Id = args.Id
//		}
//	}
//	return err
//}
//
//// 1. AddMultiple 增加多个
//func (*UserAction) AddMultiple(args []*user.User, reply *batch.BackNumm) error {
//	o := orm.NewOrm()
//	o.Using("user") //查询数据库
//	num, err := o.InsertMulti(100, args)
//	reply.Num = num
//	return err
//}
//
//// 2.UpdateByIds 修改多个,默认更改状态为-1，只适合id,更改status,update_time
//func (*UserAction) UpdateByIds(args *batch.BatchModify, reply *batch.BackNumm) error {
//	o := orm.NewOrm()
//	o.Using("user")            //查询数据库
//	qs := o.QueryTable("user") //查询表名
//	if (args.UpdateTime == 0) {
//		args.UpdateTime = time.Now().UnixNano() / 1e6
//	}
//	if (args.Status == 0) {
//		args.Status = -1
//	}
//	idsArg := strings.Split(args.Ids, ",")
//	qs = qs.Filter("id__in", idsArg)
//	num, err := qs.Update(orm.Params{
//		"status":      args.Status,
//		"update_time": args.UpdateTime,
//	})
//	reply.Num = num
//	return err
//}
//
//// 3.查询List （按ID, 按页码）
//func (*UserAction) List(args *paginator.Paginator, reply *paginator.Paginator) error {
//	o := orm.NewOrm()
//	o.Using("user")            //查询数据库
//	qs := o.QueryTable("user") //查询表名
//	qc := o.QueryTable("user") //查询表名
//	filters := args.Filters
//	// json str struct
//	var items []paginator.PaginatorItem
//	jsonErr := json.Unmarshal([]byte(filters), &items)
//	if (jsonErr == nil) {
//		for _, item := range items {
//			if (item.O == "") {
//				qs = qs.Filter(item.K, item.V)
//				qc = qc.Filter(item.K, item.V)
//			} else {
//				qc = qc.Filter(item.K+"__"+item.O, item.V)
//				qs = qs.Filter(item.K+"__"+item.O, item.V)
//			}
//		}
//	}
//	start := 0
//	if ((args.Current - 1) > 0) {
//		start = (args.Current - 1) * args.PageSize
//	}
//	if (args.MarkID != 0 && args.Direction != 0) {
//		if (args.Direction == -1) {
//			qc = qc.Filter("id__gt", args.MarkID)
//			qs = qs.Filter("id__gt", args.MarkID)
//		} else {
//			qc = qc.Filter("id__lt", args.MarkID)
//			qs = qs.Filter("id__lt", args.MarkID)
//		}
//	}
//	reply.Total, _ = qc.Count()
//	if (args.Sord != "") {
//		qs = qs.OrderBy("-" + args.Sord)
//	} else {
//		qs = qs.OrderBy("-id")
//	}
//	if (args.PageSize != 0) {
//		qs = qs.Limit(args.PageSize, start)
//	}
//	_, err := qs.Values(&reply.List)
//	return err
//}