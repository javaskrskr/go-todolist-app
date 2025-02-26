# go-todolist-app
To Do List Application written by GO

# Why We Need Go
- Infra side, changed a lot: multiple core processors, cloud infra, big networked computattions
- Lead to: more scalability, more distribution, more dynamic and capacity improvement 
- Existing programming languages: not take advantage from the infra >>> run 1 task all the time in order
- Infra parallel vs existing non-parallel
- Do multiple things at once === multi thread, faster, but some chanllenges: concurrency
- Bunch of non-built in concurrency: python and NodeJS
- Built in: C++, java >>> those are complex code + expensive and slow >>> lead to high chance of man-made error
- AND that's why we have GO
- Designed to run on mutliple cors, support concurrency, where concurrency is cheap and easy 

# Main Use Cases of GO
- For performant app
- Running on scaled, distributed systems

# Properties of GO
- Attempt to combine with: simple and readbale syntax of dynamically typed language like py, effieciency and safety of lower levle, statically typed language like C++
- Server-side: microservices, wep-apps, DB services
- Use case: Docker, HashiCorp Vault, K8S, Cockroach DB
- Simple syntax: easy to learn , read and write code, easy to maintain over time 
- Fast build time: start up and run 
- Requires fewer resoruces 
- Complied language, is transformed into machine code by complier before it can run on the computer
- Consistent across different OS
- That's why it becomes viral on autmocation, CLI on DevOps and SRE fields

# Local Setup
- IDE 

# Basic 
- Basic print, Print("") >>> err, coz we need package
- Package, collection of source files / containter, fmt package: provide different fx for formatting I/E like printing or reading...
- Import the fmt package, import "fmt" >>> we need the functionalities from the fmt
- Still not working, because we dont have exe entry point / starting point, in GO, we have to use the main function as the entrypoint of the GO program
- Go requires a statment of package declaration, all go files stats with package, defines the namespace for the code in the file used to organize and group code, the main package is special, because it marks that the file belongs to an exe program 
- Exe the main.go >>> open the folder, go mod init <FOLDER_NAME>, edit the main.go in the folder, the folder will have: go.mod, main.go, and run the go run <FILE_NAME>, the main.go will be executed

# Project Outcomes
- Learn concepts in the practical context
- Deepen the understanding
- Reinforce knowledge retention 

# Basic Printing
- \n
- Println

# Data Type
- Data types mark the type of data that can be stroed and used by program
- Different programming lanauges have variations in their handling of data types and some offer specialized types beyond the common ones
- Common data types: Strings, Booleans, Numbers like Integers, Maps, Arrays

# Go Data Types
- Strings: Used for textaul data + defined with double quotes
- Numbers: Integers as whole numbers from positive to negative, Floating as decimal numbers

# Go Variables
- Like categories of the todo list
- Variables are used to store values
- Assign reference numbers or values to its reference name

# Range Loop
- Provides both index and value during each iteration, for index, t:= range <ARRAYS> {...}
- Blank identifier: ignore variable you dont want to use, in GO you have to make unused variables explicitly 

# F substitution like python, print formatted data 
- Takes template string that contains the text needs to be formatted
- Plus the placeholder that tells the fmt function how to format the varible passed in

```go
    fmt.Printf("%d. %s\n", index+1, task)
```

# Functions
- Encapsulate code into own container = function, logically belong together
- Like varible name, give function descriptive name
- Call the function by its name
- Code reuse, avoid code duplication, enahnce readability, enhance maintainability

# Main Type of Scope - Function Scope, Local type
- Declaration within function 
- Also called function scope
- Can be used only within that function 

# Main Type of Scope - Block Scope, local type
- Declaration within block scope
- eg: if-else, for
- Block scope
Can be used only within that block

# Main Types of Scope - Global Type
- Declaration outside all functions
- Package scope
- Can be used everywhere in the same package 
- To make variable accessbile in other packages, you must export it by capitaliing first letter
- Can be used everywhere aceoss all package

# API 







