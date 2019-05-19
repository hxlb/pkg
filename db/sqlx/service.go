package sqlx

import (
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/jmoiron/sqlx/reflectx"

	_ "github.com/go-sql-driver/mysql" //mysql
	"github.com/hxlb/pkg/log"
)

// DBService is a database engine object.
type DBService struct {
	Default *sqlx.DB            // the default database engine
	List    map[string]*sqlx.DB // database engine list
}

var dbService = func() (serv *DBService) {
	serv = &DBService{
		List: map[string]*sqlx.DB{},
	}

	var errs []string
	defer func() {
		if len(errs) > 0 {
			panic("[sqlx] " + strings.Join(errs, "\n"))
		}
		if serv.Default == nil {
			log.Logger("sqlx").Panic("[sqlx] the `default` database engine must be configured and enabled")
		}
	}()

	err := loadDBConfig()
	if err != nil {
		log.Logger("sqlx").Panic(err.Error())
		return
	}

	for _, conf := range dbConfigs {
		if !conf.Enable {
			continue
		}
		db, err := sqlx.Connect(conf.Driver, conf.Connstring)
		if err != nil {
			log.Logger("sqlx").Panic(err.Error())
			errs = append(errs, err.Error())
			continue
		}

		db.SetMaxOpenConns(conf.MaxOpenConns)
		db.SetMaxIdleConns(conf.MaxIdleConns)

		var strFunc = strings.ToLower
		if conf.ColumnSnake {
			//strFunc = faygo.SnakeString
		}

		// Create a new mapper which will use the struct field tag "json" instead of "db"
		db.Mapper = reflectx.NewMapperFunc(conf.StructTag, strFunc)

		serv.List[conf.Name] = db
		if DEFAULTDB_NAME == conf.Name {
			serv.Default = db
		}
	}
	return
}()
