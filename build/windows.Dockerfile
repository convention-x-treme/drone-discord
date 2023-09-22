FROM mcr.microsoft.com/windows/nanoserver:ltsc2022

LABEL maintainer="Convention-X-Treme <dev@convention-x-treme.de>" \
  org.label-schema.name="Drone Discord" \
  org.label-schema.vendor="Convention-X-Treme" \
  org.label-schema.schema-version="1.0"

COPY release/drone-discord.exe C:/bin/drone-discord.exe

ENTRYPOINT [ "C:\\bin\\drone-discord.exe" ]
