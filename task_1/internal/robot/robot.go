package robot

type Robot struct {
	Index, ID int
}

func NewRobot(id int) *Robot {
	return &Robot{ID: id}
}
