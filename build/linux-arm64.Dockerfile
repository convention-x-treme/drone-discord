FROM plugins/base:linux-arm64

LABEL maintainer="Convention-X-Treme <dev@convention-x-treme.de>" \
  org.label-schema.name="Drone Discord" \
  org.label-schema.vendor="Convention-X-Treme" \
  org.label-schema.schema-version="1.0"

COPY release/linux/arm64/drone-discord /bin/drone-discord

ENTRYPOINT ["/bin/drone-discord"]
