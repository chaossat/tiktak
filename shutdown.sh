#!/bin/bash

kill $(ps x | grep TiktakRelase | awk '{print $1}')