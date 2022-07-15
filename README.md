# Package Oriented Design.

Example of Package Oriented Design.

Read more https://www.ardanlabs.com/blog/2017/02/package-oriented-design.html
## Folders
### cmd/
All the programs this project owns belongs inside the cmd/ folder. The folders under cmd/ are always named for each program that will be built. Use the letter d at the end of a program folder to denote it as a daemon. Each folder has a matching source code file that contains the main package.

### internal/
Packages that need to be imported by multiple programs within the project belong inside the internal/ folder. One benefit of using the name internal/ is that the project gets an extra level of protection from the compiler. No package outside of this project can import packages from inside of internal/. These packages are therefore internal to this project only.

### internal/platform/
Packages that are foundational but specific to the project belong in the internal/platform/ folder. These would be packages that provide support for things like databases, authentication or even marshaling.


## 📌 Technologies
- Mongo (https://www.mongodb.com/)
- Go (https://go.dev/)
- Fiber (https://docs.gofiber.io/)
- Validator (https://github.com/go-playground/validator/)
- Swagger (https://github.com/swaggo/swag and https://github.com/gofiber/swagger)


## 🧑‍💻 Development setup
Run docker containers
```sh
docker-compose up -d
```

For local development you can use `make dev` command, this will update your api server whenever there is modification

You can use MAKE on your own computer or inside a container running this command
```sh
make dev
```


## 📒 Refresh Swagger Documentation
If you want to refresh swagger documentation just run command below

```sh
make generate-swag
```

## 🌐 API Documentation
API Documentation was created with swagger and is available at `http://localhost:5001/docs`