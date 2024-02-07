// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FollowDao is the data access object for table follow.
type FollowDao struct {
	table   string        // table is the underlying table name of the DAO.
	group   string        // group is the database configuration group name of current DAO.
	columns FollowColumns // columns contains all the column names of Table for convenient usage.
}

// FollowColumns defines and stores column names for table follow.
type FollowColumns struct {
	Id           string //
	UserId       string //
	FollowUserId string //
	CreateTime   string //
	DeleteTime   string //
}

// followColumns holds the columns for table follow.
var followColumns = FollowColumns{
	Id:           "id",
	UserId:       "user_id",
	FollowUserId: "follow_user_id",
	CreateTime:   "create_time",
	DeleteTime:   "delete_time",
}

// NewFollowDao creates and returns a new DAO object for table data access.
func NewFollowDao() *FollowDao {
	return &FollowDao{
		group:   "default",
		table:   "follow",
		columns: followColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *FollowDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *FollowDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *FollowDao) Columns() FollowColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *FollowDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *FollowDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *FollowDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
