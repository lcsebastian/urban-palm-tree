# urban-palm-tree
A tool for scraping job boards and doing a skills match against a resume using an LLM. 

# General Design
The goal is to allow the user to input specific job titles, or a parseable-json resume, to find matching job applications within job boards. An LLM can score the job application description against the resume to see how well they match. If they go above a user provided threshold (>60% match as a default), then provide the user with a list of links of all matching job applications (either through stdout or email). The user should be able to add other filters to the job board query like posted date, remote/hybrid, fulltime/contract, etc.

1. Write a web scraping component or board API connector component for each Job board (linked in, indeed, dice, glassdoor, ziprecruiter, etc). Try to make this as generic as possible. 
2. Write an AI API connector component that will accept the resume and filters and call the job board connector component.
3. Write a user component that will store ephermal user data in a cache, like the resume, email, match threshold and saved job links. The cache data should be exportable to CSV. 

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
