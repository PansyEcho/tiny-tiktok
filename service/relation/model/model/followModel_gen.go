// Code generated by goctl. DO NOT EDIT.

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
	followFieldNames          = builder.RawFieldNames(&Follow{})
	followRows                = strings.Join(followFieldNames, ",")
	followRowsExpectAutoSet   = strings.Join(stringx.Remove(followFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	followRowsWithPlaceHolder = strings.Join(stringx.Remove(followFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheFollowIdPrefix = "cache:follow:id:"
)

type (
	followModel interface {
		Insert(ctx context.Context, data *Follow) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Follow, error)
		Update(ctx context.Context, data *Follow) error
		Delete(ctx context.Context, id int64) error
	}

	defaultFollowModel struct {
		sqlc.CachedConn
		table string
	}

	Follow struct {
		Id        int64          `db:"id"` // \'ID\
		CreatedAt sql.NullTime   `db:"created_at"`
		UpdatedAt sql.NullTime   `db:"updated_at"`
		DeletedAt sql.NullTime   `db:"deleted_at"`
		UserId    int64          `db:"user_id"`   // \'关注人ID\
		FollowId  int64          `db:"follow_id"` // \'被关注人ID\
		Cancel    int64          `db:"cancel"`    // \'是否取消关注\
		ExtraA    sql.NullString `db:"extraA"`    // \'额外字段A\
		ExtraB    sql.NullString `db:"extraB"`    // \'额外字段B\
	}
)

func newFollowModel(conn sqlx.SqlConn, c cache.CacheConf) *defaultFollowModel {
	return &defaultFollowModel{
		CachedConn: sqlc.NewConn(conn, c),
		table:      "`follow`",
	}
}

func (m *defaultFollowModel) Delete(ctx context.Context, id int64) error {
	followIdKey := fmt.Sprintf("%s%v", cacheFollowIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, followIdKey)
	return err
}

func (m *defaultFollowModel) FindOne(ctx context.Context, id int64) (*Follow, error) {
	followIdKey := fmt.Sprintf("%s%v", cacheFollowIdPrefix, id)
	var resp Follow
	err := m.QueryRowCtx(ctx, &resp, followIdKey, func(ctx context.Context, conn sqlx.SqlConn, v interface{}) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", followRows, m.table)
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

func (m *defaultFollowModel) Insert(ctx context.Context, data *Follow) (sql.Result, error) {
	followIdKey := fmt.Sprintf("%s%v", cacheFollowIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?)", m.table, followRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.DeletedAt, data.UserId, data.FollowId, data.Cancel, data.ExtraA, data.ExtraB)
	}, followIdKey)
	return ret, err
}

func (m *defaultFollowModel) Update(ctx context.Context, data *Follow) error {
	followIdKey := fmt.Sprintf("%s%v", cacheFollowIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, followRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.DeletedAt, data.UserId, data.FollowId, data.Cancel, data.ExtraA, data.ExtraB, data.Id)
	}, followIdKey)
	return err
}

func (m *defaultFollowModel) formatPrimary(primary interface{}) string {
	return fmt.Sprintf("%s%v", cacheFollowIdPrefix, primary)
}

func (m *defaultFollowModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary interface{}) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", followRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultFollowModel) tableName() string {
	return m.table
}
