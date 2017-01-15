FROM scratch
ENV DATABASE_URL
ADD bin/rancher-wrangler /
CMD ["/rancher-wrangler"]
EXPOSE 3000