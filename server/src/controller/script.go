package controller

import (
	"database/sql"
	"log"
)

func rolleback(tx *sql.Tx) {
	if p := recover(); p != nil { // panicに渡された値をrecover()でキャッチ
		tx.Rollback()
		return
	} else { // transactionの実行
		err := tx.Commit()
		if err != nil {
			tx.Rollback()
			log.Print("tx.Commit is failed", err)
		}
		return
	}
}
