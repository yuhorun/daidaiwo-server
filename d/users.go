package d

import (
	"fmt"
)

func VerifyPwd(pnumber, upwd string) (uid int64, err error) {
	stmt, err := db.Prepare("select id from upwork.user where pnumber=? and upwd = ? limit 1;")
	if err != nil {
		fmt.Println("Failed to prepare stmt, err: " + err.Error())
		return -1, err
	}
	defer stmt.Close()

	// QueryRow executes a prepared query statement with the given arguments.
	// If an error occurs during the execution of the statement, that error will
	// be returned by a call to Scan on the returned *Row, which is always non-nil.
	// If the query selects no rows, the *Row's Scan will return ErrNoRows.
	// Otherwise, the *Row's Scan scans the first selected row and discards
	// the rest.

	// Scan copies the columns from the matched row into the values
	// pointed at by dest. See the documentation on Rows.Scan for details.
	// If more than one row matches the query,
	// Scan uses the first row and discards the rest. If no row matches
	// the query, Scan returns ErrNoRows.

	err = stmt.QueryRow(pnumber, upwd).Scan(&uid)
	if err != nil {
		return -1, err
	}

	return
}

func AddUser(uname, upwd, pnumber string) (int64, bool, error) {
	stmt, err := db.Prepare("insert into upwork.user (uname, upwd, pnumber) values (?,?,?)")
	if err != nil {
		fmt.Println("Failed to prepare stmt, err: " + err.Error())
		return -1, false, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(uname, upwd, pnumber)
	if err != nil {
		return -1, false, err
	}

	uid, err := res.LastInsertId()
	if err != nil {
		return -1, false, err
	}
	return uid, true, err
}

func SetVerifyCode(key string, value interface{}, time int) error {
	err := redisSet(key, value, time)
	if err != nil {
		return err
	}
	return err
}

func GetUserInfo(uid int64) (user User, err error) {
	stmt, err := db.Prepare("select * from upwork.user where id=? limit 1;")
	if err != nil {
		fmt.Println("Failed to prepare stmt, err: " + err.Error())
		return User{}, err
	}
	defer stmt.Close()

	// QueryRow executes a prepared query statement with the given arguments.
	// If an error occurs during the execution of the statement, that error will
	// be returned by a call to Scan on the returned *Row, which is always non-nil.
	// If the query selects no rows, the *Row's Scan will return ErrNoRows.
	// Otherwise, the *Row's Scan scans the first selected row and discards
	// the rest.

	// Scan copies the columns from the matched row into the values
	// pointed at by dest. See the documentation on Rows.Scan for details.
	// If more than one row matches the query,
	// Scan uses the first row and discards the rest. If no row matches
	// the query, Scan returns ErrNoRows.

	err = stmt.QueryRow(uid).Scan(&user.Id, &user.Uname, &user.Upwd, &user.Pnumber, &user.Ctime, &user.Ltime)
	if err != nil {
		return User{}, err
	}

	return
}
