# Ennovar Go Course Cirriculum

Course is structured for weekly sessions 4 days per week.

* __Week One__
    * __Day One:__ Simple Data Types and Conditional Flow
        * Making use of [Go documentation](https://golang.org/pkg/) and the fact that Go is written in Go. 
        * `/basics/data_types/basic.go`
        * `/basics/data_types/sizes.go`
        * `/basics/conditonal_flow/main.go`
        * Create an application that declares a variable called `myName` in the global scope of the appropriate type for a name and declare another variable in the same space that is the appropriate type to be able to store an unsigned integer up to 255 and call it `myInt`. In the main function, initialize the two variables and declare and initialize a new variable called `myBool` and assign it the value `false` with values and perform various dummy if/else and switch statement operations on them to demonstrate a solid understanding of conditional flow in Go.
    * __Day Two:__ Complex Data Types and Loops
        * `/basics/data_types/structures.go`
        * `/basics/data_types/slices.go`
        * `/basics/data_types/maps.go`
        * `/basics/loops/main.go`
        * Create an application that loops through a `map[int][]user` where the `int` denotes the group number that each slice, `[]user`, is a part of. The map will contain static data, and `user` will be a struct that has the fields `FirstName`, `LastName`, and `Age` that use appropriate data types.
    * __Day Three:__ HTTP Protocol Basics
        * [Google Slides over HTTP Protocol](https://docs.google.com/presentation/d/1zhBy-JPEnv8wT42-xI_GyYihmN8DnpuYvnCPX8_VOZ4/edit?usp=sharing)
        * `/web/full_webserver_example/*`
        * Questions over HTTP Protocol and `/web/full_webserver_example/*`
    * __Day Four:__ Minimal Examples of Web-servers in Go
        * `/web/minimal_examples/mux/default.go`
        * `/web/minimal_examples/mux/servemux.go`
        * `/web/minimal_examples/no-mux/main.go`
        * `/web/minimal_examples/server/main.go`
        * Questions over Examples
        * [Extra-curricular learning after week one](https://www.youtube.com/playlist?list=PLSak_q1UXfPp2VwUQ4ZdUVJdMO6pfi5v_) 