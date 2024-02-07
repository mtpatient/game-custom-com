// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FriendLinkDao is the data access object for table friend_link.
type FriendLinkDao struct {
	table   string            // table is the underlying table name of the DAO.
	group   string            // group is the database configuration group name of current DAO.
	columns FriendLinkColumns // columns contains all the column names of Table for convenient usage.
}

// FriendLinkColumns defines and stores column names for table friend_link.
type FriendLinkColumns struct {
	Id         string //
	Name       string // 友链名称
	Url        string // 友链地址
	CreateTime string // 创建时间
	UpdateTime string //
	DeleteTime string //
}

// friendLinkColumns holds the columns for table friend_link.
var friendLinkColumns = FriendLinkColumns{
	Id:         "id",
	Name:       "name",
	Url:        "url",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	DeleteTime: "delete_time",
}

// NewFriendLinkDao creates and returns a new DAO object for table data access.
func NewFriendLinkDao() *FriendLinkDao {
	return &FriendLinkDao{
		group:   "default",
		table:   "friend_link",
		columns: friendLinkColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *FriendLinkDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *FriendLinkDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *FriendLinkDao) Columns() FriendLinkColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *FriendLinkDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *FriendLinkDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *FriendLinkDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
