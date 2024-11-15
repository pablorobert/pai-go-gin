FROM golang:1.23 

#set working directory
WORKDIR /go/src/app 

#Copy the Go app
COPY . . 

#Expose 
EXPOSE 8080

# Bulid the Go app
RUN go build -o main ./cli/main.go 

#Run the executable 
CMD ["./main"]

