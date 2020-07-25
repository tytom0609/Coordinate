package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

// Request Account DataStructure
type Request struct {
	RequestID        int    `json:"requestid"`
	Title            string `json:"title"`
	Description      string `json:"description"`
	QuantityNeeded   int    `json:"quantityneeded"`
	QuantityObtained int    `json:"quantityobtained"`
	Closed           int    `json:"closed"`
	RequestType      int    `json:"requesttype"`
}

func getRequests(db *sql.DB) ([]Request, error) {

	//debugMsg := fmt.Sprintf("Made it here to dlistmodel")
	//fmt.Println(debugMsg)

	Requests := []Request{}
	rows, err := db.Query("exec [dbo].[getRequests]")
	if err != nil {
		fmt.Println(err)
	} else {
		defer rows.Close()
		for rows.Next() {
			var n Request
			err = rows.Scan(&n.RequestID,
				&n.Title,
				&n.Description,
				&n.QuantityNeeded,
				&n.QuantityObtained,
				&n.Closed,
				&n.RequestType)
			Requests = append(Requests, n)
		}
	}
	return Requests, err
}

func (n *Request) updRequest(db *sql.DB) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_, err := db.ExecContext(ctx, "[dbo].[updateRequest]",
		sql.Named("p_iRequestID", sql.Out{Dest: &n.RequestID}),
		sql.Named("p_sTitle", n.Title),
		sql.Named("p_sDescription", n.Description),
		sql.Named("p_iQuantityNeeded", n.QuantityNeeded),
		sql.Named("p_iQuantityObtained", n.QuantityObtained),
		sql.Named("p_iClosed", n.Closed),
		sql.Named("p_iRequestType", n.RequestType))

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (n *Request) createRequest(db *sql.DB) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_, err := db.ExecContext(ctx, "[dbo].[updateRequest]",
		sql.Named("p_iRequestID", sql.Out{Dest: &n.RequestID}),
		sql.Named("p_sTitle", n.Title),
		sql.Named("p_sDescription", n.Description),
		sql.Named("p_iQuantityNeeded", n.QuantityNeeded),
		sql.Named("p_iQuantityObtained", n.QuantityObtained),
		sql.Named("p_iClosed", n.Closed),
		sql.Named("p_iRequestType", n.RequestType))

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (n *Request) delRequest(db *sql.DB) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	_, err := db.ExecContext(ctx, "[dbo].[deleteRequest]", sql.Named("p_iRequestID", &n.RequestID))
	//statement := fmt.Sprintf("Error Message: %s", err)
	//fmt.Println(statement)
	return err
}
