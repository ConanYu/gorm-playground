package main

import (
	"fmt"
	"path"
	"path/filepath"
	"runtime"
	"testing"

	"gorm.io/gen"
)

// GORM_REPO: https://github.com/go-gorm/gorm.git
// GORM_BRANCH: master
// TEST_DRIVERS: sqlite, mysql, postgres, sqlserver

func TestGORM(t *testing.T) {
	user := User{Name: "jinzhu"}

	DB.Create(&user)

	var result User
	if err := DB.First(&result, user.ID).Error; err != nil {
		t.Errorf("Failed, got error: %v", err)
	}
}

func TestGormGen(t *testing.T) {
	_, filename, _, _ := runtime.Caller(0)
	root := path.Dir(filename)
	//modelPkgPath := path.Join(root, "model") // gorm.io/gen v0.3.19 Windows下会报错
	modelPkgPath := filepath.ToSlash(path.Join(root, "model")) // gorm.io/gen v0.3.19 Windows下也会报错
	fmt.Println(modelPkgPath)
	g := gen.NewGenerator(gen.Config{
		ModelPkgPath: modelPkgPath,
	})
	g.UseDB(DB)
	g.GenerateModel("users")
	g.Execute()
}
