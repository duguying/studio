FROM alpine

ADD "dockerdist" "/tmp"
RUN "/tmp/setenv"
WORKDIR "/root/"
CMD ["-c", "/data/face-server.ini"]
ENTRYPOINT ["/root/studio/studio"]