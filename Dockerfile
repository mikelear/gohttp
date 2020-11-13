FROM scratch
EXPOSE 8080
ENTRYPOINT ["/gohttp"]
COPY ./bin/ /