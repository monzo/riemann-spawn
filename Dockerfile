FROM scratch
COPY riemann-spawn /
CMD ["/riemann-spawn"]
