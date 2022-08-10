package moduledept

type Dept struct {
	DeptId      string `json:"dep_id" gorm:"column:dep_id;primaryKey"`
	NameDept    string `json:"name" gorm:"name_dept"`
	Addr        string `json:"address" gorm:"address"`
	Description string `json:"description" gorm:"description"`
}

func (rp Dept) TableName() string {
	return "depts"
}
