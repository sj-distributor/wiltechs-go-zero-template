// Code generated by goctl. DO NOT EDIT!

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	userFieldNames          = builder.RawFieldNames(&User{})
	userRows                = strings.Join(userFieldNames, ",")
	userRowsExpectAutoSet   = strings.Join(stringx.Remove(userFieldNames, "`create_time`", "`update_time`", "`create_at`", "`update_at`"), ",")
	userRowsWithPlaceHolder = strings.Join(stringx.Remove(userFieldNames, "`id`", "`create_time`", "`update_time`", "`create_at`", "`update_at`"), "=?,") + "=?"

	cacheGoZeroUserIdPrefix = "cache:goZero:user:id:"
)

type (
	userModel interface {
		Insert(ctx context.Context, session sqlx.Session, data *User) (sql.Result, error)
		FindOne(ctx context.Context, id string) (*User, error)
		Update(ctx context.Context, session sqlx.Session, data *User) error
		Delete(ctx context.Context, session sqlx.Session, id string) error
	}

	defaultUserModel struct {
		sqlc.CachedConn
		table string
	}

	User struct {
		Id   string         `db:"id"`
		Name sql.NullString `db:"name"`
	}
)

func newUserModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultUserModel {
	return &defaultUserModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`user`",
	}
}

func (m *defaultUserModel) Delete(ctx context.Context, session sqlx.Session, id string) error {
	goZeroUserIdKey := fmt.Sprintf("%s%v", cacheGoZeroUserIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		if session != nil {
			return session.ExecCtx(ctx, query, id)
		}
		return conn.ExecCtx(ctx, query, id)
	}, goZeroUserIdKey)
	return err
}

func (m *defaultUserModel) FindOne(ctx context.Context, id string) (*User, error) {
	goZeroUserIdKey := fmt.Sprintf("%s%v", cacheGoZeroUserIdPrefix, id)
	var resp User
	err := m.QueryRowCtx(ctx, &resp, goZeroUserIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultUserModel) Insert(ctx context.Context, session sqlx.Session, data *User) (sql.Result, error) {
	goZeroUserIdKey := fmt.Sprintf("%s%v", cacheGoZeroUserIdPrefix, data.Id)
	return m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?)", m.table, userRowsExpectAutoSet)
		if session != nil {
			return session.ExecCtx(ctx, query, data.Id, data.Name)
		}
		return conn.ExecCtx(ctx, query, data.Id, data.Name)
	}, goZeroUserIdKey)
}

func (m *defaultUserModel) Update(ctx context.Context, session sqlx.Session, data *User) error {
	goZeroUserIdKey := fmt.Sprintf("%s%v", cacheGoZeroUserIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, userRowsWithPlaceHolder)
		if session != nil {
			return session.ExecCtx(ctx, query, data.Name, data.Id)
		}
		return conn.ExecCtx(ctx, query, data.Name, data.Id)
	}, goZeroUserIdKey)
	return err
}

func (m *defaultUserModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheGoZeroUserIdPrefix, primary)
}

func (m *defaultUserModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", userRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultUserModel) tableName() string {
	return m.table
}
