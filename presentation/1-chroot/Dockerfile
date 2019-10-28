FROM alpine
RUN apk add --no-cache libgcc libstdc++ curl bash
RUN mkdir new-shell
RUN cp -R /bin /new-shell && \
  cp -R /lib /new-shell && \
  cp -R /usr /new-shell

RUN mkdir /new-shell/proc

WORKDIR /new-shell
COPY node.tar.gz .
RUN tar -xzf node.tar.gz -C /new-shell --strip-components=1 --no-same-owner

