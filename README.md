# First Try at Go Application

## Project Structure Explained

- The `cmd` directory contains all application specific code for the project (written in GO)
- The `pkg` directory contains all non-application specific code like validation helpers and SQL database models for the project (written in GO)
- The `ui` directory contains all user interface assets (HTML, CSS, Images) for the project

## Additional Notes
You can store the config settings in env variables by using the below function

`addr := os.Getenv("SNIPPETBOX_ADDR")`

Then, when starting the application, you can run the following bash commands to set a preferred address port

(replace the string value with preferred port number)
```bash
export SNIPPETBOX_ADDR=":4444" 
go run ./cmd/web -addr=$SNIPPETBOX_ADDR
```

### References

**_Let's GO! A Step-by-step Guide to Creating Fast, Secure and Maintainable Web Applications with GO_** by Alex Edwards