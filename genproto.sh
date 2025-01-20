#!/bin/bash
cd common/proto && protoc --go_out=. --micro_out=.  ./$1.proto