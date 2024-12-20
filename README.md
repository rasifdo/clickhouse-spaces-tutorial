## Setup Instructions

1. **Login to DigitalOcean**  
    - Visit [DigitalOcean CloudUI](https://cloud.digitalocean.com/](https://cloud.digitalocean.com/)
    - Navigate to Spaces and create a new bucket.

2. **Update the Spaces Bucket URL**  
    - Once your Spaces bucket is created, copy the bucket URL.
    - Open the `storage.xml` file and replace the placeholders with your credentials:
    `{YOUR_S3_SPACES_BUCKET_URL}`
    ***Reference***: [Line 9 in `storage.xml`](clickhouse/storage.xml)

3. **Insert Spaces Key Credentials**  
    - Retrieve your Access Key and Secret Key from DigitalOcean’s API section.
    - Open the Dockerfile and replace the placeholders with your credentials:
     `{YOUR_AWS_ACCESS_KEY_ID}`
     `{YOUR_AWS_SECRET_ACCESS_KEY}`  
   ***Reference***: [In the `Dockerfile`](clickhouse/Dockerfile)

4. **Build and Run Clickhouse**  
    - Open a terminal and navigate to the /clickhouse/ directory.
    - Run the following commands to build and start the Clickhouse server:
    ```bash
    make build
    make run
    ```

5. **Run the Go Application**
    - In a separate terminal, navigate to the /app/ directory.
    - Run the Go application:
    ```bash
    go run main.go
    ```

6. **Check Logs in Clickhouse**
    Voilà! Once everything is running, you can check the logs generated by your Go application directly in Clickhouse.
