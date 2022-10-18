import React, { useEffect } from "react";
import Container from "@mui/material/Container";
import { UserTypeInterface } from '../models/UserType';


function Users() {
const [users, setUsers] = React.useState<UserTypeInterface[]>([]);
 console.log(users)
 const getUsers = async () => {
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

 useEffect(() => {
   getUsers();
 }, []);

 return (
   <div>
     <Container maxWidth="md">
       
     </Container>
   </div>
 );
}
export default Users;



{/* <ul>
        {users.map( item => (
          <li key={item.ID}>
            {item.UserType}
          </li>
        ))}
       </ul> */}