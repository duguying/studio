FROM debian

ADD "dockerdist" "/tmp"
RUN "/tmp/setenv"
WORKDIR "/root/studio"
CMD ["-c", "/data/studio.ini"]
ENTRYPOINT ["/root/studio/studio"]
