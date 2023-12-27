package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/LINKHA/automatix/deploy/script/mysql/genModel"

	"github.com/LINKHA/automatix/common/globalkey"

	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	serverFieldNames          = builder.RawFieldNames(&Server{})
	serverRows                = strings.Join(serverFieldNames, ",")
	serverRowsExpectAutoSet   = strings.Join(stringx.Remove(serverFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	serverRowsWithPlaceHolder = strings.Join(stringx.Remove(serverFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheAmxServermanagerServerIdPrefix         = "cache:amxServermanager:server:id:"
	cacheAmxServermanagerServerServerIdPrefix   = "cache:amxServermanager:server:serverId:"
	cacheAmxServermanagerServerTypeSwitchPrefix = "cache:amxServermanager:server:typeswitch:"
)

type (
	serverModel interface {
		Insert(ctx context.Context, data *Server) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Server, error)
		FindOneByServerId(ctx context.Context, serverId string) (*Server, error)
		Update(ctx context.Context, data *Server) error
		UpdateByServerId(ctx context.Context, newData *Server) error
		Delete(ctx context.Context, id int64) error
		SelectBuilder() squirrel.SelectBuilder
		FindAll(ctx context.Context, builder squirrel.SelectBuilder, orderBy string) ([]*Server, error)
	}

	defaultServerModel struct {
		sqlc.CachedConn
		table string
	}

	Server struct {
		Id            int64     `db:"id"`
		CreateTime    time.Time `db:"create_time"`
		UpdateTime    time.Time `db:"update_time"`
		DeleteTime    time.Time `db:"delete_time"`
		DelState      int64     `db:"del_state"`
		ServerId      string    `db:"server_id"`      // 业务Id
		Name          string    `db:"name"`           // 服务器名称
		ServerType    int64     `db:"server_type"`    // 服务类型(0:运营 1:测试)
		Switch        int64     `db:"switch"`         // 服务器是否开启(0:关闭 1:开启)
		StartTime     int64     `db:"start_time"`     // 开服时间
		Area          string    `db:"area"`           // 服务器地区
		Tags          string    `db:"tags"`           // 标签列表(使用逗号分隔)
		MaxOnline     int64     `db:"max_online"`     // 最大在线人数
		MaxQueue      int64     `db:"max_queue"`      // 最大排队人数
		MaxSign       int64     `db:"max_sign"`       // 最大注册人数
		TemplateValue string    `db:"template_value"` // 自定义参数(json格式)
	}
)

func newServerModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultServerModel {
	return &defaultServerModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`server`",
	}
}

func (m *defaultServerModel) Delete(ctx context.Context, id int64) error {
	data, err := m.FindOne(ctx, id)
	if err != nil {
		return err
	}

	amxServermanagerServerIdKey := fmt.Sprintf("%s%v", cacheAmxServermanagerServerIdPrefix, id)
	amxServermanagerServerServerIdKey := fmt.Sprintf("%s%v", cacheAmxServermanagerServerServerIdPrefix, data.ServerId)
	cacheAmxServermanagerServerTypeSwitchPrefix := fmt.Sprintf("%s%v%v", cacheAmxServermanagerServerServerIdPrefix, data.ServerType, data.Switch)

	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, amxServermanagerServerIdKey, amxServermanagerServerServerIdKey, cacheAmxServermanagerServerTypeSwitchPrefix)
	return err
}

