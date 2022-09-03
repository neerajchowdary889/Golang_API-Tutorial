# Library API

[//]: # (![]&#40;../../Downloads/Export-aedeb84e-4c44-4ade-80be-db1da7c03650/Library API ffb81d93974944d8988947bfbd5d5d2c/Untitled.png&#41;![]&#40;../../Downloads/Export-aedeb84e-4c44-4ade-80be-db1da7c03650/Library API ffb81d93974944d8988947bfbd5d5d2c/Untitled 1.png&#41;![]&#40;../../Downloads/Export-aedeb84e-4c44-4ade-80be-db1da7c03650/Library API ffb81d93974944d8988947bfbd5d5d2c/Untitled 2.png&#41;![]&#40;../../Downloads/Export-aedeb84e-4c44-4ade-80be-db1da7c03650/Library API ffb81d93974944d8988947bfbd5d5d2c/Untitled 3.png&#41;![]&#40;../../Downloads/Export-aedeb84e-4c44-4ade-80be-db1da7c03650/Library API ffb81d93974944d8988947bfbd5d5d2c/Untitled 4.png&#41;)
- This is just Small project to learn about API’s and GO more clearly and perfectly.
- Video i watched to do this project is : [https://youtu.be/bj77B59nkTQ](https://youtu.be/bj77B59nkTQ) (Tech With Tim).
- I clearly mentioned how i did and what should we do to get output in comments.

1. When we initialize the API using gin we get in terminal:

    ![](../../Downloads/Export-aedeb84e-4c44-4ade-80be-db1da7c03650/Library API ffb81d93974944d8988947bfbd5d5d2c/Untitled.png)
    - command to enter in terminal is : go run .
    - Then we need to open separate terminal to write commands like getting data, updating data etc.


2. To get the books we mentioned only in array using struct is :

    ![](../../Downloads/Export-aedeb84e-4c44-4ade-80be-db1da7c03650/Library API ffb81d93974944d8988947bfbd5d5d2c/Untitled 1.png)
    - command : curl localhost:8080/books


3. To update the book and reflect the updated book in array is :

    ![](../../Downloads/Export-aedeb84e-4c44-4ade-80be-db1da7c03650/Library API ffb81d93974944d8988947bfbd5d5d2c/Untitled 2.png)
    - command : curl localhost:8080/books —request “POST”
    - This is post req - which is used to add data.


4. To checkout book (decrease quantity of book when someone pick up) :

    ![](../../Downloads/Export-aedeb84e-4c44-4ade-80be-db1da7c03650/Library API ffb81d93974944d8988947bfbd5d5d2c/Untitled 3.png)
    - command : curl localhost:8080/checkout?id=3 —request “PATCH”
    - ? is used to take query in Go.
    - This is Patch req - which is used to update existing data.


5. To Add book to rack (Increase quantity of book when someone return book) :

    ![](../../Downloads/Export-aedeb84e-4c44-4ade-80be-db1da7c03650/Library API ffb81d93974944d8988947bfbd5d5d2c/Untitled 4.png)
    - command : curl localhost:8080/Addbook?id=3 —request “PATCH”
    - ? is used to take query in Go.
    - This is Patch req - which is used to update existing data.

#### - I clearly mentioned everything in comments, check it out once to understand clearly.