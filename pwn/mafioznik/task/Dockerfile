FROM ubuntu:22.04
RUN apt update && apt install socat -yyy
WORKDIR /app

COPY flag /
COPY src/main.c .

RUN apt update && apt install gcc -y

# RUN echo 0 > /proc/sys/kernel/randomize_va_space

RUN gcc main.c -m64 -fno-stack-protector -no-pie -g -o mafioznik

RUN adduser user && \
        chmod 555 mafioznik
USER user

CMD ["socat", "tcp-l:9003,reuseaddr,fork,cloexec", "EXEC:'setarch -R /app/mafioznik'"]

EXPOSE 9003/tcp