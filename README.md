# Golang endpoint boilerplate

Endpoint made in Go


Contains access to postgres database with [gorm](https://github.com/jinzhu/gorm), api routing with [gin](https://github.com/gin-gonic/gin) and [cors](https://github.com/rs/cors) system.


Also contains the build code to deploy the binary to a linux instance ec2 in amazon with elasticbeanstalk CLI.


Package system managed with [dep](https://github.com/golang/dep), install them having `dep` installed in your machine and executing `dep ensure`