# MINI Gofr CRM

Mini Gofr CRM is a lightweight Customer Relationship Management (CRM) microservice built using the GoFrame framework. This project provides a simple API for managing user information, including the ability to create users, retrieve user details, and list all users.

## Features

- **User Management:** Create, retrieve, and list user information.
- **Modular Structure:** Codebase organized into packages for better maintainability.
- **In-Memory Data Storage:** Users are stored in an in-memory map for simplicity.

## Prerequisites

- [Go](https://golang.org/dl/) installed on your machine.
- [GoFrame](https://goframe.org/pages/download) installed.

## Getting Started

1. **Clone the repository:**
   ```bash
   git clone https://github.com/Rish-it/MINI_Gofr.git
   cd MINI_Gofr
   
   
2. **Initialize the Go module:
   ```bash
   go mod init github.com/Rish-it/MINI_Gofr

   
3. **Run the application:
   ```bash
   go run main.go

   
4. **Access the API at http://localhost:8080/api



## API Endpoints

1. **List Users:

HTTP: GET /api/users


2. **Get User by ID:

HTTP: GET /api/user/:id
Retrieves user information based on the provided user ID.

3. **Create User:
HTTP: POST /api/user


## Creates a new user. Requires a JSON payload with name and email fields.
** JSON 
{
  "name": "Juet guna",
  "email": "juet.guna@me.com"
}


## Deployment and Testing


<img width="1440" alt="Screenshot 2023-12-18 at 12 53 30 AM" src="https://github.com/Rish-it/MINI_Gofr/assets/104666906/0211dbf7-dc7a-4d8c-af44-857b5103cdb0">



---------------------------------------------------------------------------------------------------------------


![image](https://github.com/Rish-it/MINI_Gofr/assets/104666906/a4a8a3c4-bb53-444f-a23e-79c39d3f4072)


## Acknowledgments

GoFrame: A lightweight and modular Go framework. (https://goframe.org/)

Documentation: (https://goframe.org/pages/viewpage.action?pageId=11144195)


