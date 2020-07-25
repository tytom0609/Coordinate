package main

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

// Event Account DataStructure
type Event struct {
	EventID     int    `json:"eventid"`
	Title       string `json:"title"`
	StartDate   string `json:"startdate"`
	EndDate     string `json:"enddate"`
	Description string `json:"description"`
	OrganizerID int    `json:"organizerid"`
}

func getEvents(db *sql.DB) ([]Event, error) {

	//debugMsg := fmt.Sprintf("Made it here to dlistmodel")
	//fmt.Println(debugMsg)

	Events := []Event{}
	rows, err := db.Query("exec [dbo].[getEvents]")
	if err != nil {
		fmt.Println(err)
	} else {
		defer rows.Close()
		for rows.Next() {
			var n Event
			err = rows.Scan(&n.EventID,
				&n.Title,
				&n.StartDate,
				&n.EndDate,
				&n.Description,
				&n.OrganizerID)
			Events = append(Events, n)
			//fmt.Println(err)
		}
	}
	return Events, err
}

func (n *Event) updEvent(db *sql.DB) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_, err := db.ExecContext(ctx, "[dbo].[updateEvent]",
		sql.Named("p_iEventID", sql.Out{Dest: &n.EventID}),
		sql.Named("p_sTitle", n.Title),
		sql.Named("p_sStartDate", n.StartDate),
		sql.Named("p_sEndDate", n.EndDate),
		sql.Named("p_sDescription", n.Description),
		sql.Named("p_iOrganizerID", n.OrganizerID))

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (n *Event) createEvent(db *sql.DB) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_, err := db.ExecContext(ctx, "[dbo].[updateEvent]",
		sql.Named("p_iEventID", sql.Out{Dest: &n.EventID}),
		sql.Named("p_sTitle", n.Title),
		sql.Named("p_sStartDate", n.StartDate),
		sql.Named("p_sEndDate", n.EndDate),
		sql.Named("p_sDescription", n.Description),
		sql.Named("p_iOrganizerID", n.OrganizerID))

	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (n *Event) delEvent(db *sql.DB) error {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	_, err := db.ExecContext(ctx, "[dbo].[deleteEvent]", sql.Named("p_iEventID", &n.EventID))
	//statement := fmt.Sprintf("Error Message: %s", err)
	//fmt.Println(statement)
	return err
}
