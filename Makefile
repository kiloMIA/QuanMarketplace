# Create migrations
# migrate create -ext sql -dir migrations <name_of_migration>

# generate python proto files 
# python -m grpc_tools.protoc -I. --python_out=. --grpc_python_out=. <name_of_file>.proto 
# generate go proto files
# protoc --go_out=. --go_opt=paths=source_relative       --go-grpc_out=. --go-grpc_opt=paths=source_relative       <name_of_file>.proto 