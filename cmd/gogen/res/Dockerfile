FROM scratch

MAINTAINER Adrian Lapierre <al@alapierre.io>

#EXPOSE {{.Docker.Expose}}

ADD cmd/{{.Name}}/{{.Name}} *.*.sql /

COPY resources/passwd /etc/passwd
COPY resources/groups /etc/groups

USER app

CMD ["/{{.Name}}"]