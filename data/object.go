package data

type Object struct {
	Handle  string `json:"handle"`
	Content string `json:"content"`
	StartAt string `json:"startAt"`
	EndAt   string `json:"endAt"`
}

func CreateNewObject(handle, content string, startTime, endTime string) (Object, error) {
	newObj := Object{
		Handle:  handle,
		Content: content,
		StartAt: startTime,
		EndAt:   endTime,
	}

	query := `INSERT INTO objects (handle, content, start_at, end_at) VALUES (?, ?, ?, ?)`
	_, err := DB.Exec(query, newObj.Handle, newObj.Content, newObj.StartAt, newObj.EndAt)
	return newObj, err
}

func GetObjectByHandle(handle string) ([]Object, error) {
	var obj []Object
	query := `SELECT handle, content, start_at, end_at FROM objects WHERE handle=?`
	rows, err := DB.Query(query, handle)

	for rows.Next() {
		var o Object
		err = rows.Scan(&o.Handle, &o.Content, &o.StartAt, &o.EndAt)

		obj = append(obj, o)
	}

	return obj, err
}
