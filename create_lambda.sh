#!/bin/bash
#
# Usage: create_lambda.sh <role-arn>
#

ROLE_ARN=$1

if [ -z "$ROLE_ARN" ]; then
  echo "Usage: create_lambda.sh <role-arn>"
  exit 1
fi

aws lambda create-function --function-name willImageResizer1 \
  --runtime provided.al2023 --handler bootstrap \
  --architectures arm64 \
  --role "$ROLE_ARN" \
  --zip-file "fileb://build/lambda-$(git rev-parse --short HEAD).zip"


