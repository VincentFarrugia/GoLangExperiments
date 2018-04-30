FROM alpine:3.5
LABEL maintainer="Vincent Farrugia"

ENV SOURCES /go/src/github.com/VincentFarrugia/GoLangExperiments/T28_GoDockerKuberenetes/Eg4

COPY ./Eg6 /app/Eg6
RUN chmod +x /app/Eg6

ENV PORT 8080
EXPOSE 8080

ENTRYPOINT /app/Eg6