package user

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"dev.model.360baige.com/models/user"
	"dev.model.360baige.com/models/batch"
	"dev.model.360baige.com/http/window"
	"strings"
	"time"
	"fmt"
)

type UserPositionAction struct {
}

// 新增
func (*UserPositionAction) Add(args *user.UserPosition, reply *user.UserPosition) error {
	o := orm.NewOrm()
	o.Using("user")
	id, err := o.Insert(args)
	if err == nil {
		reply.Id = id
		reply.CreateTime = args.CreateTime
		reply.UpdateTime = args.UpdateTime
		reply.Type = args.Type
		reply.CompanyId = args.CompanyId
		reply.UserId = args.UserId
		reply.PersonId = args.PersonId
		reply.AccessToken = args.AccessToken
		reply.ExpireIn = args.ExpireIn
		reply.Status = args.Status

	}
	return err
}

// 查询 by Id
func (*UserPositionAction) FindById(args *user.UserPosition, reply *user.UserPosition) error {
	o := orm.NewOrm()
	o.Using("user")
	reply.Id = args.Id
	err := o.Read(reply)
	return err
}

// 查询 by UserId
func (*UserPositionAction) ListAll(args *window.UserPositionPaginator, reply *window.UserPositionPaginator) error {
	o := orm.NewOrm()
	o.Using("user")
	if args.PageSize == 0 {
		args.PageSize = -1
	}
	num, err := o.QueryTable("user_position").SetCond(args.Cond).OrderBy(args.OrderBy...).Limit(args.PageSize).All(&reply.List, args.Cols...)
	fmt.Println(num)
	reply.Total = num
	return err
}

// 更新 by Id
func (*UserPositionAction) UpdateById(args *user.UserPosition, reply *user.UserPosition) error {
	o := orm.NewOrm()
	o.Using("user")
	num, err := o.Update(args)
	if err == nil {
		if num > 0 {
			reply.Id = args.Id
		}
	}
	return err
}

// 1. AddMultiple 增加多个
func (*UserPositionAction) AddMultiple(args []*user.UserPosition, reply *batch.BackNumm) error {
	o := orm.NewOrm()
	o.Using("user") //查询数据库
	num, err := o.InsertMulti(100, args)
	reply.Num = num
	return err
}

// 2.UpdateByIds 修改多个,默认更改状态为-1，只适合id,更改status,update_time
func (*UserPositionAction) UpdateByIds(args *batch.BatchModify, reply *batch.BackNumm) error {
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

//获取身份,涉及分库
func (*UserPositionAction) PositionListAllByUserId(args *window.UserPositionListPaginator, reply *window.UserPositionListPaginator) error {
	o := orm.NewOrm()
	o.Using("user")
	qb, _ := orm.NewQueryBuilder("mysql")
	if args.PageSize == 0 {
		args.PageSize = -1
	}
	qb.Select(args.Cols...).
		From("user_position").
		InnerJoin("db_company.company").On("user_position.company_id = db_company.company.id").
		Where("user_position.status>-1 ").And("user_position.user_id = ?").
		OrderBy(args.OrderBy...).Desc()
	// 导出 SQL 语句
	sql := qb.String()
	num, err := o.Raw(sql,args.Cond).QueryRows(&reply.List)
	reply.Total = num
	return err
}
