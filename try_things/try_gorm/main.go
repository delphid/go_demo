package main

import (
	"fmt"
	"os"
	"time"

	// "go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	// "moul.io/zapgorm2"
)

type Audit struct {
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Name      string    `json:"name"`
	Date      time.Time `json:"date"`
	Value     int       `json:"value"`
}

type User struct {
	ID    string   `json:"id"`
	Name  string   `json:"name"`
	Tags  []string `json:"tags"`
	Teams []Team   `json:"teams" gorm:"many2many:user_m2m_team;"`
}

type Team struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Users []User `json:"users" gorm:"many2many:user_m2m_team;"`
}

func NewDB() *gorm.DB {
	const ConnParams = "?charset=utf8mb4&parseTime=true&loc=Local&multiStatements=true"
	// GO_DEMO_DB = 'user:pass@tcp(host:port)/db'
	dsn := os.Getenv("GO_DEMO_DB") + ConnParams
	mysqlConfig := mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}

	db, _ := gorm.Open(mysql.New(mysqlConfig), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})

	return db
}

func main() {
	db := NewDB()
	fmt.Println(db)
	rawDate := "2022-07-20"
	loc := time.FixedZone("UTC+8", +8*60*60)
	date, err := time.ParseInLocation("2006-01-02", rawDate, loc)
	if err != nil {
		fmt.Println(err)
	}
	valueA := 3
	valueB := 4
	audits := []Audit{
		{Name: "aaa", Date: date, Value: valueA},
		{Name: "aaa", Date: date, Value: valueA},
		{Name: "bbb", Date: date, Value: valueB},
	}
	if err = db.Save(&audits).Error; err != nil {
		fmt.Println(err)
	}

	var names []string
	db.Model(&Audit{}).Select("name").Find(&names)
	fmt.Println(names)
	var sum int
	db.Model(&Audit{}).Where("name = ?", "aaa").Select("case when sum(value) is null then 0 else sum(value) end").Find(&sum)
	fmt.Println(sum)
	var m []Audit
	db.Find(&m)
	fmt.Println(m)
	db.Model(&Audit{}).Where("name != ?", "''").Find(&m)
	fmt.Println(m)
	var nv []Audit
	db.Model(&Audit{}).Select("name, sum(value) value").Group("name").Find(&nv)
	fmt.Println(nv)
	//if err = db.Where("1=1").Delete(&User{}).Error; err != nil {
	//	fmt.Println(err)
	//}
	//if err = db.Where("1=1").Delete(&User{}).Error; err != nil {
	//	fmt.Println(err)
	//}
	users := []User{
		{ID: "1", Name: "ua"},
		{ID: "2", Name: "ub"},
	}
	teams := []Team{
		{ID: "1", Name: "ta"},
		{ID: "2", Name: "tb"},
	}
	if err = db.Save(&users).Error; err != nil {
		fmt.Println(err)
	}
	if err = db.Save(&teams).Error; err != nil {
		fmt.Println(err)
	}
	//db.Model(&User{ID: "1"}).Association("Teams").Clear()
	db.Model(&User{ID: "1"}).Association("Teams").Append(&Team{ID: "1"})
	db.Model(&User{ID: "1"}).Association("Teams").Append(&Team{ID: "2"})
	u := User{Name: "ua"}
	db.FirstOrInit(&u)
	fmt.Println(u)
}
