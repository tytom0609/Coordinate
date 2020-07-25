package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

// Qualification Account DataStructure
type Qualification struct {
	QualificationID     	 int    `json:"requestid"`
	Title      		 string `json:"title"`
	Description  	 string `json:"description"`
}

func getQualifications(db *sql.DB) ([]Qualification, error) {

	//debugMsg := fmt.Sprintf("Made it here to dlistmodel")
	//fmt.Println(debugMsg)

	Qualifications := []Qualification{}
	rows, err := db.Query("exec [dbo].[getQualifications]")
	if err != nil {
		fmt.Println(err)
	} else {
		defer rows.Close()
		for rows.Next() {
			var n Qualification
			err = rows.Scan(&n.QualificationID,
				&n.Title,
				&n.Description)
				Qualifications = append(Qualifications, n)
		}
	}
	return Qualifications, err
}

func (n *Qualification) updQualification(db *sql.DB) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_, err := db.ExecContext(ctx, "[dbo].[updateQualification]",
		sql.Named("p_iQualificationID", sql.Out{Dest: &n.QualificationID}),
		sql.Named("p_sTitle", n.Title),
		sql.Named("p_sDescription", n.Description))

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (n *Qualification) createQualification(db *sql.DB) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_, err := db.ExecContext(ctx, "[dbo].[updateQualification]",
		sql.Named("p_iQualificationID", sql.Out{Dest: &n.QualificationID}),
		sql.Named("p_sTitle", n.Title),
		sql.Named("p_sDescription", n.Description))

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (n *Qualification) delQualification(db *sql.DB) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	_, err := db.ExecContext(ctx, "[dbo].[deleteQualification]", sql.Named("p_iQualificationID", &n.QualificationID))
	//statement := fmt.Sprintf("Error Message: %s", err)
	//fmt.Println(statement)
	return err
}
