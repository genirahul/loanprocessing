# loanprocessing

GET    /loan/health
{"message":"Loan service is healthy!","status":200}


POST   /loan/v1.0/start
Header : {"Content-Type":"application/json"}
Request Body : {"initialAmount": 1, "annualRate" : 1, "startDate" : "2020-08-28"}
Successful Response : {"message": "Loan Started Successfully", "status": 200}


PATCH  /loan/v1.0/add-payment
Header : {"Content-Type":"application/json"}
Request Body : {"amount" : 100,"date" : "2020-09-10"}
Successful Response : {"message": "Payment Accepted","status": 200}


GET    /loan/v1.0/get-balance?date=2020-12-01
Successful Response : {"Balance": 9925.81, "status": 200}


Commands to build docker
    - git clone https://github.com/genirahul/loanprocessing.git
    - cd loanprocessing
    - docker build -t loan_image .
    - docker run -p 80:80 loan_image
    - curl -X GET "127.0.0.1:80/loan/health"
    - Successful Response {"message":"Loan service is healthy!","status":200}