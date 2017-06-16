package company

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/astaxie/beego/orm"
	"dev.model.360baige.com/models/company"
)

type CompanyAction struct {
}

// 新增
func (*CompanyAction) Add(args *company.Company, reply *company.Company) error {
	o := orm.NewOrm()
	o.Using("company")
	id, err := o.Insert(args)
	if err == nil {
		reply.Id = id
		reply.CreateTime = args.CreateTime
		reply.UpdateTime = args.UpdateTime
		reply.Type = args.Type
		reply.Level = args.Level
		reply.Logo = args.Logo
		reply.Name = args.Name
		reply.ShortName = args.ShortName
		reply.SubDomain = args.SubDomain
		reply.ProvinceId = args.ProvinceId
		reply.CityId = args.CityId
		reply.DistrictId = args.DistrictId
		reply.Address = args.Address
		reply.PositionX = args.PositionX
		reply.PositionY = args.PositionY
		reply.Status = args.Status
		reply.Remark = args.Remark
		reply.Brief = args.Brief

	}
	return err
}

// 查询 by Id
func (*CompanyAction) FindById(args *company.Company, reply *company.Company) error {
	o := orm.NewOrm()
	o.Using("company")
	reply.Id = args.Id
	err := o.Read(reply)
	return err
}

// 更新 by Id
func (*CompanyAction) UpdateById(args *company.Company, reply *company.Company) error {
	o := orm.NewOrm()
	o.Using("company")
	num, err := o.Update(args)
	if err == nil {
		if num > 0 {
			reply.Id = args.Id
		}
	}
	return err
}
