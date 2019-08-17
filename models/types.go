package models

import "time"

//Task is struct used to identify tasks
type Task struct {
    ID 		int 		`json:"id"`
    Discord int 		`json:"discord"`
    Name 	string 		`json:"name"`
    Status 	int 		`json:"status"`
    Created time.Time 	`json:"created"`
    Updated time.Time 	`json:"updated"`
}