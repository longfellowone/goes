package goes

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func MigrateEventsTable() error {
	var err error

	DB.DropTable(&EventDB{})
	if err = DB.AutoMigrate(&EventDB{}).Error; err != nil {
		return err
	}
	return nil
}

// Init initialize the db package
func Init(db *gorm.DB) error {
	DB = db
	return DB.DB().Ping()
}

func InitDB(dbConn string, logMode bool) error {
	var err error

	DB, err = gorm.Open("postgres", dbConn)
	if err != nil {
		return err
	}
	DB.LogMode(logMode)

	return DB.DB().Ping()
}

func IsRecordNotFoundError(err error) bool {
	return gorm.IsRecordNotFoundError(err)
}
