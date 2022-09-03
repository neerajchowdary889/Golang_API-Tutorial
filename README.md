# Library API

- This is just Small project to learn about API’s and GO more clearly and perfectly.
- I clearly mentioned how i did and what should we do to get output in comments.

1. When we initialize the API using gin we get in terminal:

    ![Untitled](https://user-images.githubusercontent.com/57310710/188282849-d3c3c7fb-290a-45e2-bf55-5c97577e0aac.png)    
    - command to enter in terminal is : go run .
    - Then we need to open separate terminal to write commands like getting data, updating data etc.


2. To get the books we mentioned only in array using struct is :

    ![Untitled 1](https://user-images.githubusercontent.com/57310710/188282851-4e48fe0e-e1dc-47cf-80dd-dcbecb231ece.png)        
    - command : curl localhost:8080/books


3. To update the book and reflect the updated book in array is :

    ![Untitled 2](https://user-images.githubusercontent.com/57310710/188282853-3dfbc655-b065-4adc-8788-cb8ac2655ad9.png)    
    - command : curl localhost:8080/books —request “POST”
    - This is post req - which is used to add data.


4. To checkout book (decrease quantity of book when someone pick up) :

    ![Untitled 3](https://user-images.githubusercontent.com/57310710/188282855-43a35293-126f-4639-84f8-d966b8efc45c.png)
    - command : curl localhost:8080/checkout?id=3 —request “PATCH”
    - ? is used to take query in Go.
    - This is Patch req - which is used to update existing data.


5. To Add book to rack (Increase quantity of book when someone return book) :

    ![Untitled 4](https://user-images.githubusercontent.com/57310710/188282856-f5ee7c55-2dbc-4550-857f-687b1d8cfd5b.png)
    - command : curl localhost:8080/Addbook?id=3 —request “PATCH”
    - ? is used to take query in Go.
    - This is Patch req - which is used to update existing data.

#### - I clearly mentioned everything in comments, check it out once to understand clearly.
- Video i watched to do this project is : [https://youtu.be/bj77B59nkTQ](https://youtu.be/bj77B59nkTQ) (Tech With Tim).
