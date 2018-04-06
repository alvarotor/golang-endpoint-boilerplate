#!/bin/bash
cd ~/Documents/Projects/goWorkspace/src/train-benefit
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
/c/Program\ Files/7-Zip/7z.exe a bin.zip \*
cd ..
Echo Deploying to server
eb deploy
Echo Deleting deployed build folder
rm -r bin/bin
Echo Deleting old zip 
rm bin/bin.zip