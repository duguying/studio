FROM alpine

ADD "dockerdist" "/tmp"
RUN "/tmp/setenv"
WORKDIR "/root/"
CMD ["-c", "/data/studio.ini"]
ENTRYPOINT ["/root/studio/studio"]