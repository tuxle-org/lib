package entities_test

import (
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/tuxle-org/lib/tuxle/entities"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gotest.tools/assert"
)

func TestEntities(test *testing.T) {
	path := filepath.Join(os.TempDir(), "tuxle-test.db")
	os.Remove(path)
	fmt.Println("- New database file:", path)

	db, err := gorm.Open(sqlite.Open(path), &gorm.Config{})
	assert.NilError(test, err)

	fmt.Println("- Running migrations")
	err = db.AutoMigrate(
		&entities.PermissionMask{},
		&entities.Tag{},
		&entities.Role{},
		&entities.User{},
		&entities.Channel{},
		&entities.Server{},
		&entities.Directory{},
		&entities.MessageVote{},
		&entities.TextMessage{},
	)
	assert.NilError(test, err)

	fmt.Println("- Inserting example values to see if they are valid in SQL")
	assert.NilError(test, db.Create(entities.NewPermissionMask(true)).Error)
	assert.NilError(test, db.Create(entities.NewTag("test", 0, 1)).Error)
	assert.NilError(test, db.Create(entities.NewRole("test", 17591, "/tmp.png", 1)).Error)
	assert.NilError(test, db.Create(entities.NewUser("test", 1)).Error)
	assert.NilError(test, db.Create(entities.NewChannel("test", 0, 1, 1)).Error)
	assert.NilError(
		test,
		db.Create(
			entities.NewServer(
				"test",
				"hello world",
				"1. test",
				"/tmp.png",
				"/banner.png",
				1,
				"en_us",
			),
		).Error,
	)
	assert.NilError(test, db.Create(entities.NewDirectory("test", 0, 1)).Error)
	assert.NilError(test, db.Create(entities.NewMessageVote(true, 1, 1, 1)).Error)
}
