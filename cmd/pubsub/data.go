package pubsub

type Data struct {
	Lane   string
	Points int
}

func NewData(Lane string, Points int) *Data {
	return &Data{
		Lane:   Lane,
		Points: Points,
	}
}
func (d *Data) GetLane() string {
	return d.Lane
}
func (d *Data) GetPoints() int {
	return d.Points
}
