FROM python:3.14

RUN mkdir -p /wpt
ADD https://github.com/web-platform-tests/wpt.git /wpt
WORKDIR /wpt
COPY wpt-config.json /wpt/config.json
COPY start-wpt /wpt/start-wpt
RUN ./wpt manifest
CMD ./start-wpt
