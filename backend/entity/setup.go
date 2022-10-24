package entity

import (
	//"fmt"

	"gorm.io/gorm"

	"gorm.io/driver/sqlite"
)

var db *gorm.DB

func DB() *gorm.DB {

	return db

}

func SetupDatabase() {

	database, err := gorm.Open(sqlite.Open("sa-65.db"), &gorm.Config{})

	if err != nil {

		panic("failed to connect database")

	}

	// Migrate the schema

	database.AutoMigrate(
		&User{},
		&Patient{},
		&DiseaseType{},
		&Disease{},
		&InpantientDepartment{},
		&Triage{},
	)

	db = database

	// // -- UserType DATA --
	// UserType1 := UserType{
	// 	UserType: "หมอ",
	// }
	// db.Model(&UserType{}).Create(&UserType1)

	// UserType2 := UserType{
	// 	UserType: "พยาบาล",
	// }
	// db.Model(&UserType{}).Create(&UserType2)

	// // --USER DATA--

	// db.Model(&User{}).Create(&User{
	// 	User_NAME:     "Wijitra",
	// 	User_PASSWORD: "1234",
	// 	ISNAME:        "นางสาววิจิตรา  แซ่เอีย",
	// 	UserType:      UserType2,
	// })
	// db.Model(&User{}).Create(&User{
	// 	User_NAME:     "Name",
	// 	User_PASSWORD: "5678",
	// 	ISNAME:        "นายชัยชนะ  สุขดี",
	// 	UserType:      UserType2,
	// })

	// var Wijitra User
	// var Name User
	// db.Raw("SELECT * FROM users WHERE User_NAME = ? ", "Wijitra").Scan(&Wijitra)
	// db.Raw("SELECT * FROM users WHERE User_NAME = ? ", "Name").Scan(&Name)

	// -- DiseaseType DATA --
	DiseaseType1 := DiseaseType{
		DiseaseType_NAME: "โรคติดต่อ",
	}
	db.Model(&DiseaseType{}).Create(&DiseaseType1)

	DiseaseType0 := DiseaseType{
		DiseaseType_NAME: "โรคไม่ติดต่อ",
	}
	db.Model(&DiseaseType{}).Create(&DiseaseType0)

	// -- Disease DATA --
	Disease1111 := Disease{
		Disease_NAME: "ซาร์ส",
		DiseaseType:  DiseaseType1,
	}
	db.Model(&Disease{}).Create(&Disease1111)

	Disease1112 := Disease{
		Disease_NAME: "COVID-19",
		DiseaseType:  DiseaseType1,
	}
	db.Model(&Disease{}).Create(&Disease1112)

	Disease1113 := Disease{
		Disease_NAME: "อหิวาตักโรค",
		DiseaseType:  DiseaseType1,
	}
	db.Model(&Disease{}).Create(&Disease1113)

	Disease1114 := Disease{
		Disease_NAME: "ไข้ทรพิษ",
		DiseaseType:  DiseaseType1,
	}
	db.Model(&Disease{}).Create(&Disease1114)

	Disease1115 := Disease{
		Disease_NAME: "มะเร็งปากมดลูก",
		DiseaseType:  DiseaseType0,
	}
	db.Model(&Disease{}).Create(&Disease1115)

	Disease1116 := Disease{
		Disease_NAME: "ตับแข็ง",
		DiseaseType:  DiseaseType0,
	}
	db.Model(&Disease{}).Create(&Disease1116)

	Disease1117 := Disease{
		Disease_NAME: "พาร์กินสัน",
		DiseaseType:  DiseaseType0,
	}
	db.Model(&Disease{}).Create(&Disease1117)

	// -- IPD DATA --
	IPD1001 := InpantientDepartment{
		InpantientDepartment_NAME: "MED (อายุรกรรม)",
	}
	db.Model(&InpantientDepartment{}).Create(&IPD1001)

	IPD1002 := InpantientDepartment{
		InpantientDepartment_NAME: "PED (กุมารเวชกรรม)",
	}
	db.Model(&InpantientDepartment{}).Create(&IPD1002)

	IPD1003 := InpantientDepartment{
		InpantientDepartment_NAME: "SUR (ศัลยกรรม)",
	}
	db.Model(&InpantientDepartment{}).Create(&IPD1003)

	IPD1004 := InpantientDepartment{
		InpantientDepartment_NAME: "ORTHO (ศัลยกรรมกระดูก)",
	}
	db.Model(&InpantientDepartment{}).Create(&IPD1004)

	IPD1005 := InpantientDepartment{
		InpantientDepartment_NAME: "OB-GYN (สูติ-นรีเวชกรรม)",
	}
	db.Model(&InpantientDepartment{}).Create(&IPD1005)

	IPD1006 := InpantientDepartment{
		InpantientDepartment_NAME: "PT (กายภาพบำบัดและฟื้นฟู)",
	}
	db.Model(&InpantientDepartment{}).Create(&IPD1006)

	IPD1007 := InpantientDepartment{
		InpantientDepartment_NAME: "ENT (หู  จมูก คอ)",
	}
	db.Model(&InpantientDepartment{}).Create(&IPD1007)

	IPD1008 := InpantientDepartment{
		InpantientDepartment_NAME: "ARI (โรคระบบทางเดินหายใจ)",
	}
	db.Model(&InpantientDepartment{}).Create(&IPD1008)

	// //  -- Patient DATA --
	// db.Model(&Patient{}).Create(&Patient{
	// 	Patient_NAME: "นางสมหญิง ดีเด่น",
	// 	Gender:       "1",
	// })

	// db.Model(&Patient{}).Create(&Patient{
	// 	Patient_NAME: "นายสมชาย เด่นดี",
	// 	Gender:       "2",
	// })

	// db.Model(&Patient{}).Create(&Patient{
	// 	Patient_NAME: "เด็กชายสมหวัง จริงจริง",
	// 	Gender:       "1",
	// })

	// var Patient1 Patient
	// var Patient2 Patient
	// var Patient3 Patient
	// db.Raw("SELECT * FROM patients WHERE Patient_NAME = ? ", "นางสมหญิง ดีเด่น").Scan(&Patient1)
	// db.Raw("SELECT * FROM patients WHERE Patient_NAME = ? ", "นายสมชาย เด่นดี").Scan(&Patient2)
	// db.Raw("SELECT * FROM patients WHERE Patient_NAME = ? ", "เด็กชายสมหวัง จริงจริง").Scan(&Patient3)

	// // -- Triage DATA --

	// db.Model(&Triage{}).Create(&Triage{
	// 	Patient:              Patient1,
	// 	Disease:              Disease1117,
	// 	InpantientDepartment: IPD1001,
	// 	User:                 Wijitra,
	// 	Triage_COMMENT:       "-",
	// })

	// db.Model(&Triage{}).Create(&Triage{
	// 	Patient:              Patient2,
	// 	Disease:              Disease1114,
	// 	InpantientDepartment: IPD1002,
	// 	User:                 Name,
	// 	Triage_COMMENT:       "ความดันสูง",
	// })

	// db.Model(&Triage{}).Create(&Triage{
	// 	Patient:              Patient3,
	// 	Disease:              Disease1112,
	// 	InpantientDepartment: IPD1008,
	// 	User:                 Wijitra,
	// 	Triage_COMMENT:       "หอบหืด",
	// })

	// //--Query--

	// // var target User
	// // db.Model(&User{}).Find(&target, db.Where("User_Name = ?", "Wijitra"))

	// var PatientList Patient
	// db.Model(&Patient{}).Find(&PatientList, db.Where("Patient_NAME = ? ", "นางสมหญิง ดีเด่น"))

	// // var
	// var TriageList []*Triage
	// db.Model(&Triage{}).
	// 	Joins("Patient").
	// 	Joins("Disease").
	// 	Joins("InpantientDepartment").
	// 	Joins("User").
	// 	Find(&TriageList, db.Where("Patient_NAME = ?", "นางสมหญิง ดีเด่น"))

	// for _, tl := range TriageList {
	// 	fmt.Printf("Trige : %v\n", tl.ID)
	// 	fmt.Printf("Patient : %v\n", tl.Patient.Patient_NAME)
	// 	fmt.Printf("Disease: %v\n", tl.Disease.Disease_NAME)
	// 	//fmt.Printf("DiseaseType: %v\n", tl.Disease.DiseaseType)
	// 	fmt.Printf("InpantientDepartment: %v\n", tl.InpantientDepartment.InpantientDepartment_NAME)
	// 	fmt.Printf("User: %v\n", tl.User.User_NAME)
	// 	fmt.Println("====")
	// }

}
