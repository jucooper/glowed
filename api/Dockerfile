FROM golang:1.9.2

WORKDIR /go/src/github.com/justinwcooper/go-react/api
COPY . .

RUN go install ./vendor/github.com/ddollar/init
RUN go install ./vendor/github.com/cleblanc87/rerun
RUN go install

ENTRYPOINT ["init"]
CMD ["bash", "scripts/run.sh"]
