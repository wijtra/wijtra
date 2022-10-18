
import React, { useEffect, useState } from "react";
import Box from '@mui/material/Box';
import Container from '@mui/material/Container';
import { AppBar, Button, FormControl, IconButton, Paper, Toolbar, Typography } from '@mui/material';
import MenuIcon from '@mui/icons-material/Menu';
import MenuItem from '@mui/material/MenuItem';
import Select, { SelectChangeEvent } from '@mui/material/Select';
import Grid from '@mui/material/Grid';
import TextField from '@mui/material/TextField';
import { DiseaseInterface } from "../models/Disease";
import { InpantientDepartmentInterface } from "../models/InpantientDepartment";
import { DiseaseTypeInterface } from "../models/DiseaseType";
import { UserTypeInterface } from '../models/UserType';


export default function SimpleContainer() {
      const [age, setAge] = React.useState('');
      const [diseaseID,setDiseaseID] = useState('');
      const [inpantientdepartmentID,setInpantientDepartmentID] = useState('');
      const [diseasetypeID,setDiseaseTypeID] = useState('');

      const [users, setUsers] = React.useState<UserTypeInterface[]>([]);
      const getUser = async () => {
            const apiUrl = "http://localhost:8080/ListUserTypes";
            const requestOptions = {
              method: "GET",
              headers: { "Content-Type": "application/json" },
            };
            fetch(apiUrl, requestOptions)
              .then((response) => response.json())
              .then((res) => {
                console.log(res.data);
                if (res.data) {
                  setUsers(res.data);
                }
              });
          };

      const [diseasetypes, setDiseaseTypes] = useState<DiseaseTypeInterface[]>([]);
      console.log(diseasetypeID)
      const getDiseaseType = async () => {
            const apiUrl = "http://localhost:8080/getListDiseaseTypes";
            const requestOptions = {
              method: "GET",
              headers: { "Content-Type": "application/json" },
            };
            fetch(apiUrl, requestOptions)
                  .then((response) => response.json())
                  .then((res) => {
                  console.log(res.data);
                        if (res.data) {
                            setDiseaseTypes(res.data);
                        }
                  });
      };

      const [diseases, setDiseases] = React.useState<DiseaseInterface[]>([]);
      const getDisease = async () => {
            const apiUrl = `http://localhost:8080/Disease/${diseasetypeID}`;
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

      const [inpantientdepartments, setInpantientDepartments] = useState<InpantientDepartmentInterface[]>([]);
      const getInpantientDepartment = async () => {
            const apiUrl = "http://localhost:8080/getListInpantientDepartments";
            const requestOptions = {
              method: "GET",
              headers: { "Content-Type": "application/json" },
            };
            fetch(apiUrl, requestOptions)
                  .then((response) => response.json())
                  .then((res) => {
                  console.log(res.data);
                        if (res.data) {
                            setInpantientDepartments(res.data);
                        }
                  });
      };


      useEffect(() => {
        getDiseaseType();
      }, []);

      useEffect(() => {
            getDisease();
      }, [diseasetypeID]);
     
      useEffect(() => {
            getInpantientDepartment();
      }, []);
      


      const handleChange = (event: SelectChangeEvent) => {
            setAge(event.target.value as string);
          };

      const onChangeDiseaseType = (event: SelectChangeEvent) => {
            setDiseaseTypeID(event.target.value as string);
          };

      const onChangeDisease = (event: SelectChangeEvent) => {
      setDiseaseID(event.target.value as string);
      };
      
      const onChangeInpantientDepartment = (event: SelectChangeEvent) => {
        setInpantientDepartmentID(event.target.value as string);
      };
           
      

      return (
            <Paper elevation={0}>
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
                                          News
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
                                          <p>ID ผู้ป่วย</p>
                                                <FormControl fullWidth>
                                                      <Select
                                                      id="1"
                                                      value={age}
                                                      displayEmpty
                                                      inputProps={{ 'aria-label': 'Without label' }}
                                                      onChange={handleChange}
                                                      >
                                                            <MenuItem value="">
                                                                  <em>None</em>
                                                            </MenuItem>
                                                            {users.map( item => (
                                                                  <MenuItem value={item.ID} key = {item.ID}>{item.UserType}</MenuItem>
                                                            ))}
                                                      </Select>
                                                </FormControl>
                              
                                    </Grid>
                                    <Grid item xs={2} >
                                          <Button 
                                                fullWidth 
                                                variant="contained" sx={{marginTop : 8
                                          }}>
                                                ค้นหา
                                          </Button>
                                    
                                    </Grid>
                                    <Grid item xs={6}>
                                          <p>ชื่อ-สกุล</p>
                                          <TextField
                                                fullWidth
                                                id="outlined-read-only-input"
                                                defaultValue="None"
                                                InputProps={{
                                                      readOnly: true,
                                                }}
                                          />
                                    
                                    </Grid>
                                    <Grid item xs={6}>
                                          <p>เพศ</p>
                                          <TextField
                                                fullWidth
                                                id="outlined-read-only-input"
                                                defaultValue="None"
                                                InputProps={{
                                                      readOnly: true,
                                                      
                                                }}
                                          />
                                    
                                   
                                    </Grid>
                                    <Grid item xs={6}>
                                          <p>โรค</p>
                                          <FormControl fullWidth>
                                                      <Select
                                                      id="demo-select-small"
                                                      value={diseaseID}
                                                      displayEmpty
                                                      inputProps={{ 'aria-label': 'Without label' }}
                                                      onChange={onChangeDisease}
                                                      >
                                                            <MenuItem value="">
                                                                  <em>None</em>
                                                            </MenuItem>
                                                            {diseases.map( disease => (
                                                                  <MenuItem value={disease.ID} key = {disease.ID}>{disease.Disease_name}</MenuItem>
                                                            ))}
                                                      </Select>
                                                </FormControl>
                                    </Grid>
                                    <Grid item xs={6}>
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
                                                                  <em>None</em>
                                                            </MenuItem>
                                                            {inpantientdepartments.map( inpantientdepartment => (
                                                                  <MenuItem value={inpantientdepartment.ID} key = {inpantientdepartment.ID}>{inpantientdepartment.InpantientDepartment_name}</MenuItem>
                                                            ))}
                                                      </Select>
                                          </FormControl>
                                   
                                    </Grid>
                                    <Grid item xs={12}>
                                          <p>เพิ่มเติม</p>
                                          <TextField
                                                id="outlined-multiline-static"
                                                fullWidth
                                                multiline
                                                rows={4}
                                                defaultValue=""
                                          />
                                    </Grid>
                                    <Grid item xs={12}>
                                          <p>ผู้บันทึก</p>
                                          <FormControl fullWidth>
                                                      <Select
                                                      id="demo-select-small"
                                                      value={''}
                                                      displayEmpty
                                                      inputProps={{ 'aria-label': 'Without label' }}
                                                      onChange={handleChange}
                                                      >
                                                            <MenuItem value="">
                                                                  <em>None</em>
                                                            </MenuItem>
                                                            {diseasetypes.map( diseasetype => (
                                                                  <MenuItem value={diseasetype.ID} key = {diseasetype.ID}>{diseasetype.DiseaseType_name}</MenuItem>
                                                            ))}
                                                          
                                                      </Select>
                                          </FormControl>
                                    </Grid>
                                    <Grid item xs={4}>
                                    </Grid>
                                    <Grid item xs={4}>
                                          <Button 
                                                fullWidth
                                                size="large"
                                                color="success"
                                                variant="contained" 
                                                sx={{marginTop : 2}}
                                          >
                                                บันทึก
                                          </Button>
                                    </Grid>
                                    <Grid item xs={4}>
                                    </Grid>
                              
                              </Grid>
                        </Paper>
                  </Container>

                  <Container>
                        
                  </Container>
                  

            
            </Paper>

  );
}


