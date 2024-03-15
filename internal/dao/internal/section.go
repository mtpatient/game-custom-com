// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// SectionDao is the data access object for table section.
type SectionDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns SectionColumns // columns contains all the column names of Table for convenient usage.
}

// SectionColumns defines and stores column names for table section.
type SectionColumns struct {
	Id         string //
	Name       string //
	Dc         string // 描述
	Role       string // 该板块所属用户，0：普通用户，1：仅管理员用户
	Icon       string // url链接
	CreateTime string //
	UpdateTime string //
	DeleteTime string //
}

// sectionColumns holds the columns for table section.
var sectionColumns = SectionColumns{
	Id:         "id",
	Name:       "name",
	Dc:         "dc",
	Role:       "role",
	Icon:       "icon",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	DeleteTime: "delete_time",
}

// NewSectionDao creates and returns a new DAO object for table data access.
func NewSectionDao() *SectionDao {
	return &SectionDao{
		group:   "default",
		table:   "section",
		columns: sectionColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *SectionDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *SectionDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *SectionDao) Columns() SectionColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *SectionDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *SectionDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *SectionDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
