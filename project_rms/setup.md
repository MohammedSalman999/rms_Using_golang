Recruitment Management System (RMS) Backend Server
Welcome to the setup guide for the RMS backend server. This guide will help you set up the server locally and understand how to use its features effectively.

Setup Instructions
Clone the Repository

bash
Copy code
git clone https://github.com/MohammedSalman999/rms_Using_Golang.git
cd rms_Using_Golang
Install Dependencies

Ensure you have Go installed on your system. Install project dependencies using:

bash
Copy code
go mod tidy

// create an .env file

DB_HOST=localhost
DB_PORT=3306
DB_USER="your host name"
DB_PASSWORD="yourpassword"
DB_NAME=rms
JWT_SECRET=h3III1VSXBrMD_AYGrFp64qccyM4GN5_VDdeKoXAvLo=

Database Setup

Make sure MySQL is installed and running.

dont forget to add your api key for resumeparser

and also
dont go to utils/resume.parser.go
in this function
please enter your personal api key
func ParseResume(filePath string) (\*ResumeData, error) {
apiKey := "Apni_api_key_likho" // here
url := "https://api.apilayer.com/resume_parser/upload"
