package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"sync"
)

var (
	// CommonDB session, version, log db handler
	db *gorm.DB

	cveDBSet     = make(map[string]*gorm.DB)
	cveSetLocker sync.Mutex
)

// Init init db
func Init(host  string,pwd_sql string) {
	var err error
	db, err = gorm.Open("mysql","root:"+pwd_sql+"@tcp("+host+":32680)/deepin_cve?parseTime=true")
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&Session{})
	db.AutoMigrate(&Version{})
	db.AutoMigrate(&Log{})
	// TODO(jouyouyun): add to configuration
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	var verList VersionList
	err = db.Find(&verList).Error
	if err != nil {
		panic(err)
	}

	for _, ver := range verList {
		err = doSetDBHandler(ver.Version)
		if err != nil {
			panic(err)
		}
	}
}

// GetDBHandler return db handler by version
func GetDBHandler(version string) (*gorm.DB) {
	handler, ok := cveDBSet[version]
	if !ok {
		return nil
	}
	return handler
}

// SetDBHandler init db handler
func SetDBHander(version string) error {
	cveSetLocker.Lock()
	defer cveSetLocker.Unlock()

	if handler, ok := cveDBSet[version]; ok && handler != nil {
		return nil
	}
	return doSetDBHandler(version)
}

// DeleteDBHandler delete db handler
func DeleteDBHandler(version string) error {
	cveSetLocker.Lock()
	defer cveSetLocker.Unlock()

	handler, ok := cveDBSet[version]
	delete(cveDBSet, version)
	if !ok || handler == nil {
		return nil
	}
	return handler.Close()
}

func doSetDBHandler(version string) error {

	db.AutoMigrate(&CVE{VersionId: version})
	db.AutoMigrate(&Package{VersionId: version})
	db.AutoMigrate(&CVEScore{VersionId: version})
	// TODO(jouyouyun): add to configuration
	db.DB().SetMaxIdleConns(0)
	db.DB().SetMaxOpenConns(100)
	cveDBSet[version] = db
	return nil
}

