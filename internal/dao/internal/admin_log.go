// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// AdminLogDao is the data access object for table admin_log.
type AdminLogDao struct {
	table   string          // table is the underlying table name of the DAO.
	group   string          // group is the database configuration group name of current DAO.
	columns AdminLogColumns // columns contains all the column names of Table for convenient usage.
}

// AdminLogColumns defines and stores column names for table admin_log.
type AdminLogColumns struct {
	Id         string //
	UserId     string //
	Type       string //
	Content    string //
	CreateTime string //
	DeleteTime string //
}

// adminLogColumns holds the columns for table admin_log.
var adminLogColumns = AdminLogColumns{
	Id:         "id",
	UserId:     "user_id",
	Type:       "type",
	Content:    "content",
	CreateTime: "create_time",
	DeleteTime: "delete_time",
}

// NewAdminLogDao creates and returns a new DAO object for table data access.
func NewAdminLogDao() *AdminLogDao {
	return &AdminLogDao{
		group:   "default",
		table:   "admin_log",
		columns: adminLogColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *AdminLogDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *AdminLogDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *AdminLogDao) Columns() AdminLogColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *AdminLogDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *AdminLogDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *AdminLogDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
