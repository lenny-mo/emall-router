FROM alpine
ADD router-service /router-service
ENTRYPOINT [ "/router-service" ]
