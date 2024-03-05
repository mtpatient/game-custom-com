// ==========================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// UserDao is the data access object for table user.
type UserDao struct {
	table   string      // table is the underlying table name of the DAO.
	group   string      // group is the database configuration group name of current DAO.
	columns UserColumns // columns contains all the column names of Table for convenient usage.
}

// UserColumns defines and stores column names for table user.
type UserColumns struct {
	Id          string // 用户id，唯一标识
	Username    string // 用户名
	Password    string // 密码
	Email       string // 邮箱，可通过邮箱找回密码
	Avatar      string // 头像id
	Sex         string // 0：女，1：男；2：保密
	Signature   string // 个性签名
	Role        string // 管理员：1，,普通用户：0
	Status      string // 用户所处状态，0为正常，1为被封禁
	FansCount   string //
	LikeCount   string //
	FollowCount string //
	CreateTime  string //
	UpdateTime  string //
	DeleteTime  string //
}

// userColumns holds the columns for table user.
var userColumns = UserColumns{
	Id:          "id",
	Username:    "username",
	Password:    "password",
	Email:       "email",
	Avatar:      "avatar",
	Sex:         "sex",
	Signature:   "signature",
	Role:        "role",
	Status:      "status",
	FansCount:   "fans_count",
	LikeCount:   "like_count",
	FollowCount: "follow_count",
	CreateTime:  "create_time",
	UpdateTime:  "update_time",
	DeleteTime:  "delete_time",
}

// NewUserDao creates and returns a new DAO object for table data access.
func NewUserDao() *UserDao {
	return &UserDao{
		group:   "default",
		table:   "user",
		columns: userColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *UserDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *UserDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *UserDao) Columns() UserColumns {
	return dao.columns
}

// Group returns the configuration group name of database of current dao.
func (dao *UserDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *UserDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *UserDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
