package main

import (
	"database/sql"
	f "fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	//connecting ..........................................................

	c := 0
	//var dn string
	db, err := sql.Open("mysql", "root:Thiru#2001@tcp(localhost:3306)/")
	f.Println("Connecting Golang with Mysql !")
	if err != nil {
		f.Println(err)
	}
	err = db.Ping()
	if err != nil {
		f.Println(err)
	}
	f.Println("Succcessful Connected !")
	defer db.Close()

	/*//user wish name ............................

	f.Print("Enter the database name : ")
	f.Scan(&dn)*/

	//create database ....................................................

	data, err := db.Query(`create database emp`)
	if err != nil {
		f.Println(err)
	} else {
		f.Println("Succesfully Created Database !")
	}
	defer data.Close()

	//database change ................................................

	_, err = db.Query("use emp ")
	if err != nil {
		f.Println(err)
	} else {
		f.Println("Datebases Changed To Emp !")
	}

	//create table.........................................................

	create, err := db.Query(`create table emp.record
							 (id int , sno int , name varchar(30),age int,primary key(id));`)
	if err != nil {
		f.Println(err)
	} else {
		f.Println("Successfully Table Was Created !")
	}
	defer create.Close()

	//insert query .......................................................
	insert, err := db.Query(`insert into emp.record values(1,102,'RAM',29),
													 (2,102,'CHERAN',35),
													 (3,103,'ARUN',45),
													 (4,104,'RANU',12),
													 (5,105,'THIRU',21),
													 (6,106,'RAM',23),
													 (7,107,'VICKY',19);`)
	if err != nil {
		f.Println(err)
	} else {
		f.Println("Successfully Inserted The Given Values !")
	}
	defer insert.Close()

	//delete query ...............................................................

	del, err := db.Query(`delete from emp.record where id in(5,6,7);`)
	if err != nil {
		f.Println(err)
	} else {
		f.Println("Successfully Deleted With Given Condition !")
	}
	defer del.Close()

	//alter query .................................................................

	Alter, err := db.Query(`alter table emp.record add Gender varchar(10);`)
	if err != nil {
		f.Println(err)
	} else {
		f.Println("Successfully Alter The Table !")
	}
	defer Alter.Close()

	//update query ...............................................................

	update, err := db.Query(`update emp.record set gender='M';`)
	if err != nil {
		f.Println(err)
	} else {
		f.Println("Successfully Updated !")
	}
	defer update.Close()

	//select query .......................................

	rows, err := db.Query("select *from emp.record")
	checkerr(err)
	if err != nil {
		f.Println(err)
	}
	for rows.Next() {
		var id int
		var sno int
		var name string
		var age int
		var gender string
		err = rows.Scan(&id, &sno, &name, &age, &gender)
		c = c + id
		checkerr(err)
		f.Printf("The result of i%d ,%d ,%s ,%d ,%s", id, sno, name, age, gender)
		f.Print("\n")
	}
	defer rows.Close()

	//rename query...................................................

	rename, err := db.Query(`alter table emp.record rename column gender to sex;`)
	if err != nil {
		f.Println(err)
	} else {
		f.Println("Successfully Renamed Column !")
	}
	defer rename.Close()

	//count function .................................................................
	f.Printf("Sum of the id is %d", c)
	f.Print("\n")
}

func checkerr(err error) {
	if err != nil {
		f.Println(err)
	}
}
