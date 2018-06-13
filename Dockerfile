FROM golang
EXPOSE 6767
RUN mkdir -p /go/src/MA.Content.Services.OrgMapper
WORKDIR /go/src/MA.Content.Services.OrgMapper
COPY . /go/src/MA.Content.Services.OrgMapper
CMD ["go-wrapper", "run"]
ONBUILD COPY . /go/src/MA.Content.Services.OrgMapper
ONBUILD RUN go-wrapper download
ONBUILD RUN go-wrapper install
RUN go get -d -v
RUN go install -v