FROM debian:jessie

COPY web /web
RUN useradd web
USER web

EXPOSE 8080:8080
CMD ["/web"]

