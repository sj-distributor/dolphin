package templates

var Database = `package gen

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/iancoleman/strcase"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
)

type DB struct {
	db *gorm.DB
}

type SyncShardingFunc func(db *DB) error

func NewDBFromEnvVars(name string) *DB {
	urlString := name

	if urlString == "" {
		urlString = os.Getenv("DATABASE_URL")
	}

	if urlString == "" {
		panic(fmt.Errorf("missing DATABASE_URL environment variable"))
	}
	return NewDBWithString(urlString)
}

func NewDBWithString(urlString string) *DB {
	u, err := url.Parse(urlString)
	if err != nil {
		panic(err)
	}

	// urlString = GetConnectionString(u)

	gormConfig := &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix: os.Getenv("TABLE_NAME_PREFIX"),
		},
	}

	if os.Getenv("DEBUG") == "true" {
		gormConfig.Logger = logger.Default.LogMode(logger.Info)
	}

	dsn, err := GetConnectionString(u)
	if err != nil {
		panic(err)
	}

	db, err := gorm.Open(dsn, gormConfig)
	if err != nil {
		panic(err)
	}

	return NewDB(db)
}

func GetConnectionString(u *url.URL) (gorm.Dialector, error) {
	if u.Scheme == "postgres" {
		password, _ := u.User.Password()
		params := u.Query()
		params.Set("host", strings.Split(u.Host, ":")[0])
		params.Set("port", u.Port())
		params.Set("user", u.User.Username())
		params.Set("password", password)
		params.Set("dbname", strings.TrimPrefix(u.Path, "/"))
		return postgres.Open(strings.Replace(params.Encode(), "&", " ", -1)), nil
	}

	if u.Scheme == "sqlite3" {
		return sqlite.Open(strings.Replace(u.String(), u.Scheme+"://", "", 1)), nil
	}

	if u.Scheme == "mysql" {
		u.Host = "tcp(" + u.Host + ")"
		q := u.Query()
		q.Set("parseTime", "true")
		u.RawQuery = q.Encode()
		return mysql.Open(strings.Replace(u.String(), u.Scheme+"://", "", 1)), nil
	}

	return nil, fmt.Errorf("db not support")
}

func NewDB(db *gorm.DB) *DB {
	v := DB{db}
	// InitGorm(db)
	return &v
}

var ShardingArray = []string{ {{range $obj := .Model.ObjectEntities}}{{if $obj.IsSharding}}
	"{{.TableName}}",{{end}}{{end}}
}

var ShardingStruct = []any{ {{range $obj := .Model.ObjectEntities}}{{if $obj.IsSharding}}
	{{.Name}}{},{{end}}{{end}}
}

var TableMap = map[string]interface{}{ {{range $obj := .Model.ObjectEntities}}
	"{{.TableName}}": {{.Name}}{},{{end}}
}

// 获取表名
func GetShardingTableName(name string, shardingId string) string {
	// secretKey := ctx.Value(config.KeyAppSecret)
	if shardingId != "" && utils.StrIndexOf(ShardingArray, name) != -1 {
		return name + "_" + shardingId
	}

	return name
}

func TableName(name string, ctx context.Context) string {
	secretKey := ctx.Value(config.KeyAppSecret)
	shardingId := utils.ExtractShardingTableName(secretKey)

	prefix := os.Getenv("TABLE_NAME_PREFIX")
	tableName := strcase.ToSnake(strcase.ToLowerCamel(name))
	if prefix != "" {
		tableName = prefix + "_" + tableName
	}

	// return strcase.ToSnake(strcase.ToLowerCamel(name))
	return GetShardingTableName(tableName, shardingId)
}

// Close ...
func (db *DB) Close() error {
	return db.Close()
}

// Query ...
func (db *DB) Query() *gorm.DB {
	return db.db
}

// AutoMigrate ...
func (db *DB) AutoMigrate() error {
	return db.db.AutoMigrate({{range $obj := .Model.ObjectEntities}}
		{{.Name}}{},{{end}}
	)
}

// 同步分表字段结构
func (db *DB) SyncShardingTable(callback SyncShardingFunc) error {
	return callback(db)
}

func (db *DB) Ping() error {
	return db.Ping()
}`
