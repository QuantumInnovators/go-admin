package models

type Sequence struct {
	Id                  int    `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	SequenceId          string `json:"sequence_id" gorm:"type:varchar(100);comment:SequenceId"`
	SequenceDescription string `json:"sequence_desc" gorm:"type:varchar(255);comment:SequenceDescription"`
	Sequence            string `json:"sequence" gorm:"type:text;comment:Sequence"`
	ModelTime
	ControlBy
}

func (Sequence) TableName() string {
	return "sequence"
}
