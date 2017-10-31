package dh

import (
	"database/sql"
	"log"
)

type txManager struct {
	txMap map[uint64]*sql.Tx
}

func (tm *txManager) get() (*sql.Tx, *sql.DB, error) {
	if v, ok := tm.txMap[getGID()]; ok {
		return v, nil, nil
	}
	db, err := getDB()
	if err != nil {
		return nil, nil, err
	}
	return nil, db, nil
}

func (tm *txManager) remove() {
	delete(tm.txMap, getGID())
}

func (tm *txManager) set(tx *sql.Tx) {
	tm.txMap[getGID()] = tx
}

var tm *txManager

func transaction(f func() error) error {
	db, err := getDB()
	if err != nil {
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	tm.set(tx)

	if debug {
		log.Print("transaction begin")
	}

	defer func() {

		if err := recover(); err != nil {
			//* 遇到error,回滚事务
			err = tx.Rollback()
		}

		if debug {
			if err != nil {
				log.Print("rollback failure")
			} else {
				log.Print("rollback success")
			}

		}

		tm.remove()
	}()

	//* 执行用户指定的业务
	err = f()
	if err != nil {
		panic(err)
	}

	//* 提交事务,如果有error就抛出
	err = tx.Commit()
	if debug {
		if err != nil {
			log.Print("commit failure")
		} else {
			log.Print("commit success")
		}
	}
	return err
}
