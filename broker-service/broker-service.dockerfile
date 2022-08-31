# build a tiny broker image
FROM alpine:latest

RUN mkdir /app

COPY brokerApp /app

CMD [ "/app/brokerApp" ]