package lib

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"gorm.io/gorm"

	"gorm.io/driver/mysql"
)

// SQLHandler ...
type SQLHandler struct {
	DB  *gorm.DB
	Err error
}

var dbConn *SQLHandler

// DBOpen は DB connectionを張る。
func DBOpen() {
	dbRetryAttempts, err := strconv.ParseUint(os.Getenv("DB_RETRY_ATTEMPTS"), 10, 64)
	if err != nil {
		panic(err)
	}
	dbConn = NewSQLHandler(uint(dbRetryAttempts))
}

// DBClose は DB connectionを張る。
func DBClose() {
	sqlDB, _ := dbConn.DB.DB()
	sqlDB.Close()
}

// NewSQLHandler ...
// count: 接続試行回数
func NewSQLHandler(count uint) *SQLHandler {
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_DATABASE")
	fmt.Println(user, password, host, port)

	var db *gorm.DB
	var err error
	// Todo: USE_HEROKU = 1のときと場合分け
	if os.Getenv("USE_HEROKU") != "1" {
		dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?parseTime=true&loc=Asia%2FTokyo"
		// メモ：docker container名で宛先を指定する場合（どちらでも良い
		// dsn := user + ":" + password + "@tcp(goDockerDB)/" + dbName + "?parseTime=true&loc=Asia%2FTokyo"

		for count >= 1 {
			db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
			if err != nil {
				time.Sleep(time.Second * 2)
				if count == 1 {
					panic(err)
				}
				count--
				fmt.Printf("retry... count:%v\n", count)
				continue
			}
			break
		}

		// db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		// if err != nil {
		// 	panic(err)
		// }

	} /*else {
	    var (
	        instanceConnectionName = os.Getenv("DB_CONNECTION_NAME") // e.g. 'project:region:instance'
	    )
	    dbURI := fmt.Sprintf("%s:%s@unix(/cloudsql/%s)/%s?parseTime=true", user, password, instanceConnectionName, database)
	    // dbPool is the pool of database connections.
	    db, err = gorm.Open(mysql.Open(dbURI), &gorm.Config{})
	    if err != nil {
	        panic(err)
	    }
	}*/

	sqlDB, _ := db.DB()
	//コネクションプールの最大接続数を設定。
	sqlDB.SetMaxIdleConns(100)
	//接続の最大数を設定。 nに0以下の値を設定で、接続数は無制限。
	sqlDB.SetMaxOpenConns(100)
	//接続の再利用が可能な時間を設定。dに0以下の値を設定で、ずっと再利用可能。
	sqlDB.SetConnMaxLifetime(100 * time.Second)

	sqlHandler := new(SQLHandler)
	db.Logger.LogMode(4)
	sqlHandler.DB = db

	return sqlHandler
}

// GetDBConn ...
func GetDBConn() *SQLHandler {
	return dbConn
}

// BeginTransaction ...
func BeginTransaction() *gorm.DB {
	dbConn.DB = dbConn.DB.Begin()
	return dbConn.DB
}

// Rollback ...
func RollBack() {
	dbConn.DB.Rollback()
}
