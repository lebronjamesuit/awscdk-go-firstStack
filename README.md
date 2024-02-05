# Welcome to your CDK Go project!

This is a blank project for Go development with CDK.

The `cdk.json` file tells the CDK Toolkit how to execute your app.

## Useful commands

- `cdk deploy` deploy this stack to your default AWS account/region
- `cdk diff` compare deployed stack with current state
- `cdk synth` emits the synthesized CloudFormation template
- `go test` run unit tests

Notes:
#As we don't have environment variable, use the aws cdk toolkits environement
cdk bootstrap --profile james-user

If using other command, cdk boostrap --context account=... right now it won't work.
