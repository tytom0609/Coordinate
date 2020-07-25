//dbnd_model.go

package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

// NonProfit Account DataStructure
type NonProfit struct {
	AccountID int    `json:"accountid"`
	Name      string `json:"name"`
}

func getNonProfits(db *sql.DB) ([]NonProfit, error) {
	NonProfits := []NonProfit{}
	rows, err := db.Query("exec [dbo].[getNonProfits]")
	if err != nil {
		fmt.Println(err)
	} else {
		defer rows.Close()
		for rows.Next() {
			var n NonProfit
			err = rows.Scan(&n.AccountID,
				&n.Name)
			NonProfits = append(NonProfits, n)
		}
	}
	return NonProfits, err
}

func (n *NonProfit) updNonProfit(db *sql.DB) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_, err := db.ExecContext(ctx, "[dbo].[updateNonProfit]",
		sql.Named("p_iAccountID", sql.Out{Dest: &n.AccountID}),
		sql.Named("p_sName", n.Name))

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (n *NonProfit) createNonProfit(db *sql.DB) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_, err := db.ExecContext(ctx, "[dbo].[updateNonProfit]",
		sql.Named("p_iAccountID", sql.Out{Dest: &n.AccountID}),
		sql.Named("p_sName", n.Name))

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (n *NonProfit) delNonProfit(db *sql.DB) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	_, err := db.ExecContext(ctx, "[dbo].[deleteNonProfit]", sql.Named("p_iAccountID", &n.AccountID))
	//statement := fmt.Sprintf("Error Message: %s", err)
	//fmt.Println(statement)
	return err
}
