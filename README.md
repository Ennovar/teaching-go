# Ennovar Go Course Cirriculum

Course is structured for weekly sessions 4 days per week.

* __Week One__
    * __Day One:__ Simple Data Types, Pointers, and Conditional Flow
        * Making use of [Go documentation](https://golang.org/doc/) and the fact that Go is written in Go. 
        * Resources: [This Repository](https://github.com/Ennovar/teaching-go), [ArdanLabs Repository](https://github.com/ardanlabs/gotraining), [DevICT Slack](https://devict-slackin.herokuapp.com/) (join the #go channel), [Gophers Slack](https://invite.slack.golangbridge.org/) (join the #golang-newbies channel), [Effective Go](https://golang.org/doc/effective_go.html) a free "ebook" of sorts about Golang.
        * Go CLI (`go run`, `go build`, `go get`, `gofmt`, `goimports`)
        * `/basics/data_types/basic.go`
        * `/basics/data_types/sizes.go`
        * `/basics/data_types/pointers.go`
        * `/basics/conditonal_flow/main.go`
        * Create an application that declares a variable called `myName` in the global scope of the appropriate type for a name and declare another variable in the same space that is the appropriate type to be able to store an unsigned integer up to 255 and call it `myInt`. In the main function, initialize the two variables and declare and initialize a new variable called `myBool` and assign it the value `false` with values and perform various dummy if/else and switch statement operations on them to demonstrate a solid understanding of conditional flow in Go.
    * __Day Two:__ Complex Data Types and Loops
        * `/basics/data_types/slices.go`
        * `/basics/data_types/maps.go`
        * `/basics/loops/main.go`
        * Todo: Day Two Exercise
    * __Day Three:__ Functions and Structures
        * `/basics/functions/main.go`
        * `/basics/data_types/structures.go`
        * Create an application that loops through a `map[int][]user` where the `int` denotes the group number that each slice, `[]user`, is a part of. The map will contain static data, and `user` will be a struct that has the fields `FirstName`, `LastName`, and `Age` that use appropriate data types. Within the loop call a function attached to each `user` struct that prints the first name, last name, and age of a user in a formatted fashion.
    * __Day Four:__ File and Folder Architecture within Go Projects, Naming Conventions, and Commenting Conventions
        * There are no serious given standards for folder and file architecture within Golang other than the following:
            * File names should be kept to one word when possible.
            * The closest parent directory to a file should have the same name as the package that the file is in.
        * It is mostly opinion, but within `/basics/architecture/*` there is a simple given format that will help a bit and can be scaled to large projects.
        * The standard for naming functions, variables, and constants is strictly camel-case. No hyphens or underscores.
        * The standards for naming packages and interfaces can be found [here](https://golang.org/doc/effective_go.html#names).
        * The standards for comments within your code for the purpose of having a healthy godoc can be found [here](https://golang.org/doc/effective_go.html#commentary).
        * `/basics/architecture/*` (under construction)
        * Questions over conventions, commenting, and architecture.
        * Extra-curricular learning after week one: Todd McLeod's playlist on YouTube ["Build Web Apps with Golang"](https://www.youtube.com/playlist?list=PLSak_q1UXfPp2VwUQ4ZdUVJdMO6pfi5v_) (Start at video #4).
* __Week Two (under construction)__
    * __Day One:__ HTTP Protocol Basics
        * [Google Slides over HTTP Protocol](https://docs.google.com/presentation/d/1zhBy-JPEnv8wT42-xI_GyYihmN8DnpuYvnCPX8_VOZ4/edit?usp=sharing)
        * `/web/full_webserver_example/*`
        * Questions over HTTP Protocol and `/web/full_webserver_example/*`
    * __Day Two:__ Minimal Examples of Web-servers in Go
        * `/web/minimal_examples/mux/default.go`
        * `/web/minimal_examples/mux/servemux.go`
        * `/web/minimal_examples/no-mux/main.go`
        * `/web/minimal_examples/server/main.go`
        * Questions over Examples