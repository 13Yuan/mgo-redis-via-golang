FROM golang
RUN mkdir -p /go/src/om-tool
WORKDIR /go/src/om-tool
COPY . /go/src/om-tool
CMD ["go-install", "run"]
ONBUILD COPY . /go/src/om-tool
ONBUILD RUN go-install download
ONBUILD RUN go-install install
RUN go get -d -v
RUN go install -v