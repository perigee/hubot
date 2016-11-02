FROM scratch

COPY cmd/ubot /

ENTRYPOINT ["/ubot"]