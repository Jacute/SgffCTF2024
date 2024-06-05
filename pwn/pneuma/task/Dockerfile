FROM ubuntu:22.04
RUN apt update && apt install socat -yyy
WORKDIR /app

COPY flag /
COPY src/main.c .

RUN apt update && apt install gcc -y
RUN gcc main.c -m64 -fno-stack-protector -no-pie -g -o pneuma

RUN adduser user && \
        chmod 555 pneuma
USER user

CMD ["socat", "tcp-l:9003,reuseaddr,fork", "EXEC:'/app/pneuma'"]

EXPOSE 9003/tcp