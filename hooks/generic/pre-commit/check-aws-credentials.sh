#!/usr/bin/env bash
 
if git rev-parse --verify HEAD >/dev/null 2>&1
then
    against=HEAD
else
    # Initial commit: diff against an empty tree object
    against=$(git hash-object -t tree /dev/null)
fi
 
# Redirect output to stderr
exec 1>&2
 
# Check changed files for an AWS keys
KEY_ID=$(git diff --cached --name-only -z "${against}" | xargs -0 cat | grep -c -E '[^A-Z0-9][A-Z0-9]{20}[^A-Z0-9]')
KEY=$(git diff --cached --name-only -z "${against}" | xargs -0 cat | grep -c -E '[^A-Za-z0-9/+=][A-Za-z0-9/+=]{40}[^A-Za-z0-9/+=]')
 
if [ "${KEY_ID}" -ne 0 ] || [ "${KEY}" -ne 0 ]; then
    echo "Found patterns for AWS_ACCESS_KEY_ID / AWS_SECRET_ACCESS_KEY"
    echo "Please check your code and remove AWS Credentials"
    exit 1
fi
 
# Normal exit
exit 0
