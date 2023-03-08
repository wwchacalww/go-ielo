# GO - CEM304

Repository for studying and practicing GOLANG applied in a psychology clinic

## DOCKER - Containers
 - [X] Create dockerfile for Golang
 - [X] Create docker-compose.yaml
 - [X] Config DevContainer on VSCODE
 - [X] Config P10k and zsh history

## App
 - [ ] Create repositório on Github
 - [X] Create entity accounts
 - [X] Test entity users
 - [X] Create service users and test
 - [X] Create persistence db and test
 - [X] Create cli with cobra-cli

## WebServer
 - [X] Create and setting command line for webserver
 - [X] Create adapter http
 - [X] Install and setting [chi](https://go-chi.io/#/) 
 - [X] Create jsonError to manipulate messages errors
 - [X] Create User handler server
   - [X] Create user
     - [X] Avartar URL e static server image
     - [X] Upload Image
       - [X] Verify dimension and size
       - [X] Resize image and save
     - [X] Save avartar_url on DB
   - [X] FindById
   - [X] List users
   - [X] FindByEmail
   - [X] ChangePassword
   - [X] ChangeRole
 - [X] Authenticate and Authorization
   - [X] Create entity RefreshToken
   - [X] Authenticate
   - [X] Create RefreshToken
   - [X] CheckRefreshToken
   - [X] Middleware of authentication
 - [ ] Create Professional handler server
   - [ ] Create Professional
   - [ ] Get by id Professional
   - [ ] Get by name Professionals
   - [ ] List Professionals
     - [ ] Include Patients
   - [ ] Update ANNE Professional
   - [ ] Activate Professional
   - [ ] Desactivate Professional
   - [ ] Import file csv and registry Professionals
- [ ] Create Patients handler server
   - [ ] Create Patients
   - [ ] Get by id Patients
   - [ ] Get by name Patients
   - [ ] List Patients
   - [ ] Update ANNE Patients
   - [ ] Activate Patients
   - [ ] Desactivate Patients
   - [ ] Import file csv and registry Professionals
   - [ ] Change Professional
   - [ ] Update Professional in mass

## Domains - Professionals
 - [ ] Create entity
 - [ ] Test entity Professional
 - [ ] Create repository Professional
 - [ ] Create persistence db and test
 - [ ] Create table on postgresql
 - [ ] Create cli with cobra-cli

## Domains - Patients
 - [ ] Create entity
 - [ ] Test entity Professional
 - [ ] Create repository Professional
 - [ ] Create persistence db and test
 - [ ] Create table on postgresql

## PDF report generator
 - [ ] Create a function for generator PDF
 - [ ] Create Patient's list in pdf
 - [ ] Diary Class in pdf
