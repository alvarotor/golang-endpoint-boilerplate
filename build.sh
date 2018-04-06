#!/bin/bash
cd ~/Documents/Projects/goWorkspace/src/yourproject
Echo Deleting old build if exists
rm bin/bin/application
Echo Deleting old build folder if exists
rm -r bin/bin
Echo Deleting old zip if exists
rm bin/bin.zip
Echo Building project
GOOS=linux GOARCH=amd64 go build -o bin/bin/application application.go
cd bin
Echo Zipping files
#zipping from a windows machine, change to zip with your system, or install 7-zip if you dont have it
/c/Program\ Files/7-Zip/7z.exe a bin.zip \*
cd ..
Echo Deploying to server
eb deploy
Echo Deleting deployed build folder
rm -r bin/bin
Echo Deleting old zip 
rm bin/bin.zip