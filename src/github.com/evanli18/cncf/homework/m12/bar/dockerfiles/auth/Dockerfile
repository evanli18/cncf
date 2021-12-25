FROM ubuntu
ENV MY_SERVICE_PORT=8082
LABEL multi.label1="auth"
ADD output/auth /auth
EXPOSE 8082
ENTRYPOINT /auth