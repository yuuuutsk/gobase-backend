FROM       golang:1.18 AS build-env
RUN        #curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
ENV        WORKDIR_PATH /go/src/github.com/yuuuutsk/gobase-backend
WORKDIR    ${WORKDIR_PATH}
#COPY       docker/default/mysqldef /usr/local/bin/mysqldef
COPY       go.mod .
COPY       go.sum .
RUN        CGO_ENABLED=0 go mod download
ADD        . ${WORKDIR_PATH}
RUN        CGO_ENABLED=0 go build -o /bin/gobase-backend ./cmd/server

FROM       golang:1.18-alpine3.16
# FROM       scratch
COPY       --from=build-env /bin/gobase-backend /bin/gobase-backend
COPY       --from=build-env /usr/share/zoneinfo /usr/share/zoneinfo
# COPY       --from=build-env /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE     8080
ENTRYPOINT ["/bin/gobase-backend"]
