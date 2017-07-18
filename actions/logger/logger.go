package logger

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"dev.model.360baige.com/models/paginator"
	"dev.model.360baige.com/models/batch"
	"strings"
	"encoding/json"
	"time"
)

// 2.UpdateByIds 修改多个,默认更改状态为-1，只适合id,更改status,update_time
func (*LoggerAction) UpdateByIds(args *batch.BatchModify, reply *batch.BackNumm) error {
	o := orm.NewOrm()
	o.Using("logger")            //查询数据库
	qs := o.QueryTable("logger") //查询表名
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
func (*LoggerAction) List(args *paginator.Paginator, reply *paginator.Paginator) error {
	o := orm.NewOrm()
	o.Using("logger")            //查询数据库
	qs := o.QueryTable("logger") //查询表名
	qc := o.QueryTable("logger") //查询表名
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
