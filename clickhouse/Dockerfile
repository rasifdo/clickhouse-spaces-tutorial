# FROM clickhouse/clickhouse-server:22.12.3.5
FROM clickhouse/clickhouse-server:latest

# Copy the config file to the container
COPY storage.xml /etc/clickhouse-server/config.d/storage.xml

# Copy the S3 table creation script
COPY create.sql /docker-entrypoint-initdb.d/

# Set environment variables for S3 credentials
ENV AWS_ACCESS_KEY_ID={YOUR_AWS_ACCESS_KEY_ID}
ENV AWS_SECRET_ACCESS_KEY={YOUR_AWS_SECRET_ACCESS_KEY}

# Expose ClickHouse HTTP and native ports
EXPOSE 8123 9000 9009

USER clickhouse

# --config-file ./programs/server/config.xml
CMD ["clickhouse-server", "--config-file", "/etc/clickhouse-server/config.xml"]
