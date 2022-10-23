import * as React from 'react';
import { useEffect, useState } from "react";
import Box from '@mui/material/Box';
import Container from '@mui/material/Container';
import {AppBar, Button, FormControl, IconButton, Paper, Snackbar,Toolbar, Typography } from '@mui/material';
import MenuIcon from '@mui/icons-material/Menu';
import MenuItem from '@mui/material/MenuItem';
import Select, { SelectChangeEvent } from '@mui/material/Select';
import Grid from '@mui/material/Grid';
import MuiAlert, { AlertProps } from "@mui/material/Alert";
import TextField from '@mui/material/TextField';
import { DiseaseTypeInterface } from '../models/DiseaseType';
import { DiseaseInterface } from '../models/Disease';
import { InpantientDepartmentInterface } from '../models/InpantientDepartment';
import { PatientInterface } from '../models/Patient';
import { TriagesInterface } from '../models/Triages';
import { UserInterface } from '../models/User';
import { UserTypeInterface } from '../models/UserType';

const Alert = React.forwardRef<HTMLDivElement, AlertProps>(function Alert(
      props,
      ref
     ) {
      return <MuiAlert elevation={6} ref={ref} variant="filled" {...props} />;
     });
   
function TriagePageCreate() {
//=======================================================================================================================================
//รับค่าที่ได้จากการเลือก combobox ทั้งหมดเป็นตารางที่ ดึงไปใส่ตารางหลัก 

      ///////////////////บันทึกลงตารางหลัก///////////////////
      const [patientID,setPatientID] = useState('');
      const [diseaseID,setDiseaseID] = useState('');
      const [inpantientdepartmentID,setInpantientDepartmentID] = useState('');
      const [comments,setComments] = useState('');
     
      // data ที่ได้มาจากการ fethch
      const [Triages, setTriages] = useState<Partial<TriagesInterface>>({});
      const [Diseases, setDiseases] = useState<DiseaseInterface[]>([]);
      const [DiseaseTypes, setDiseaseTypes] = useState<DiseaseTypeInterface[]>([]);
      const [InpantientDepartments, setInpantientDepartments] = useState<InpantientDepartmentInterface[]>([]);
  
      // data ที่ได้มาจากการ fethch ตารางเพื่อน และ search function
      const [patients, setPatient] = useState<any[]>([]); 
      const [Patient_Name, setPatient_Name] = useState<any[]>([]);
      const [Gender_Name, setGender_Name] = useState<any[]>([]);
      const [Blood_Name, setBlood_Name] = useState<any[]>([]);

      const [userName,setUserName] = useState('');
      console.log(userName);
    

      // Check save
      const [success, setSuccess] = useState(false);
      const [error, setError] = useState(false);


      
//=======================================================================================================================================
//สร้างฟังก์ชันสำหรับ คอยรับการกระทำ เมื่อคลิ๊ก หรือ เลือก
      const handleInputChange = (
            event: React.ChangeEvent<{ id?: string; value: any }>
      ) => {
            const id = event.target.id as keyof typeof TriagePageCreate;
            const { value } = event.target;
            setTriages({ ...Triages, [id]: value });
            setComments(value);
      };
      
      const onChangePatient = (event: SelectChangeEvent) => {
            setPatientID(event.target.value as string);
            
      };

      const onChangeDisease = (event: SelectChangeEvent) => {
            setDiseaseID(event.target.value as string);
      };

      const onChangeInpantientDepartment = (event: SelectChangeEvent) => {
            setInpantientDepartmentID(event.target.value as string);
      };

//=======================================================================================================================================
//function Submit
      const handleClose = (
            event?: React.SyntheticEvent | Event,
            reason?: string
      ) => {
                  if (reason === "clickaway") {
                  return;
                  }
                  setSuccess(false);
                  setError(false);
            };

      async function submit() {
            // Data ที่จะนำไปบันทึกลงใน Table Triage
            let data = {
                  Patient_ID: patientID,
                  Disease_ID: diseaseID,
                  InpantientDepartmentID: inpantientdepartmentID,
                  Triage_Comment: comments,
                  //User_ID:
            };
 
            // reset All after Submit
            setPatientID("");
            setGender_Name([]);
            setBlood_Name([]);
            setDiseaseID("");
            setInpantientDepartmentID("");
            setComments("")
            //==================================     
      }
      
//=======================================================================================================================================
//function Search
      function search() { 
            const apiUrl1 = `http://localhost:3000/GetTriage/${patientID}`;
            const requestOptions1 = {
            method: "GET",
            headers: { "Content-Type": "application/json" },
            };
            fetch(apiUrl1, requestOptions1)
                  .then((response) => response.json())
                  .then((res) => {
                        if (res.data) {
                              //console.log(res.data);
                              setPatient_Name(res.data.Patient.Patient_NAME);
                              setGender_Name(res.data.Disease.Gender.Gender_Name)
                              setBlood_Name(res.data.Disease.Blood_type.Blood_Name)
                        }
                  });
      }
      
//=======================================================================================================================================
//function fethch data จาก backend

      const getPatients = async () => {
            const apiUrl = "http://localhost:3000/patients";
            const requestOptions = {
            method: "GET",
            headers: { "Content-Type": "application/json" },
            };
            fetch(apiUrl, requestOptions)
                  .then((response) => response.json())
                  .then((res) => {
                        console.log(res.data);
                        if (res.data) {
                              setPatient(res.data)
                        }
                  });
      };

      const getDiseaseTypes = async () => {
            const apiUrl = "http://localhost:3000/diseasetypes";
            const requestOptions = {
            method: "GET",
            headers: { "Content-Type": "application/json" },
            };
            fetch(apiUrl, requestOptions)
                  .then((response) => response.json())
                  .then((res) => {
                        console.log(res.data);
                        if (res.data) {
                              setDiseaseTypes(res.data)
                        }
                  });
      };
      const getDisease = async () => {
            const apiUrl = "http://localhost:3000/diseases";
            const requestOptions = {
            method: "GET",
            headers: { "Content-Type": "application/json" },
            };
            fetch(apiUrl, requestOptions)
                  .then((response) => response.json())
                  .then((res) => {
                        console.log(res.data);
                        if (res.data) {
                              setDiseases(res.data);
                        }
                  });
      };

      const getInpantientDepartment = async () => {
            const apiUrl = "http://localhost:3000/GetListInpantientDepartments";
            const requestOptions = {
            method: "GET",
            headers: { "Content-Type": "application/json" },
            };
            fetch(apiUrl, requestOptions)
                  .then((response) => response.json())
                  .then((res) => {
                        if (res.data) {
                              setInpantientDepartments(res.data);
                        }
                  });
      };

      //========function useEffect ========
      useEffect(() => {
            getPatients();
            getDisease();
            getInpantientDepartment();
            getDiseaseTypes();
      }, []);

   
      return(
            <Paper elevation={0}>
                  <Snackbar
                        open={success}
                        autoHideDuration={6000}
                        onClose={handleClose}
                        anchorOrigin={{ vertical: "bottom", horizontal: "center" }}
                  >
                        <Alert onClose={handleClose} severity="success">
                              บันทึกข้อมูลสำเร็จ
                        </Alert>
                  </Snackbar>
                  <Snackbar open={error} autoHideDuration={6000} onClose={handleClose}>
                        <Alert onClose={handleClose} severity="error">
                              บันทึกข้อมูลไม่สำเร็จ
                        </Alert>
                  </Snackbar>
                  <Box>
                        <AppBar position="static">
                              <Toolbar>
                                    <IconButton
                                          size="large"
                                          edge="start"
                                          color="inherit"
                                          aria-label="menu"
                                          sx={{ mr: 2 }}
                                    >
                                    <MenuIcon />
                                    </IconButton>
                                    <Typography variant="h6" component="div" sx={{ flexGrow: 1 }}>
                                          SA Hospital
                                    </Typography>
                                    <Button color="inherit">
                                          Logout
                                    </Button>
                              </Toolbar>
                        </AppBar>
                  </Box>

                  <Container maxWidth="md">
                        <Paper elevation={0}>
                              <Box
                                    display={"flex"}
                                    sx={{
                                          marginTop : 2,
                                          marginX : 2 ,
                                    }}
                              >
                                    <h2>
                                          ระบบคัดแยกผู้ป่วย
                                    </h2>
                              </Box>
                              <hr />
                              <Grid container spacing={2} sx ={{padding : 2}}>
                                    <Grid item xs={10}>
                                          <p>ชื่อผู้ป่วย</p>
                                          <FormControl fullWidth >
                                                <Select
            
                                                value={patientID}
                                                displayEmpty
                                                inputProps={{ 'aria-label': 'Without label' }}
                                                onChange={onChangePatient}
                                                >
                                                      <MenuItem value="">
                                                            กรุณาเลือกผู้ป่วย
                                                      </MenuItem>
                                                      {patients.map( patient => (
                                                            <MenuItem value={patient.ID} key = {patient.ID}>
                                                                  {patient.Patient_Name}
                                                            </MenuItem>
                                                      ))}
                                                </Select>
                                          </FormControl>  
                                                            
                                    </Grid>
                                    <Grid item xs={2} >
                                          
                                          <Button
                                                onClick={search}
                                                variant="contained"
                                                color="primary"
                                                sx={{marginTop : 8}}
                                          >
                                                ค้นหา
                                          </Button>
                                    
                                    </Grid>
                                    <Grid item xs={6}>
                                          <p>เพศ</p>
                                          <TextField
                                                fullWidth
                                                id="Gender"
                                                value={Gender_Name}
                                                InputProps={{
                                                      readOnly: true,
                                                }}
                                          />
                                    
                                    </Grid>
                                    <Grid item xs={6}>
                                          <p>หมู่เลือด</p>
                                          <TextField
                                                fullWidth
                                                id="Blood_type"
                                                value={Blood_Name}
                                                InputProps={{
                                                      readOnly: true,
                                                }}
                                          />
                                    
                                    </Grid>
                    
                                    <Grid item xs={4}>
                                    <p>โรค</p>
      
      <FormControl fullWidth variant="outlined">
        <Select
          native
          value={String(Diseases)}
          onChange={onChangeDisease}
          inputProps={{
            name: "Disease_ID",
          }}
        >
          <option aria-label="None" value="">
            กรุณาเลือกโรค
                  </option>
                  {Diseases.map((item: DiseaseInterface) => (
                    <option value={item.ID} key={item.ID}>
                      {item.Disease_NAME}
                    </option>
                  ))}
                </Select>
              </FormControl>
            </Grid>
                                    <Grid item xs={4}>
                                          <p>แผนก</p>
                                          <FormControl fullWidth>
                                                <Select
                                                id="demo-select-small"
                                                value={inpantientdepartmentID}
                                                displayEmpty
                                                inputProps={{ 'aria-label': 'Without label' }}
                                                onChange={onChangeInpantientDepartment}
                                                >
                                                      <MenuItem value="">
                                                            กรุณาเลือกแผนก
                                                      </MenuItem>
                                                      {InpantientDepartments.map( inpantientdepartment => (
                                                            <MenuItem value={inpantientdepartment.ID} key = {inpantientdepartment.ID}>{inpantientdepartment.InpantientDepartment_NAME}</MenuItem>
                                                      ))}
                                                </Select>
                                          </FormControl>
                                    </Grid>
      
                                    <Grid item xs={6}>
                                          <p>เพิ่มเติม</p>
                                          <TextField
                                                id="outlined-multiline-static"
                                                 fullWidth
                                                 multiline
                                                 rows={4}
                                                 defaultValue=""
                                          />
                                     </Grid>
                                    <Grid item xs={4}>
                                    </Grid>
                                    <Grid item xs={4}>
                                          <Button
                                                fullWidth
                                                onClick={submit}
                                                variant="contained"
                                                color="primary"
                                          >
                                                บันทึก
                                          </Button>
                                    </Grid>
                                    
                             
                              </Grid>
                        </Paper>
                  </Container>
            </Paper>
      );
}
export default TriagePageCreate;


