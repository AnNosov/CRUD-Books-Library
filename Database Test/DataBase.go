package main

import (
	"database/sql"
	"fmt"
	"github.com/mattn/go-oci8"
	"log"
)

func main() {

	fmt.Println(sql.Drivers())

	db, err := sql.Open("Oracle", "Oracle://cms:cms@tcp(nba48:7621)")

	if err != nil {
		panic(err)
	}

	value, err := db.Query("select crm_camp_id from v_crm_informing where camp_id in (454601);", 1)

	var crm_camp_id string

	if err != nil {
		log.Fatal(err)
	}
	defer value.Close()
	for value.Next() {
		err := value.Scan(&crm_camp_id)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(crm_camp_id)
	}
	err = value.Err()
	if err != nil {
		log.Fatal(err)
	}

}
