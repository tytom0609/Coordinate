package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

// Volunteer Account DataStructure
type Volunteer struct {
	AccountID int    `json:"accountid"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func getVolunteers(db *sql.DB) ([]Volunteer, error) {

	//debugMsg := fmt.Sprintf("Made it here to dlistmodel")
	//fmt.Println(debugMsg)

	Volunteers := []Volunteer{}
	rows, err := db.Query("exec [dbo].[getVolunteers]")
	if err != nil {
		fmt.Println(err)
	} else {
		defer rows.Close()
		for rows.Next() {
			var n Volunteer
			err = rows.Scan(&n.AccountID,
				&n.FirstName,
				&n.LastName)
				Volunteers = append(Volunteers, n)
		}
	}
	return Volunteers, err
}

func (n *Volunteer) updVolunteer(db *sql.DB) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_, err := db.ExecContext(ctx, "[dbo].[updateVolunteer]",
		sql.Named("p_iAccountID", sql.Out{Dest: &n.AccountID}),
		sql.Named("p_sFirstName", n.FirstName),
		sql.Named("p_sLastName", n.LastName))

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (n *Volunteer) createVolunteer(db *sql.DB) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_, err := db.ExecContext(ctx, "[dbo].[updateNonProfit]",
		sql.Named("p_iAccountID", sql.Out{Dest: &n.AccountID}),
		sql.Named("p_sFirstName", n.FirstName),
		sql.Named("p_sLastName", n.LastName))

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (n *Volunteer) delVolunteer(db *sql.DB) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	
	_, err := db.ExecContext(ctx,"[dbo].[deleteVolunteer]", sql.Named("p_iAccountID",&n.AccountID))
	//statement := fmt.Sprintf("Error Message: %s", err)
	//fmt.Println(statement)
	return err
	}
	