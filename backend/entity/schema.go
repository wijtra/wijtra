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
	UserTypeID *uint
	UserType   UserType
	//1user เป็นเจ้าของได้หลาย Triage
	Triages []Triage `grom:"foreignKey:User_ID"`
}

type Gender struct {
	gorm.Model
	Gender_NAME string
	Patient     []Patient `grom:"foreignKey:Gender_ID"`
}

type Blood_type struct {
	gorm.Model
	Blood_Name string
	// หมู่เลือด 1 ประเภท เป็นของผู้ป่วยในได้ หลายคน
	Patient []Patient `gorm:"foreignKey:Blood_type_ID"`
}

type Patient struct {
	gorm.Model
	Patient_NAME  string
	Gender_ID     *uint
	Gender        string
	Blood_type_ID *uint
	Blood_type    string
	Triages       []Triage `grom:"foreignKey:Patient_ID"`
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
	Triages                   []Triage `grom:"foreignKey: InpantientDepartment_ID"`
}

type Triage_State struct {
	gorm.Model
	Triage_State string
	Triages      []Triage `grom:"foreignKey: Triage_State_ID"`
}

type Triage struct {
	gorm.Model
	Patient_ID              *uint
	Patient                 Patient
	Disease_ID              *uint
	Disease                 Disease
	InpantientDepartment_ID *uint
	InpantientDepartment    InpantientDepartment
	Triage_State_ID         *uint
	Triage_State            Triage_State
	Triage_COMMENT          string
	User_ID                 *uint
	User                    User
}