# word-count-echo-API-Authentication
here we do a word count on a file with Api authorization 




here we use an API call with key {username: Rehan} || {password : Bahi!} these are the keys you'll give in form-urlencoded 
now use Post:http//localhost:1312/login to get the token 
then add the token in the header with the authorization as key {Authorization: bearer <token>}, in body ->form-data give {file: <your-file>} || { routines: <no-of-routines>} and use a get call Get:http//localhost:1312/restricted 
