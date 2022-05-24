package event

type Event struct {
	Id      uint   `db:"id,omitempty"`
	Master  string `db:"master"`
	Title   string `db:"title"`
	Descrip string `db:"description"`
	Price   int    `db:"price"`
}
