#!/bin/bash

exe_name="run"
compiler_flags="-Wall"

source_dir=(../singlyLinkedList/)


include_flag=""
for dir in "${source_dir[@]}"; do
	include_flag+="-I $dir "
done

source_files=""
for dir in "${source_dir[@]}"; do
	file=($(find "$dir" -name "*.cpp")) # && echo "$file"
	source_files+=$file
done

compiler_command="g++ test.cpp ${source_files[@]} $include_flag $compiler_flags -o $exe_name"

echo "Executing '$compiler_command'"
eval "$compiler_command"
if [ 0 -eq $? ]; then
	echo "Success"
fi
