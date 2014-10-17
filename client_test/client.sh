#!/bin/sh
./client_main -sender=1 -receiver=2 -v=1 -stderrthreshold=0 &
./client_main -sender=3 -receiver=2 -v=1 -stderrthreshold=0 &
./client_main -sender=4 -receiver=2 -v=1 -stderrthreshold=0 &
./client_main -sender=5 -receiver=2 -v=1 -stderrthreshold=0 &
./client_main -sender=6 -receiver=2 -v=1 -stderrthreshold=0 &