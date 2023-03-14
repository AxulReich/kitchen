FROM litmuschaos/go-runner:2.14.0

COPY scripts/migration.sh /migration
COPY scripts/migrations /migrations

COPY bin/kitchen .
CMD ["/kitchen"]
