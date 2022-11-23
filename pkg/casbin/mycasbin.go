package mycasbin

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"gorm.io/gorm"
)

// Initialize the model from a string.
var text = `
[request_definition]
r = sub, obj, act

#在sys_casbin_rule使用RBAC 则必须添加
[role_definition]
g = _, _

[policy_definition]
p = sub, obj, act

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = r.sub == p.sub && (keyMatch2(r.obj, p.obj) || keyMatch(r.obj, p.obj)) && (r.act == p.act || p.act == "*")
`

/*
pType有2中p和g
	例如:
	p admin    /user  GET   p是策略 admin是角色 /user是资源路径  GET是请求方式

	g zhangsan admin        g是具体某个用户策略 zhangsan是张三  admin是所属角色

*/

func New(db *gorm.DB, isProd bool) (*casbin.SyncedEnforcer, error) {
	adapt, err := NewAdapterByDBUseTableName(db, "sys", "casbin_rule")
	if err != nil && err.Error() != "invalid DDL" {
		return nil, err
	}
	m, err := model.NewModelFromString(text)
	if err != nil {
		panic(err)
	}
	var enforcer *casbin.SyncedEnforcer
	enforcer, err = casbin.NewSyncedEnforcer(m, adapt)
	if err != nil {
		return nil, err
	}

	err = enforcer.LoadPolicy()
	if err != nil {
		return nil, err
	}
	if !isProd {
		enforcer.EnableLog(true)
	}

	return enforcer, nil
}
