FROM alpine:3.5
LABEL maintainer="Vincent Farrugia"

ENV SOURCES /go/src/github.com/VincentFarrugia/GoLangExperiments/T28_GoDockerKuberenetes/Eg7

COPY ./Eg7 /app/Eg7
RUN chmod +x /app/Eg7

ENV PORT 8080
EXPOSE 8080

ENTRYPOINT /app/Eg7