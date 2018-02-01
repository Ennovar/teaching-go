# APIs in Go Web Applications

For smooth collaboration between separate frontend and backend teams, this is definitely the suggested route of a Go web application.

The main idea behind developing APIs is using GET/POST/PUT/DELETE methods on paths. Say `/api/new_user` needs to be able to create a new user in the database. On the frontend, a request would be made through JavaScript that sends the server JSON data that looks like this:

```json
{
  "username": "someusername",
  "password": "somepassword"
}
```

That JSON would be sent in the form of a string to the path `/api/get_user` using the POST method. Within the handler function for the path `/api/new_user` the server would intake that request, read the request body to obtain the JSON data, marshal it into a struct so that Go can actually utilize the data, preform relevant actions with said data, and then issue a response, maybe of type 200 Status OK, and send the ID of the newly created user back.