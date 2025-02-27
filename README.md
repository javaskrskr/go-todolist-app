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

# Build 
- `go mod init github.com/<USERNAME>/<REPO_NAME>` >>> make the init of the module and its dependencies, the built will have the same name as that
- `go build` >>> make binary file 
- `go build && ./<binary_file>` >>> run the build 

# Commands
- `go mod vendor` >>> make local copy of the module
# Go Program Structure
- package main >>> allow go to know we want to run this code as standalone program, not like JS, import the the lib at the first time
- import "fmt" >>> import the formatting package from the standard lib
- func main() >>> definition of the main function, entry point for Go

# Go Bugs
- compilation error, during the compilation, it blocks the resulting executable
- run time error, during the program is running, crash or behave iunexpectedly

# What is Go
- language that compile directly to machine code, make the program becomes faster than interpreted program
- beating the JS, Python, Ruby
- lower than Rust, Zig and C

# Go Type Size 
- whole numbers (no decimal): int int8 int16 int32 int64
- positive whole numbers (no decimal): uint uint8 uint16 uint32 uint64 uintptr >>> unsigned integer
- signed decimal numbers: float32 float64
- imaginary numbers: complex64 complex128 >>> rarely used
- size (aka the number after the data_type) >>> tells us how many bits in memory will be used
- standard size usage: int uint float64 complex128
- type convertion: casting, truncate the rest of the data, like from 88.26 casts as int64, the floating point will be gone

# How to choose the type
- use the default type, like the type we defined is not the same as the type we are trying to input, cast into int from uint16
- deafult types: bool, string, int, uint, byte, rune, float64, complex128

# Go is statically typed
- types are known before the code runs
- type error will be thrown before the code is run

# Compiled Program
- the one we can run without original source code >>> compiled program
- no need compiler after it is done its job
- download the exe and we can run it
- only the comiled executable runs in production non compiler required
- run the compiled program, users need compiled executable
- GO, C, C++, Rust

# Interpreted Program
- code is interpreted at runtime by another program
- end user need interpreter so they can run the code
- code and interpreter required in production
- JS, Python, Ruby

# Single Line Declarations
- mileage, company := 802, "Hello"

# Small Memory Footpring
- lightweight
- program need small amount extra code, included in the executable binary >>> go runtime
- we need go run, to clean up the unused memory = garbage collector, free up memory which is no longer in use
- compare to java and go, go need more memory, cause the java is using virtual machine to interpret bytecode at runtime, make it allocating more heap
- and rust and c use less memory, due to that program allows developer to handle the memory usage, but go handles it auto

# What is heap
- memoery will be taken up: new variales...
- stack memory, store the LOCAL variables
- heap memory, DYNAMIC for programmer to allocate
- data: store global varaible >>> uninit and init data
- text: store the code being executed
- address location, stack > heap > data > text
- address with base 16 numbers
- 
# Env
```bash
PORT=3000
```




