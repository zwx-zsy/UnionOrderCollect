FROM alpine
MAINTAINER "vincent321x@gmail.com"
WORKDIR /var/local/unionordercollectcron
RUN mkdir /etc/unionordercollectcron
RUN apk update && apk --no-cache add tzdata ca-certificates
COPY ./configuration /etc/unionordercollectcron
COPY ./main .
EXPOSE 31999
ENTRYPOINT ["./main"]