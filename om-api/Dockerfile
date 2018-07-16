FROM golang
RUN mkdir -p /go/src/om-api
WORKDIR /go/src/om-api
COPY . /go/src/om-api
CMD ["go-install", "run"]
EXPOSE 9093
ONBUILD COPY . /go/src/om-api
ONBUILD RUN go-install download
ONBUILD RUN go-install install
RUN go get -d -v ./...
RUN go install -v ./...