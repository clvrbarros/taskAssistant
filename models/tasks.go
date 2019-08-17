package models

func TasksById (idDiscord string) ([]*Task, error) {
	rows, err := db.Query("SELECT id, discord, name, status, data_create FROM TASKS WHERE discord = $1",idDiscord)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	tasks := make([]*Task,0)
	for rows.Next() {
		task := new(Task)
		err := rows.Scan(&task.ID, &task.Discord, &task.Name, &task.Status, &task.Created)
		if err != nil {
			return nil, err
		}
		tasks = append(tasks, task)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}

func AddTask(title string, idDiscord string) error {
	var id int
	row := db.QueryRow("select count(id) from tasks where discord = $1", idDiscord)
	err := row.Scan(&id)
	if err != nil {
		return err
	}
	sqlStmt := "INSERT INTO tasks(discord, id, name, status) values($1,$2,$3,1)"
	_, err = db.Exec(sqlStmt, idDiscord, id + 1, title)
	if err != nil {
		return err
	}
	return nil
}