package helper

import "gorm.io/gorm"

func CommitOrRollback(tx *gorm.DB) {
	err := recover()
	if err != nil {
		rollback := tx.Rollback()
		PanicIfError(rollback.Error)
		panic(err)
	} else {
		commit := tx.Commit()
		PanicIfError(commit.Error)
	}
}