func (m *defaultServerModel) FindOne(ctx context.Context, id int64) (*Server, error) {
	amxServermanagerServerIdKey := fmt.Sprintf("%s%v", cacheAmxServermanagerServerIdPrefix, id)
	var resp Server
	err := m.QueryRowCtx(ctx, &resp, amxServermanagerServerIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", serverRows, m.table)
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

func (m *defaultServerModel) FindOneByServerId(ctx context.Context, serverId string) (*Server, error) {
	amxServermanagerServerServerIdKey := fmt.Sprintf("%s%v", cacheAmxServermanagerServerServerIdPrefix, serverId)
	var resp Server
	err := m.QueryRowIndexCtx(ctx, &resp, amxServermanagerServerServerIdKey, m.formatPrimary, func(ctx context.Context, conn sqlx.SqlConn, v any) (i any, e error) {
		query := fmt.Sprintf("select %s from %s where `server_id` = ? limit 1", serverRows, m.table)
		if err := conn.QueryRowCtx(ctx, &resp, query, serverId); err != nil {
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

func (m *defaultServerModel) FindAll(ctx context.Context, builder squirrel.SelectBuilder, orderBy string) ([]*Server, error) {

	builder = builder.Columns(serverRows)

	if orderBy == "" {
		builder = builder.OrderBy("id DESC")
	} else {
		builder = builder.OrderBy(orderBy)
	}

	query, values, err := builder.Where("del_state = ?", globalkey.DelStateNo).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Server
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)

	switch err {
	case nil:
		return resp, nil
	case sqlc.ErrNotFound:
		return nil, genModel.ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultServerModel) Insert(ctx context.Context, data *Server) (sql.Result, error) {
	data.DeleteTime = time.Unix(0, 0)
	data.DelState = globalkey.DelStateNo
	amxServermanagerServerIdKey := fmt.Sprintf("%s%v", cacheAmxServermanagerServerIdPrefix, data.Id)
	amxServermanagerServerServerIdKey := fmt.Sprintf("%s%v", cacheAmxServermanagerServerServerIdPrefix, data.ServerId)
	cacheAmxServermanagerServerTypeSwitchPrefix := fmt.Sprintf("%s%v%v", cacheAmxServermanagerServerServerIdPrefix, data.ServerType, data.Switch)

	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, serverRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.DeleteTime, data.DelState, data.ServerId, data.Name, data.ServerType, data.Switch, data.StartTime, data.Area, data.Tags, data.MaxOnline, data.MaxQueue, data.MaxSign, data.TemplateValue)
	}, amxServermanagerServerIdKey, amxServermanagerServerServerIdKey, cacheAmxServermanagerServerTypeSwitchPrefix)
	return ret, err
}

func (m *defaultServerModel) Update(ctx context.Context, newData *Server) error {
	newData.DeleteTime = time.Unix(0, 0)
	newData.DelState = globalkey.DelStateNo

	data, err := m.FindOne(ctx, newData.Id)
	if err != nil {
		return err
	}

	amxServermanagerServerIdKey := fmt.Sprintf("%s%v", cacheAmxServermanagerServerIdPrefix, data.Id)
	amxServermanagerServerServerIdKey := fmt.Sprintf("%s%v", cacheAmxServermanagerServerServerIdPrefix, data.ServerId)
	cacheAmxServermanagerServerTypeSwitchPrefix := fmt.Sprintf("%s%v%v", cacheAmxServermanagerServerServerIdPrefix, data.ServerType, data.Switch)

	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, serverRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.DeleteTime, newData.DelState, newData.ServerId, newData.Name, newData.ServerType, newData.Switch, newData.StartTime, newData.Area, newData.Tags, newData.MaxOnline, newData.MaxQueue, newData.MaxSign, newData.TemplateValue, newData.Id)
	}, amxServermanagerServerIdKey, amxServermanagerServerServerIdKey, cacheAmxServermanagerServerTypeSwitchPrefix)
	return err
}

func (m *defaultServerModel) UpdateByServerId(ctx context.Context, newData *Server) error {
	newData.DeleteTime = time.Unix(0, 0)
	newData.DelState = globalkey.DelStateNo

	data, err := m.FindOneByServerId(ctx, newData.ServerId)
	if err != nil {
		return err
	}

	amxServermanagerServerIdKey := fmt.Sprintf("%s%v", cacheAmxServermanagerServerIdPrefix, data.Id)
	amxServermanagerServerServerIdKey := fmt.Sprintf("%s%v", cacheAmxServermanagerServerServerIdPrefix, data.ServerId)
	cacheAmxServermanagerServerTypeSwitchPrefix := fmt.Sprintf("%s%v%v", cacheAmxServermanagerServerServerIdPrefix, data.ServerType, data.Switch)

	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `server_id` = ?", m.table, serverRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, newData.DeleteTime, newData.DelState, newData.ServerId, newData.Name, newData.ServerType, newData.Switch, newData.StartTime, newData.Area, newData.Tags, newData.MaxOnline, newData.MaxQueue, newData.MaxSign, newData.TemplateValue, newData.ServerId)
	}, amxServermanagerServerIdKey, amxServermanagerServerServerIdKey, cacheAmxServermanagerServerTypeSwitchPrefix)
	return err
}

func (m *defaultServerModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheAmxServermanagerServerIdPrefix, primary)
}

func (m *defaultServerModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", serverRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultServerModel) tableName() string {
	return m.table
}

func (m *defaultServerModel) SelectBuilder() squirrel.SelectBuilder {
	return squirrel.Select().From(m.table)
}
