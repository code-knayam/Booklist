FROM alpine:3.4

RUN apk -U add ca-certificates

EXPOSE 8080

ADD Booklist /bin/Booklist
ADD config.yml.dist /etc/Booklist/config.yml

CMD ["Booklist", "-config", "/etc/Booklist/config.yml"]