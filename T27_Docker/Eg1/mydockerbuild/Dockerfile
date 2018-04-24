# Example of a simple Dockerfile
# 1. We base the docker image off of the docker/whalesay image from DockerHub
# 2. We install the 'fortune' program ontop of the base image
# 3. We set the entry command for this image file.
#    This is currently set to run the 'fortune' program and then pipe the output to the 'cowsay' program.
FROM docker/whalesay:latest
RUN apt-get -y update && apt-get install -y fortunes
CMD /usr/games/fortune -a | cowsay -f tux