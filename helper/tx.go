package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {

	// catch error
	err := recover()
	if err != nil {
		// error rollback
		errorRollback := tx.Rollback()
		// send to function PanicIfError, if errorRollback is not nil
		PanicIfError(errorRollback)

		// panic error if errorRollback is nil, it's mean error but not from rollback
		panic(err)
	} else {
		// commit if no error
		errorCommit := tx.Commit()

		// send to function PanicIfError, if errorCommit is not nil
		PanicIfError(errorCommit)
	}
}
