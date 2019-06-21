package data

import "time"

type Thread struct {
	Id       int
	Uuid     string
	Topic    string
	UserId   int
	createAt time.Time
}

//get all threads
func Threads() (threads []Thread, err error) {
	rows, err := Db.Query("SELECT id,uuid,topic,user_id,create_at FROM threads ORDER BY create_at DESC")
	if err != nil {
		return
	}
	// process the query_data
	for rows.Next() {
		th := Thread{}
		if err = rows.Scan(&th.Id, &th.Uuid, &th.Topic, &th.UserId, &th.createAt); err != nil {
			return
		}
		threads = append(threads, th)
	}
	rows.Close()
	return
}
