#!/bin/bash

echo "Configuring project..."

echo "Creating build folders..."
if [ ! -d "cpp/bin" ]; then
  mkdir -p cpp/bin
fi
if [ ! -d "bin" ]; then
  mkdir bin
fi
if [ ! -L "bin/lib" ]; then
  cd bin && ln -s ../cpp/lib lib && cd ..
fi

echo "Copying scripts..."
cp scripts/run.sh bin
