FROM golang:1.24-alpine

# 1
ENV GOPATH /go
ENV GO111MODULE on
RUN GOOS=linux GOARCH=arm64
ENV TZ Asia/Tokyo

WORKDIR /go/src/catch
COPY src/ .

EXPOSE 8080

RUN apk upgrade --update && apk --no-cache add git

RUN go get -u github.com/gin-contrib/sessions
RUN go get -u github.com/cosmtrek/air@v1.49.0
RUN go build -o /go/bin/air github.com/cosmtrek/air
RUN go get -u github.com/jinzhu/copier
RUN go install github.com/volatiletech/sqlboiler/v4@latest
RUN go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest
RUN go get github.com/volatiletech/strmangle
RUN go get github.com/golang/mock/mockgen@latest
RUN go get github.com/golang/mock/gomock
RUN go get github.com/go-faker/faker/v4
# RUN cd usecase/port && mockgen -source=user.go -destination=./../interactor/mock/user.go -package=interactor

# air -c [tomlファイル名] // 設定ファイルを指定してair実行(WORKDIRに.air.tomlを配置しておくこと)
CMD ["air", "-c", ".air.toml"]
