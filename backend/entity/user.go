package entity

import (
	"gorm.io/gorm"
)

type UserType struct {
	gorm.Model

	UserType string

	// 1 User_Type เป็นเจ้าของได้หลาย Users
	Users []User `gorm:"foreignKey:UserType_ID"`
}

type User struct {
	gorm.Model
	User_NAME     string `grom:"uniqueIndex"`
	User_PASSWORD string
	ISNAME        string
	// User_Type_ID  ทำหน้าที่เป็น FK
	UserType_ID *uint
	UserType    UserType
	//1user เป็นเจ้าของได้หลาย Triage
	Triages []Triage `grom:"foreignKey:User_ID"`
}

type Patient struct {
	gorm.Model
	Patient_NAME string
	Gender       string
	Triages      []Triage `grom:"foreignKey:Patient _ID"`
}

type DiseaseType struct {
	gorm.Model
	DiseaseType_NAME string
	Disease          []Disease `grom:"foreignKey:DiseaseType_ID"`
}

type Disease struct {
	gorm.Model
	Disease_NAME   string
	DiseaseType_ID *uint
	DiseaseType    DiseaseType
	Triages        []Triage `grom:"foreignKey:Disease_ID"`
}

type InpantientDepartment struct {
	gorm.Model
	InpantientDepartment_NAME string
	Triages                   []Triage `grom:"foreignKey:IPD_ID"`
}

type Triage struct {
	gorm.Model
	Patient_ID              *uint
	Patient                 Patient
	Disease_ID              *uint
	Disease                 Disease
	InpantientDepartment_ID *uint
	InpantientDepartment    InpantientDepartment
	Triage_COMMENT          string
	User_ID                 *uint
	User                    User
}
