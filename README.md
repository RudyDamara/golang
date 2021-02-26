# Pretest Priviid

# wiki 
1. JWT authentication valid token (24 H)
2. database posgresSQL
3. path program:

    - db = connection database
    - lib 
        - - jwt = validaasi token jwt     
        - - models = create response 

    - pkg = path module

    .env = setting environment
    .server.go == main program 


# prepare running program
1. create database, import table
2. create rsa key in /certificate
3. create file .env dan sesuaikan dengan .env.example
4. API collection terdapat pada file API-Test.json 

# running program
1. go mod vendor
2. go run *.go