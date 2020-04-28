package d

import "log"

func AddTask(id int64, task *Task) (bool, error) {
	stmt, err := db.Prepare("insert into upwork.task (createby, tittle, category, detail, question, tperiod, addition) values (?,?,?,?,?,?,?)")
	if err != nil {
		log.Println("Failed to prepare stmt, err: " + err.Error())
	}
	defer stmt.Close()

	_, err = stmt.Exec(id, task.Tittle, task.Category, task.Detail, task.Question, task.Tperiod, task.Addition)
	if err != nil {
		return false, err
	}
	return true, nil
}
