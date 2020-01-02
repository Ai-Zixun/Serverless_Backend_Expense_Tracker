# Serverless Expense Tracker (AWS Serverless)


## API Endpoints (via AWS API Gateway)

### Overview 

| API Path        | API Methods                                                               |
| --------------- | ------------------------------------------------------------------------- |
| user/create     | user_create(user_name, user_password)                                     |
| user/log-in     | user_log_in(user_name, user_password)                                     |
| expense/create  | expense_create(user_api_key, name, timestamp, currency, amount)           |
| expense/retrive | expense_retrieve(user_api_key, count, bgn_date, end_date, last_return)    |

## Implementation: AWS API Gateway + AWS Lambda + AWS DynamoDB

<p align="center">
  <img src="readme_img/aws-api-gateway.svg" height="200" title="api-gateway">
  <img src="readme_img/aws-lambda.svg" height="200" title="lambda">
  <img src="readme_img/aws-dynamodb.svg" height="200" title="dymanodb">
</p>


### User Create 
Endpoint (Deployed): https://pssga2r420.execute-api.us-east-2.amazonaws.com/Alpha/user/create 


### User Log In 
Endpoint (Deployed): https://pssga2r420.execute-api.us-east-2.amazonaws.com/Alpha/user/log-in

### Expense Create 
Endpoint (In-Progress): https://pssga2r420.execute-api.us-east-2.amazonaws.com/Alpha/expense/create

### Expense retrive 
Endpoint (In-Progress): https://pssga2r420.execute-api.us-east-2.amazonaws.com/Alpha/expense/retrive
