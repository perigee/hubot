FROM scratch

ADD cmd/ubot /root/

ENTRYPOINT ["/root/ubot"]