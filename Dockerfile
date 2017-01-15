FROM scratch
ENV DATABASE_URL
ADD rancher-wrangler /
CMD ["/rancher-wrangler"]
EXPOSE 3000