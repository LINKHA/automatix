// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/LINKHA/automatix/deploy/script/mysql/genModel"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	roleFieldNames          = builder.RawFieldNames(&Role{})
	roleRows                = strings.Join(roleFieldNames, ",")
	roleRowsExpectAutoSet   = strings.Join(stringx.Remove(roleFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	roleRowsWithPlaceHolder = strings.Join(stringx.Remove(roleFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheAmxRolemanagerRoleIdPrefix     = "cache:amxRolemanager:role:id:"
	cacheAmxRolemanagerRoleRoleIdPrefix = "cache:amxRolemanager:role:roleId:"
)

type (
	roleModel interface {
		Insert(ctx context.Context, data *Role) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Role, error)
		FindOneByRoleId(ctx context.Context, roleId string) (*Role, error)
		Update(ctx context.Context, data *Role) error
		UpdateByRoleId(ctx context.Context, data *Role) error
		Delete(ctx context.Context, id int64) error
		DeleteByRoleId(ctx context.Context, roleId string) error
	}

	defaultRoleModel struct {
		sqlc.CachedConn
		table string
	}

	Role struct {
		Id               int64     `db:"id"`
		CreateTime       time.Time `db:"create_time"`
		UpdateTime       time.Time `db:"update_time"`
		RoleId           string    `db:"role_id"`            // 业务Id
		BornServerId     string    `db:"born_server_id"`     // 出生服id
		CurServerId      string    `db:"cur_server_id"`      // 当前服id
		HistoryServerIds string    `db:"history_server_ids"` // 历史区服
		Tags             string    `db:"tags"`               // 标签组
		TemplateValue    string    `db:"template_value"`     // 自定义参数
	}
)

func newRoleModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultRoleModel {
	return &defaultRoleModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`role`",
	}
}

func (m *defaultRoleModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	amxRolemanagerRoleIdKey := fmt.Sprintf("%s%v", cacheAmxRolemanagerRoleIdPrefix, id)
	amxRolemanagerRoleRoleIdKey := fmt.Sprintf("%s%v", cacheAmxRolemanagerRoleRoleIdPrefix, data.RoleId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, amxRolemanagerRoleIdKey, amxRolemanagerRoleRoleIdKey)
	return err
}

func (m *defaultRoleModel) DeleteByRoleId(ctx context.Context, roleId string) error {
	data, err := m.FindOneByRoleId(ctx, roleId)
	if err != nil {
		return err
	}

	amxRolemanagerRoleRoleIdKey := fmt.Sprintf("%s%v", cacheAmxRolemanagerRoleRoleIdPrefix, data.RoleId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `role_id` = ?", m.table)
		return conn.ExecCtx(ctx, query, roleId)
	}, amxRolemanagerRoleRoleIdKey)
	return err
}

func (m *defaultRoleModel) FindOne(ctx context.Context, id int64) (*Role, error) {
	amxRolemanagerRoleIdKey := fmt.Sprintf("%s%v", cacheAmxRolemanagerRoleIdPrefix, id)
	var resp Role
	err := m.QueryRowCtx(ctx, &resp, amxRolemanagerRoleIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", roleRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, genModel.ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultRoleModel) FindOneByRoleId(ctx context.Context, roleId string) (*Role, error) {
	amxRolemanagerRoleRoleIdKey := fmt.Sprintf("%s%v", cacheAmxRolemanagerRoleRoleIdPrefix, roleId)
	var resp Role
	err := m.QueryRowIndexCtx(ctx, &resp, amxRolemanagerRoleRoleIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `role_id` = ? limit 1", roleRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, roleId); err != nil {
			return nil, err
		}
		return resp.Id, nil
	}, m.queryPrimary)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, genModel.ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultRoleModel) Insert(ctx context.Context, data *Role) (sql.Result, error) {
	amxRolemanagerRoleIdKey := fmt.Sprintf("%s%v", cacheAmxRolemanagerRoleIdPrefix, data.Id)
	amxRolemanagerRoleRoleIdKey := fmt.Sprintf("%s%v", cacheAmxRolemanagerRoleRoleIdPrefix, data.RoleId)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?)", m.table, roleRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Id, data.RoleId, data.BornServerId, data.CurServerId, data.HistoryServerIds, data.Tags, data.TemplateValue)
	}, amxRolemanagerRoleIdKey, amxRolemanagerRoleRoleIdKey)
	return ret, err
}

func (m *defaultRoleModel) Update(ctx context.Context, newData *Role) error {
	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	amxRolemanagerRoleIdKey := fmt.Sprintf("%s%v", cacheAmxRolemanagerRoleIdPrefix, data.Id)
	amxRolemanagerRoleRoleIdKey := fmt.Sprintf("%s%v", cacheAmxRolemanagerRoleRoleIdPrefix, data.RoleId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, roleRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.RoleId, newData.BornServerId, newData.CurServerId, newData.HistoryServerIds, newData.Tags, newData.TemplateValue, newData.Id)
	}, amxRolemanagerRoleIdKey, amxRolemanagerRoleRoleIdKey)
	return err
}

func (m *defaultRoleModel) UpdateByRoleId(ctx context.Context, newData *Role) error {
	data, err := m.FindOneByRoleId(ctx, newData.RoleId)
	if err != nil {
		return err
	}

	amxRolemanagerRoleIdKey := fmt.Sprintf("%s%v", cacheAmxRolemanagerRoleIdPrefix, data.Id)
	amxRolemanagerRoleRoleIdKey := fmt.Sprintf("%s%v", cacheAmxRolemanagerRoleRoleIdPrefix, data.RoleId)
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, roleRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.RoleId, newData.BornServerId, newData.CurServerId, newData.HistoryServerIds, newData.Tags, newData.TemplateValue, newData.Id)
	}, amxRolemanagerRoleIdKey, amxRolemanagerRoleRoleIdKey)
	return err
}

func (m *defaultRoleModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheAmxRolemanagerRoleIdPrefix, primary)
}

func (m *defaultRoleModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", roleRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultRoleModel) tableName() string {
	return m.table
}
