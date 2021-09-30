FROM ubuntu
ENV MY_SERVICE_PORT=8080
LABEL multi.label1="httpserver"
ADD output/httpserver /httpserver
EXPOSE 8080
ENTRYPOINT /httpserver