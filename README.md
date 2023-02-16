# League Backend Server

# What it does?
League server serves following requests

Given an uploaded csv file
```
1,2,3
4,5,6
7,8,9
```

1. Echo (given)
    - Return the matrix as a string in matrix format.
    
    ```
    // Expected output
    1,2,3
    4,5,6
    7,8,9
    ``` 
2. Invert
    - Return the matrix as a string in matrix format where the columns and rows are inverted
    ```
    // Expected output
    1,4,7
    2,5,8
    3,6,9
    ``` 
3. Flatten
    - Return the matrix as a 1 line string, with values separated by commas.
    ```
    // Expected output
    1,2,3,4,5,6,7,8,9
    ``` 
4. Sum
    - Return the sum of the integers in the matrix
    ```
    // Expected output
    45
    ``` 
5. Multiply
    - Return the product of the integers in the matrix
    ```
    // Expected output
    362880
    ``` 

## Steps to build (MacOS)
### Prerequisites
1. Install [homebrew](https://brew.sh)
2. Install golang
```
brew install go
```
3. Install [bazel](https://bazel.build/install) in your system.
```
brew install bazel
```
4. Install [gazelle](https://github.com/bazelbuild/bazel-gazelle/blob/15828e8077542449fea13491bcb2e404a12412fd/README.rst#running-gazelle-with-bazel)

### Clone Repo
1. Make directory and cd into it to clone the project.
```
mkdir -p ${GOPATH}/src/github.com/palak92/ 
```
Note : substitute your GOPATH
2. Checkout the code
```
git clone https://www.github.com/palak92/league.git
```
3. Change directory to workspace
```
cd ${GOPATH}/src/github.com/palak92/league
```
### Build
1. Generate build files
```
bazel run //:gazelle
```
8. Run build command
```
bazel build //...
```
8. Run unit tests
```
go test ./...
```

## Steps to Run Server (MacOS)
1. Run web server
```
bazel run //cmd:cmd
```
2. Send request to the server from another client
```
curl -F 'file=@/path/matrix.csv' "localhost:8080/echo"
```

## Known Issues
1. Cannot run tests through bazel configuration. Workaround is to run test via go test util.