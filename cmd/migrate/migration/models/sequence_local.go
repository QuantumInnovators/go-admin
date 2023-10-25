package models

type SequenceLocal struct {
	Id                  int    `json:"id" gorm:"primaryKey;autoIncrement;comment:主键编码"`
	SequenceId          string `json:"sequence_id" gorm:"type:varchar(100);comment:SequenceId"`
	Name                string `json:"name" gorm:"type:varchar(255);comment:Name"`
	NameZh              string `json:"name_cn" gorm:"type:varchar(255);comment:中文名称"`
	SequenceDescription string `json:"sequence_desc" gorm:"type:varchar(255);comment:SequenceDescription"`
	Sequence            string `json:"sequence" gorm:"type:text;comment:Sequence"`
	PrimerName          string `json:"primer_name" gorm:"type:varchar(255);comment:PrimerName"`
	Type                string `json:"type" gorm:"type:varchar(255);comment:Type"`
	ModelTime
	ControlBy
}

func (SequenceLocal) TableName() string {
	return "sequence_local"
}
