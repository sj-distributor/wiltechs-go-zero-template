package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	"github.com/golang-migrate/migrate/v4/source/file"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
)

var configFile = flag.String("f", "../rpc/etc/app.yaml", "the config file")
var sourceFile = flag.String("sf", "./sql", "the sql source file")

type Config struct {
	Mysql struct {
		DataSource string
	}
}

func main() {
	var c Config
	var cmd string
	var version int
	flag.StringVar(&cmd, "c", "up", "migrate command")
	flag.IntVar(&version, "v", -1, "migrate to the version")

	flag.Parse()
	conf.MustLoad(*configFile, &c)

	db, err := sql.Open("mysql", c.Mysql.DataSource)
	if err != nil {
		logx.Errorf("migrate error:", err)
		return
	}

	driver, err := mysql.WithInstance(db, &mysql.Config{})
	if err != nil {
		logx.Errorf("migrate error:", err)
		return
	}

	// 获取sql文件
	fs, err := (&file.File{}).Open("file://" + *sourceFile)
	if err != nil {
		logx.Errorf("migrate error:", err)
		return
	}

	m, err := migrate.NewWithInstance(
		"file", fs,
		"mysql", driver)
	if err != nil {
		logx.Errorf("migrate error:", err)
		return
	}

	// 执行迁移
	if err = migrateLogic(m, cmd, version); err != nil {
		logx.Errorf("migrate error:", err)
		return
	}

	logx.Info("database migrate success.")

	defer driver.Close()
	defer db.Close()
	defer fs.Close()
}

// 迁移执行逻辑
func migrateLogic(m *migrate.Migrate, cmd string, version int) error {
	if version == -1 && cmd == "up" {
		if err := m.Up(); err != nil {
			return err
		}
	} else if version > 0 && cmd == "goto" {
		var isContinue string

		fmt.Println("危险操作！！！回滚数据库可能导致数据丢失！！！是否操作[N/y]")

		if _, err := fmt.Scanln(&isContinue); err != nil {
			return err
		}

		if isContinue == "y" {
			if err := m.Migrate(uint(version)); err != nil {
				return err
			}
		} else {
			return errors.New("migrate error: nothing migrate")
		}
	} else {
		return errors.New("migrate error: parameter error")
	}
	return nil
}
