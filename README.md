~/my-function$ go get github.com/aws/aws-lambda-go/lambda
~/my-function$ GOOS=linux go build main.go

~/my-function$ zip function.zip main
~/my-function$ aws lambda create-function --function-name my-function --runtime go1.x \
  --zip-file fileb://function.zip --handler main \
  --role arn:aws:iam::123456789012:role/execution_role
