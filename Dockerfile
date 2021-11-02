FROM scratch
COPY genpw /
ENTRYPOINT ["/genpw"]
