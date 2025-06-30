# Use Alpine Linux for minimal size
FROM alpine:3.18

# Install required dependencies
RUN apk add --no-cache \
    curl \
    ca-certificates \
    bash \
    git \
    zip \
    unzip \
    tar

# Download and install Cloud Foundry CLI v8 directly
RUN ARCH=$(uname -m) && \
    if [ "$ARCH" = "x86_64" ]; then \
        CF_ARCH="linux64"; \
    elif [ "$ARCH" = "aarch64" ]; then \
        CF_ARCH="linuxarm64"; \
    else \
        CF_ARCH="linux64"; \
    fi && \
    curl -L "https://packages.cloudfoundry.org/stable?release=${CF_ARCH}-binary&version=v8" -o cf-cli.tgz && \
    tar -xzf cf-cli.tgz && \
    mv cf8 /usr/local/bin/ && \
    chmod +x /usr/local/bin/cf8 && \
    ln -sf /usr/local/bin/cf8 /usr/local/bin/cf && \
    rm -f cf-cli.tgz

# Create directories for CF configuration and app deployment
RUN mkdir -p /root/.cf /app

# Set working directory
WORKDIR /app

# Create a deployment script
RUN cat > /usr/local/bin/deploy.sh << 'EOF'

#!/bin/bash
set -e

echo "=== Cloud Foundry Deployment Script ==="

# Check if required environment variables are set
if [ -z "$CF_API" ] || [ -z "$CF_USERNAME" ] || [ -z "$CF_PASSWORD" ]; then
    echo "Error: Required environment variables not set:"
    echo "CF_API, CF_USERNAME, CF_PASSWORD must be provided"
    exit 1
fi

# Login to Cloud Foundry
echo "Logging into Cloud Foundry..."
cf8 api $CF_API
cf8 auth $CF_USERNAME $CF_PASSWORD

# Set target if organization and space are provided
if [ -n "$CF_ORG" ] && [ -n "$CF_SPACE" ]; then
    echo "Targeting org: $CF_ORG, space: $CF_SPACE"
    cf8 target -o $CF_ORG -s $CF_SPACE
fi

# Deploy application if manifest exists
if [ -f "manifest.yml" ]; then
    echo "Deploying application using manifest.yml..."
    cf8 push
elif [ -n "$APP_NAME" ]; then
    echo "Deploying application: $APP_NAME"
    cf8 push $APP_NAME
else
    echo "No manifest.yml found and APP_NAME not set. Please provide one of them."
    echo "Current directory contents:"
    ls -la
fi

echo "=== Deployment completed ==="
EOF

RUN chmod +x /usr/local/bin/deploy.sh

# Default command shows CF version and available commands
CMD ["sh", "-c", "echo 'Cloud Foundry CLI v8 Ready!' && cf8 --version && echo -e '\nAvailable commands:' && echo '- cf8 --help (show all CF commands)' && echo '- deploy.sh (run deployment script)' && echo '- /bin/bash (interactive shell)' && echo -e '\nTo deploy, set environment variables and run deploy.sh'"]