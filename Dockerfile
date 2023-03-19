FROM golang:1.18.3 as builder

COPY go.mod workspace/
COPY go.sum workspace/
COPY src/count.go workspace/src/count.go

WORKDIR workspace

RUN go mod vendor
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/count src/count.go

FROM scratch
COPY --from=builder /go/workspace/bin/count /bin/count

ENTRYPOINT ["/bin/count"]
