package repository

import (
	"bioskop/model"
	"database/sql"
)

func GetAllCustomer(db *sql.DB) (result []model.Customer, err error) {
	sql := `SELECT * FROM customer`

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var customer = model.Customer{}

		err = rows.Scan(&customer.CustomerID, &customer.Email, &customer.MobileNumber)
		if err != nil {
			panic(err)
		}

		result = append(result, customer)
	}
	return
}

func UpdateProfile(db *sql.DB, cus model.Customer, CusId int) (err error) {
	sql := `
	UPDATE customer
	SET email=$2, mobilenumber=$3
	WHERE customerid=$1;
	`
	errs := db.QueryRow(sql, CusId, cus.Email, cus.MobileNumber)
	return errs.Err()
}

func SaveProfile(db *sql.DB, cus model.Customer, CusId int) (err error) {
	sql := `
	INSERT INTO customer(
	customerid, email, mobilenumber)
	VALUES ($1, $2, $3);
	`
	errs := db.QueryRow(sql, cus.CustomerID, cus.Email, cus.MobileNumber)
	return errs.Err()
}

func FindCustomerById(db *sql.DB, id int) (customer model.Customer, err error) {
	sql := `SELECT * FROM customer WHERE customerid = $1`
	rows, err := db.Query(sql, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&customer.CustomerID, &customer.Email, &customer.MobileNumber)
		if err != nil {
			panic(err)
		}
	}
	return
}

func MaxcustomerID(db *sql.DB) (max int, err error) {
	sql := "select max(customerid) as max_id from account"
	rows, err := db.Query(sql)
	if err != nil {
		return 0, err
	} else {
		var max_id int
		for rows.Next() {
			rows.Scan(&max_id)
		}
		return max_id, nil
	}
}
