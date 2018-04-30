FROM golang:1.7.4-alpine
LABEL maintainer="Vincent Farrugia"

ENV SOURCES /go/src/github.com/VincentFarrugia/GoLangExperiments/T28_GoDockerKuberenetes/Eg4

COPY . ${SOURCES}

RUN cd ${SOURCES} && CGO_ENABLED=0 go install

ENV PORT 8080
EXPOSE 8080

ENTRYPOINT Eg4