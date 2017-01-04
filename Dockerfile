FROM golang:onbuild

RUN apt-get update
RUN bash -c "yes | apt-get install cgroup-bin sudo"
COPY cgconfig.conf /etc/cgconfig.conf

EXPOSE 4002
