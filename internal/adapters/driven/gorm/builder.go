package orm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"strconv"
	//"strings"
)

func BuildGormDb(dbDSN, mysqlMaxOpenConns, mysqlMaxIDLEConns string) (*gorm.DB, error) {

	//dbDSN = strings.Trim(dbDSN, "\"")

	db, err := gorm.Open(mysql.Open(dbDSN), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // use singular table name, table for `User` would be `user` with this option enabled
			NoLowerCase:   true, // skip the snake_casing of names
		},
	})
	if err != nil {
		panic(err)
	}

	//TODO: Handle errors.
	maxOpenConns, err := strconv.Atoi(mysqlMaxOpenConns)
	if err != nil {
		panic(err)
	}
	maxIDLEConns, err := strconv.Atoi(mysqlMaxIDLEConns)
	if err != nil {
		panic(err)
	}

	// Let's make shure that we have a pool for concurrent requests (because of the gin)
	mysql_pool_config, err := db.DB()
	mysql_pool_config.SetMaxOpenConns(maxOpenConns)
	mysql_pool_config.SetMaxIdleConns(maxIDLEConns)

	//TODO: Some load balancers or cloud environments (e.g., AWS, GCP)
	// may distribute connections dynamically.
	// Setting a shorter lifetime (e.g., 1 hour) can help your application reconnect periodically
	// to distribute load more evenly across the infrastructure.
	//mysql_pool_config.SetConnMaxLifetime(time.Hour)

	if err != nil {
		return nil, err
	}
	return db, nil
}
