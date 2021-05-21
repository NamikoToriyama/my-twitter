package model

import "time"

// Tweet ... Structure of tweet data.
type Tweet struct {
	ID           string    `datastore:"id,omitempty"`
	Username     string    `datastore:"username"`
	Tweet        string    `datastore:"tweet"`
	RegisterDate time.Time `datastore:"registerDate"`
}
