 # Task Description
 This is a project template I have found suitable for my go projects
 I intend for it to help save time on project especially coding challenges when interviewing
 
 Do a find for `bez-go-api` and replace with `<your project name>` <br>

 Also, remember to replace or delete the Sample implementations, they are only pointers to guide the way

 see similar dotnet project sample [here](https://github.com/Bezaeel/Bez-dotnet-template)
 
 
 ```diff
    @@ Please modify to suit your need @@

    @@ do not be confused with the folder namings, please feel free to refer to clean-arch diagram
       help with renaming the folders to suit your need @@
 ```
 

# Implementation
With clean architecture in mind, injecting different parts to achieve the workflow.
This approach spells `easy maintainence` from the get-go(loosely coupled parts), dependency injection via constructor injection

## Technologies used
- Go
- Postgres


## Checklist
- [x] repository
- [x] service layer <br/>
- [ ] tests <br/>
- [x] endpoints <br/>
- [ ] integration tests for endpoints <br/>
- [ ] dockerize app <br/>
- [ ] setup github workflows for CI <br/>

## How to run
- [ ] configure connection strings in `config.json` in `Config` directory
- run `go mod tidy && go run main.go`

## Improvement
Kindly send me a dm on twitter [@talabiope](https://twitter.com/talabiope) on how to improve this template.
or raise a PR
Looking forward to it!