# Real-Time Server Health Monitor

![imagen](https://github.com/user-attachments/assets/19ab0608-8de0-47fe-8914-6e364ea147d1)

The idea of this project is to get info about a server's hardware in real time to future monitoring.

## Techonologies Used
- Go
- Gorilla Websocket
- Echo Framework
- HTMX
- github.com/shirou/gopsutil/v4 (to get the hardware info)

## To use
1. Clone this repo
2. in the repo folder do 
```go
go run main.go
```
3. Go to localhost:1313/ to view your system's data
   

### Future Ideas
- In case of horizontal scaling, be able to monitor each server and load
- Connect this to a database and obtain server's load data to analyze trend



