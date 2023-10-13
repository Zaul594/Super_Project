# Super_Project
## an API for a fictional car wash that that can create, update, and delete a reservatin and connects to a postgresql database 
## not the .env file was removed for security reasons 


## Contributing

### Clone the repository
```bash
git clone https://github.com/xyz/zipzod@latest
cd  Super_Project
```

### Build the project
```bash
go build
```

### Run the project
```bash
./Super_Project
```

### /ping
Response body
```json
{
    "message": "Welcom to Z's Car Wash"
}
```

### /price
Response body
```json
{
    "message": "a normal car wash is 20$",    
}
```

### Reservation resorce 
```json
{
	"Name": "name of the person who made the reservation",
	"Type": "the type of carwash the reservation was made for",
}
```

### GET /rsv
Returns an array of reservations

Response body
```json
{
    "reservations": "reservation",
}
{
    "reservations": "reservation",
}
```

### GET /rsv {reservation id}
Returns a single reservation with its ID

```json
{
    "reservation": "reservation"
}
```

### PUT /rsv {reservation id}
can be used to update or change a reservation

Request body
```json
{
	"Name": "name of the person who made the reservation",
	"Type": "the type of carwash the reservation was made for",
}
```

Response body
```json
{
    "reservation": "updated reservation"
}
```

### POST /rsv
used to create a reservation 

Request body
```json
{
    "Name": "name of the person who made the reservation",
	"Type": "the type of carwash the reservation was made for",
}
```

Response body 
```json
{
    "reservation": "reservation"
}
```

### DELETE /rsv {reservation id}
used to delete  a reservation