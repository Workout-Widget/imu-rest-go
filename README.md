# Sense HAT REST API
## Overview
This is a minimalist REST API designed to collect IMU data from WiFi/network devices.
# How To Run
First navigate to the `docker/` directory and start the database:
```
docker-compose up -d
```
Next, navigate to the `src/main/` directory and build the executable:
```
go build
```
Run the executable

Linux/MacOS
```
./sensehatrest
```

Windows
```
sensehatrest.exe
```