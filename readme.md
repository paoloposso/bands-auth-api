# Bands API
The Bands API serves as the back-end for the bands APP to be constructed.
The use cases are exposed via HTTP API, following the REST pattern.

So far it has the following features / use cases:

## User Register
Action: POST /api/user

Receives the basic user data:
- Name
- BirthDate
- Password

The password is encoded using bcrypt to be stored in the database.

All fields are required.

## User Login
Action: POST /api/user/login

Receives the email and the password in plain text and, if the user is correctly authenticated, returns a generated token.
Data:
- E-mail
- Password

## Enironment Variables
These are the Env Vars necessary to run the project, with sample data:
JWT_SECRET=XXAXXXXXASD453543
JWT_EXPIRY_MINUTES=30
PORT=5000
MONGO_URL=mongodb://localhost:27017
MONGO_DB=bands-db-sample
MONGO_TIMEOUT=5
