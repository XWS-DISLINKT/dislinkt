#!/bin/(shell)

full_path=$1
files=$(echo $full_path | tr "/" " ")
files=($files)
path=""

declare -i i=1
for file in ${files[@]}
do
    if [ ${#files[@]} -gt $i ]
    then
        path="$path/$file"
    fi
    i=$i+1
done

protoc -I "./common/proto$path" \
       --go_out "./common/proto$path" --go_opt paths=source_relative \
       --go-grpc_out "./common/proto$path" --go-grpc_opt paths=source_relative \
       --grpc-gateway_out "./common/proto$path" --grpc-gateway_opt paths=source_relative \
       "./common/proto$full_path"

read -p "Press any key to resume ..."