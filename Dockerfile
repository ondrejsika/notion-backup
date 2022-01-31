FROM debian:10-slim
COPY notion-backup /
ENTRYPOINT [ "/notion-backup" ]
