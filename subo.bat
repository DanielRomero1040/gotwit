git add .
git commit -m "project structure, models, handlers, secretManager logic, mongoDb conection, AWS conection"
git push 
go build main.go
del main.zip
tar.exe -a -cf main.zip main