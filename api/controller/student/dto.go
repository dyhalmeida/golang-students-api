package student

type StudentInput struct {
	ID       string `json:"id,omitempty"`
	Fullname string `json:"fullname"`
	Age      int64  `json:"age"`
}
