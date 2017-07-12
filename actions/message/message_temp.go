package message

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"dev.model.360baige.com/models/message"
	"dev.model.360baige.com/models/paginator"
	"dev.model.360baige.com/models/batch"
	"strings"
	"encoding/json"
	"time"
)

type MessageTempAction struct {
}

// 新增
func (*MessageTempAction) Add(args *message.MessageTemp, reply *message.MessageTemp) error {
	o := orm.NewOrm()
	o.Using("message")
	id, err := o.Insert(args)
	if err == nil {
		reply.Id = id
		reply.CreateTime = args.CreateTime
		reply.UpdateTime = args.UpdateTime
		reply.Type = args.Type
		reply.SendTime = args.SendTime
		reply.SendCompanyId = args.SendCompanyId
		reply.SendUserId = args.SendUserId
		reply.SendUserPositionId = args.SendUserPositionId
		reply.SendUserPositionType = args.SendUserPositionType
		reply.SendDest = args.SendDest
		reply.ReceiveTime = args.ReceiveTime
		reply.ReceiveCompanyId = args.ReceiveCompanyId
		reply.ReceiveUserId = args.ReceiveUserId
		reply.ReceiveUserPositionId = args.ReceiveUserPositionId
		reply.ReceiveUserPositionType = args.ReceiveUserPositionType
		reply.ReceiveDest = args.ReceiveDest
		reply.Content = args.Content
		reply.Total = args.Total
		reply.Status = args.Status
		
	}
	return err
}

// 查询 by Id
func (*MessageTempAction) FindById(args *message.MessageTemp, reply *message.MessageTemp) error {
	o := orm.NewOrm()
	o.Using("message")
	reply.Id = args.Id
	err := o.Read(reply)
	return err
}

// 更新 by Id
func (*MessageTempAction) UpdateById(args *message.MessageTemp, reply *message.MessageTemp) error {
	o := orm.NewOrm()
	o.Using("message")
	num, err := o.Update(args)
	if err == nil {
		if num > 0 {
			reply.Id = args.Id
		}
	}
	return err
}

// 1. AddMultiple 增加多个
func (*MessageTempAction) AddMultiple(args []*message.MessageTemp, reply *batch.BackNumm) error {
	o := orm.NewOrm()
	o.Using("message") //查询数据库
	num, err := o.InsertMulti(100, args)
	reply.Num = num
	return err
}

// 2.UpdateByIds 修改多个,默认更改状态为-1，只适合id,更改status,update_time
func (*MessageTempAction) UpdateByIds(args *batch.BatchModify, reply *batch.BackNumm) error {
	o := orm.NewOrm()
	o.Using("message")            //查询数据库
	qs := o.QueryTable("message") //查询表名
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

// 3.查询List （按ID, 按页码）
func (*MessageTempAction) List(args *paginator.Paginator, reply *paginator.Paginator) error {
	o := orm.NewOrm()
	o.Using("message")            //查询数据库
	qs := o.QueryTable("message") //查询表名
	qc := o.QueryTable("message") //查询表名
	filters := args.Filters
	// json str struct
	var items []paginator.PaginatorItem
	jsonErr := json.Unmarshal([]byte(filters), &items)
	if (jsonErr == nil) {
		for _, item := range items {
			if (item.O == "") {
				qs = qs.Filter(item.K, item.V)
				qc = qc.Filter(item.K, item.V)
			} else {
				qc = qc.Filter(item.K+"__"+item.O, item.V)
				qs = qs.Filter(item.K+"__"+item.O, item.V)
			}
		}
	}
	start := 0
	if ((args.Current - 1) > 0) {
		start = (args.Current - 1) * args.PageSize
	}
	if (args.MarkID != 0 && args.Direction != 0) {
		if (args.Direction == -1) {
			qc = qc.Filter("id__gt", args.MarkID)
			qs = qs.Filter("id__gt", args.MarkID)
		} else {
			qc = qc.Filter("id__lt", args.MarkID)
			qs = qs.Filter("id__lt", args.MarkID)
		}
	}
	reply.Total, _ = qc.Count()
	if (args.Sord != "") {
		qs = qs.OrderBy("-" + args.Sord)
	} else {
		qs = qs.OrderBy("-id")
	}
	if (args.PageSize != 0) {
		qs = qs.Limit(args.PageSize, start)
	}
	_, err := qs.Values(&reply.List)
	return err
}