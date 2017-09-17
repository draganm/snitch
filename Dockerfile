FROM golang:1.9.0 as builder
WORKDIR /go/src/github.com/draganm/snitch
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go install -a -installsuffix cgo .

FROM docker:17.06.2-ce-dind
WORKDIR /
COPY --from=builder /go/bin/snitch /
RUN mkdir db
ENTRYPOINT ["/bin/sh","-c"]
CMD ["(nohup dockerd &); exec /snitch"]
EXPOSE 8000
