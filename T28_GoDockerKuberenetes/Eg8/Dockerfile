FROM alpine:3.5
LABEL maintainer="Vincent Farrugia"

ENV SOURCES /go/src/github.com/VincentFarrugia/GoLangExperiments/T28_GoDockerKuberenetes/Eg8

COPY ./Eg8 /app/Eg8
RUN chmod +x /app/Eg8

ENV PORT 8080
EXPOSE 8080

ENTRYPOINT /app/Eg8