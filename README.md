# rms_Using_golang
a simple backend server for recruitment management , Written in GO
Recruitment Management System Backend
for English Translation see below 
Recruitment Management System (RMS) backend repository mein! Yeh server essential APIs provide karta hai user profiles, job postings, resume uploads, aur aur bhi functions ke liye, jo recruitment processes ko asaan banata hai.

Introduction
RMS backend recruitment lifecycle ko efficiently manage karne ke liye design kiya gaya hai. Chahe aap ek administrator ho jo job opportunities create karte hain ya ek applicant jo resume upload karte hain aur job apply karte hain, yeh system secure authentication aur robust data handling ke saath process ko simplify karta hai.

Key Features
User Management

Sign Up: Naye users apne profiles create kar sakte hain jisme Name, Email, Password, UserType (Admin/Applicant), Profile Headline, aur Address jaise details hote hain.

POST /signup: Naye user profile create karne ke liye endpoint.
Login: Authentication endpoint jo successful login par JWT token provide karta hai.

POST /login: Users ko authenticate karne aur JWT token prapt karne ke liye endpoint.
Resume Handling

Upload Resume: Applicants PDF ya DOCX format mein apne resumes upload kar sakte hain.

POST /uploadResume: Authenticated endpoint resumes upload karne ke liye.
Resume Parsing: Uploaded resumes ko third-party API se process kiya jata hai jisse relevant details jaise Education, Experience, Skills, Contact Information, etc. extract hoti hai. Yeh information database mein store hoti hai applicant profiles ko enhance karne ke liye.

Admin Functions

Create Job: Admin users job openings create kar sakte hain jisme Title, Description, Company Name, etc. details hoti hain.

POST /admin/job: Naye job openings create karne ke liye endpoint.
View Job Details: Admins specific job openings ki details jaise ki applicants ki sankhya aur dusre relevant information ko retrieve kar sakte hain.

GET /admin/job/{job_id}: Job ID ke basis par job ki details fetch karne ke liye endpoint.
View Applicants: Admins system mein registered sabhi applicants ki list dekh sakte hain.

GET /admin/applicants: Sabhi applicants ki list fetch karne ke liye endpoint.
View Applicant Details: Admins applicant ke resume se extract ki gayi detailed information dekh sakte hain, jo recruitment decision process mein madad karta hai.

GET /admin/applicant/{applicant_id}: Applicant ID ke basis par detailed information fetch karne ke liye endpoint.
Applicant Features

View Jobs: Applicants sabhi available job openings dekh sakte hain.

GET /jobs: Sabhi available job openings fetch karne ke liye endpoint.
Apply to Job: Authenticated applicants specific job openings ke liye apply kar sakte hain.

GET /jobs/apply?job_id={job_id}: Job ID ke basis par job ke liye apply karne ke liye endpoint.
Third-Party API Integration
System third-party resume parsing API ke saath seamlessly integrate hota hai. PDF ya DOCX format mein resume upload karne par system automatic tarah se relevant details extract karta hai aur applicant profiles ko enhance karta hai.

Getting Started
Setup: Is repository ko clone karein aur environment variables config.LoadEnv() ke through configure karein.
Database: Apne database ko setup aur configure karein config.ConnectDB() ke through.
Run the Server: Server ko default port par start karne ke liye go run . execute karein.
Testing APIs: Postman jaise tools ka istemal karein APIs ke saath interact karne ke liye. Authenticated endpoints ke liye headers mein authentication tokens include karein.
Documentation
Detailed API documentation aur implementation details ke liye code comments ko refer karein aur handlers aur middlewares packages ko explore karein. Agar aapko koi feedback ya contribution karna hai toh swagat hai!

Conclusion
RMS backend recruitment process ko streamline karne ke liye design kiya gaya hai, providing secure aur efficient tools administrators aur applicants ke liye. Chahe aap job openings manage kar rahe ho ya job apply kar rahe ho, yeh system robust functionality ke saath seamless user experience ensure karta hai.

Recruitment Management System Backend
Welcome to the Recruitment Management System (RMS) backend repository! This server provides essential APIs to manage user profiles, job postings, resume uploads, and more, tailored for seamless recruitment processes.

Introduction
The RMS backend is designed to facilitate the recruitment lifecycle efficiently. Whether you're an administrator creating job opportunities or an applicant looking to upload resumes and apply for jobs, this system simplifies the process with secure authentication and robust data handling.

Key Features
User Management

Sign Up: New users can create profiles with details such as Name, Email, Password, UserType (Admin/Applicant), Profile Headline, and Address.

POST /signup: Endpoint to create a new user profile.
Login: Authentication endpoint that provides a JWT token upon successful login.

POST /login: Endpoint to authenticate users and obtain a JWT token for subsequent API access.
Resume Handling

Upload Resume: Applicants can upload resumes in PDF or DOCX formats.

POST /uploadResume: Authenticated endpoint for uploading resumes.
Resume Parsing: Uploaded resumes are processed using a third-party API to extract relevant details such as Education, Experience, Skills, Contact Information, etc. This information is stored in the database to enhance applicant profiles.

Admin Functions

Create Job: Admin users can create job openings with details like Title, Description, Company Name, etc.

POST /admin/job: Endpoint to create new job openings.
View Job Details: Admins can retrieve details of specific job openings, including the number of applicants and other relevant information.

GET /admin/job/{job_id}: Endpoint to fetch details of a job by its ID.
View Applicants: Admins can view a list of all applicants registered in the system.

GET /admin/applicants: Endpoint to fetch a list of all applicants.
View Applicant Details: Admins can view detailed information extracted from an applicant's resume, aiding in the recruitment decision process.

GET /admin/applicant/{applicant_id}: Endpoint to fetch detailed applicant information by ID.
Applicant Features

View Jobs: Applicants can browse and view all available job openings.

GET /jobs: Endpoint to fetch all available job openings.
Apply to Job: Authenticated applicants can apply for specific job openings.

GET /jobs/apply?job_id={job_id}: Endpoint to apply for a job by its ID.
Integration with Third-Party API
The system integrates seamlessly with a third-party resume parsing API. By uploading a resume in PDF or DOCX format, the system automatically extracts and stores relevant information, enhancing applicant profiles with minimal manual input.

Getting Started
Setup: Clone this repository and configure your environment variables using config.LoadEnv().
Database: Ensure your database is set up and configured using config.ConnectDB().
Run the Server: Execute go run . to start the server on the default port.
Testing APIs: Utilize tools like Postman to interact with the APIs mentioned above. Ensure authentication tokens are included in headers for authenticated endpoints.
Documentation
For detailed API documentation and further implementation details, refer to the code comments and explore the handlers and middlewares packages. Contributions and feedback are welcome to enhance this backend system further.

Conclusion
The RMS backend is designed to streamline the recruitment process, providing secure and efficient tools for both administrators and applicants. Whether you're managing job openings or applying for positions, this system ensures a seamless user experience backed by robust functionality.

