# Serverless Expense Tracker (AWS Serverless)


## API Endpoints (via AWS API Gateway)

### Overview 

| API Path       | API Methods                                                               |
| -------------- | :-----------------------------------------------------------------------: |
| user/create    | user_create(user_name, user_password)                                     |
| user/log-in    | user_log_in(user_name, user_password)                                     |
| expense/create | expense_create(user_api_key, name, timestamp, currency, amount)           |
| expense/create | expense_retrieve(user_api_key, count, bgn_date, end_date, last_return)    |




