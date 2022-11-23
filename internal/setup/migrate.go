package setup

import (
	"back-admin/store"
	"github.com/pkg/errors"
	"io/ioutil"
	"strings"
)

type Service struct {
	Db store.Db
}

func NewSrv(db store.Db) *Service {
	return &Service{
		db,
	}
}

func (s *Service) Setup(tables []interface{}, filePath string) error {
	var err error
	//if err = s.Db.CreateSchema(); err != nil {
	//	return err
	//}
	//判断表是否存在
	//for _, t := range tables {
	//	s.Db.HasTable(t)
	//}
	if err = s.Db.Schema().Migrate(tables); err != nil {
		return err
	}

	//初始化表数据
	if err = s.execSql(filePath); err != nil {
		return err
	}
	return s.Db.Schema().Close()
}

func (s *Service) execSql(filePath string) error {
	sql, err := s.ioUtil(filePath)
	if err != nil {
		return errors.WithStack(err)
	}
	sqlList := strings.Split(sql, ";")
	for i := 0; i < len(sqlList)-1; i++ {
		if strings.Contains(sqlList[i], "--") {
			continue
		}
		sql := strings.Replace(sqlList[i]+";", "\n", "", -1)
		sql = strings.TrimSpace(sql)
		if err = s.Db.Schema().Exec(sql); err != nil {
			if !strings.Contains(err.Error(), "Query was empty") {
				return errors.WithStack(err)
			}
		}
	}
	return nil
}

func (s *Service) ioUtil(filePath string) (string, error) {
	if contents, err := ioutil.ReadFile(filePath); err == nil {
		//因为contents是[]byte类型，直接转换成string类型后会多一行空格,需要使用strings.Replace替换换行符
		result := strings.Replace(string(contents), "\n", "", 1)
		return result, nil
	} else {
		return "", errors.WithStack(err)
	}
}
