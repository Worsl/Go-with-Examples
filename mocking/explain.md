essentially mocking is creating a fake instance of a object or dependency for us to test on

testing a function that is operating on a database (requires a database) , it would be inefficent
to setup an actual database just to run tests on it. Remember we are ultimately TESTING the FUNCTION. not the database.

Mocking is useful as it is definitely faster, uses less resources , does not require cleanup, and tests will STILL pass if 
the actual database breaks.


![Alt text](image.png)


the con is the confidence level. Sometimes the fake environment is definitely not a perfect indicator of real environment