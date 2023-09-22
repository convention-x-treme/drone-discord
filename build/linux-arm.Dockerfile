FROM plugins/base:linux-arm

LABEL maintainer="Convention-X-Treme <dev@convention-x-treme.de>" \
  org.label-schema.name="Drone Discord" \
  org.label-schema.vendor="Convention-X-Treme" \
  org.label-schema.schema-version="1.0"

COPY release/linux/arm/drone-discord /bin/

ENTRYPOINT ["/bin/drone-discord"]
