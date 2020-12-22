#!/usr/local/bin/bash
curl -X GET http://localhost:8080/
curl -X GET http://localhost:8080/probe/ready
curl -X GET http://localhost:8080/probe/live
curl -X GET http://localhost:8080/buildInfo
