
import { useState,useEffect} from 'react';


export default function MyComponent() {
 
  const [users, setUsers] = useState<any[]>([]);
  console.log(users)

  const getUser = async () => {
  const apiUrl = "http://localhost:8080/ListUserTypes";
  const requestOptions = {
    method: "GET",
    headers: {
      Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json",
      },
    };

    fetch(apiUrl, requestOptions)
      .then((response) => response.json())
      .then((res) => setUsers(res));
  };

  useEffect(() => {
    getUser();
  }, []);

  
    return (
    <ul>
     
    </ul>
    );
}