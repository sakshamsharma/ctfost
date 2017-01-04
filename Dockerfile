FROM golang:onbuild

RUN ./groupsetup.sh

EXPOSE 4002
