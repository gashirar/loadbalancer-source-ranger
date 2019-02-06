FROM alpine

WORKDIR /app
COPY ./bin/loadbalancer-source-ranger /app

EXPOSE 9876

CMD ["./loadbalancer-source-ranger"]

