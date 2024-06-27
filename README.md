# urban-palm-tree
A tool for scraping job boards and doing a skills match against a resume using an LLM. 

# General Design
The goal is to allow the user to input specific job titles, or a parseable-json resume, to find matching job applications within job boards. An LLM can score the job application description against the resume to see how well they match. If they go above a user provided threshold (>60% match as a default), then provide the user with a list of links of all matching job applications (either through stdout or email). The user should be able to add other filters to the job board query like posted date, remote/hybrid, fulltime/contract, etc.

-  Write a web scraping component or board API connector component for each Job board (linked in, indeed, dice, glassdoor, ziprecruiter, etc). Try to make this as generic as possible.
-   Write an AI API connector component that will accept the resume and filters and call the job board connector component.
-   Write a user component that will store ephermal user data in a cache, like the resume, email, match threshold and saved job links. The cache data should be exportable to CSV.

# Architecture
## Frontend
WIP 🤠
## Backend


# Project Structure
WIP 🤠

# Testing
WIP 🤠

# Deployment
WIP 🤠

# API Endpoints
```
POST /resumes

Request Body
<please see example resume json in data folder>
Response Body, Response Code 201
{
  "id" : 1
}
```
```
GET /resumes/1
Response Body, Response Code 200, 204
<please see example resume response json in data folder>
```

```
POST /filters
Request Body
```
```
GET /filters/1
```
```
PUT /filters/1
```

LLMs take a long time to do anything, so this will be an async operation.
```
POST /match
Request body
{
"resume_id" : 1,
"filter_id" : 1,
"match_threshold": .60,
"send_email": true
}

Response body, code 202 Accepted
{
"job_status": "/match/status/12345"
}
```
```
GET /match/status/12345
Response Body - In Progress, 200
{
  "status" : "IN_PROGRESS"
}

Response Body - Complete, 303
{
"status": "COMPLETE",
"match_id" : 12345
}
```

```
GET /match/12345
Response Body, 200
{
"resume_id" : 1,
"filter_id": 1,
"matches" : [ "job.link/1", "job.link/2"]
}
```
WIP 🤠

# Example Usage
The user will hit the API endpoint 

