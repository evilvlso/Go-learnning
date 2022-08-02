package serialize

import "example.com/to_list/model"

type Response struct {
	Status int         `json:"status"`
	Data   interface{} `json:"data,omitempty"`
	Msg    string      `json:"msg,omitempty"`
	Error  string      `json:"error,omitempty"`
}

type Token struct {
	Data string
	err  error
}

func (t *Token) SerializeToken() map[string]string {
	return map[string]string{"token": t.Data}
}

type DataList struct {
	Count int           `json:"count"`
	Data  []interface{} `json:"data"`
}

type Task struct {
	Tid     uint   `json:"tid"`
	Title   string `json:"title"`
	Comment string `json:"comment"`
	Status  uint8  `json:"status"`
}

func SerializeTask(tasks []model.Task) DataList {
	var dataList DataList
	dataList.Count = len(tasks)
	for _, task := range tasks {
		t := &Task{
			Tid:     task.ID,
			Title:   task.Title,
			Comment: task.Comment,
			Status:  task.Status,
		}
		dataList.Data = append(dataList.Data, t)
	}
	return dataList
}
