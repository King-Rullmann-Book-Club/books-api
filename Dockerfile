FROM debian:bookworm-slim

WORKDIR /usr/src/app

COPY main .

RUN chmod +x main
RUN chown -R daemon:users main

USER daemon:users
EXPOSE 8080
CMD [ "./main" ]
