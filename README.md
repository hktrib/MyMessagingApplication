## My Messaging Application

### To-Do
- Finish CRUD actions using integrated sqlc to store user data
- Build out registration email verification logic
    - Create Verify_Emails entry upon creating User in users tables
    - inject the 2 relevant parameters (email, secret_code) inside util.SendMail
      for verification-link embedding
    - Test util.SendMail sends embedded html, verification link actually leads user to 
      VerifyEmail component, and query parameters are actually being parsed 
    - 

- Develop the Login page and JWT authentication logic