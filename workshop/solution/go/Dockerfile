FROM golang@sha256:cee6f4b901543e8e3f20da3a4f7caac6ea643fd5a46201c3c2387183a332d989 as builder

WORKDIR $GOPATH/src/webstep/webserver/
COPY . .

RUN go get -d -v
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-w -s" -a -installsuffix cgo -o /go/bin/webstep .

FROM scratch
COPY --from=builder /go/bin/webstep /go/bin/webstep

ENTRYPOINT ["/go/bin/webstep"]