git add .
git commit -m "- uploadImgage, newRelation "
git push 
go build main.go
del main.zip
tar.exe -a -cf main.zip main