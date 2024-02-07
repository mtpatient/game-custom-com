// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// VisitDataDao is the data access object for table visit_data.
type VisitDataDao struct {
	table   string           // table is the underlying table name of the DAO.
	group   string           // group is the database configuration group name of current DAO.
	columns VisitDataColumns // columns contains all the column names of Table for convenient usage.
}

// VisitDataColumns defines and stores column names for table visit_data.
type VisitDataColumns struct {
	Id         string //
	Date       string //
	ViewCount  string // 每日访客数
	CreateTime string //
	UpdateTime string //
	DeleteTime string //
}

// visitDataColumns holds the columns for table visit_data.
var visitDataColumns = VisitDataColumns{
	Id:         "id",
	Date:       "date",
	ViewCount:  "view_count",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	DeleteTime: "delete_time",
}

// NewVisitDataDao creates and returns a new DAO object for table data access.
func NewVisitDataDao() *VisitDataDao {
	return &VisitDataDao{
		group:   "default",
		table:   "visit_data",
		columns: visitDataColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *VisitDataDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *VisitDataDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *VisitDataDao) Columns() VisitDataColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *VisitDataDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *VisitDataDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *VisitDataDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
