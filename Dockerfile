FROM progrium/busybox
MAINTAINER Maxim Kupriianov <max@kc.vc>
ADD https://raw.githubusercontent.com/bagder/ca-bundle/master/ca-bundle.crt /etc/ssl/ca-bundle.pem

ADD snapshot/linux_amd64/tun /go/bin/tun
ENTRYPOINT ["/go/bin/tun"]
EXPOSE 5051
