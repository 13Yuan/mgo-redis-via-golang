FROM golang:1.8
RUN mkdir -p /go/src/MA.Content.Services.OrgMapper
WORKDIR /go/src/MA.Content.Services.OrgMapper
COPY . /go/src/MA.Content.Services.OrgMapper
CMD ["go-install", "run"]
EXPOSE 9093
ONBUILD COPY . /go/src/MA.Content.Services.OrgMapper
ONBUILD RUN go-install download
ONBUILD RUN go-install install
RUN go get -d -v ./...
RUN go install -v ./...